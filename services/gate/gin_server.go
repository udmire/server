package gate

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	// "github.com/east-eden/gate/msg"
	"github.com/east-eden/server/logger"
	"github.com/east-eden/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
)

var (
	httpReadTimeout           = time.Second * 5
	httpWriteTimeout          = time.Second * 5
	ginConcurrentRequestLimit = 1000
)

var (
	opsSelectGameCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "gate",
		Name:      "select_game_ops",
		Help:      "选择服务器操作总数",
	})

	timeCounterHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "gate",
		Name:      "select_game_addr_latency",
		Help:      "请求延迟",
	},
		[]string{"method"},
	)
)

type GinServer struct {
	g         *Gate
	router    *gin.Engine
	tlsRouter *gin.Engine
	wg        utils.WaitGroupWrapper
}

// wrap http.HandlerFunc to gin.HandlerFunc
func ginHandlerWrapper(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func (s *GinServer) setupHttpRouter() {
	s.router.Use(limit.MaxAllowed(ginConcurrentRequestLimit))
	s.router.Use(gin.LoggerWithWriter(logger.Logger))

	// pprof
	s.router.GET("/debug/pprof", ginHandlerWrapper(pprof.Index))
	s.router.GET("/debug/cmdline", ginHandlerWrapper(pprof.Cmdline))
	s.router.GET("/debug/symbol", ginHandlerWrapper(pprof.Symbol))
	s.router.GET("/debug/profile", ginHandlerWrapper(pprof.Profile))
	s.router.GET("/debug/trace", ginHandlerWrapper(pprof.Trace))
	s.router.GET("/debug/allocs", ginHandlerWrapper(pprof.Handler("allocs").ServeHTTP))
	s.router.GET("/debug/heap", ginHandlerWrapper(pprof.Handler("heap").ServeHTTP))
	s.router.GET("/debug/goroutine", ginHandlerWrapper(pprof.Handler("goroutine").ServeHTTP))
	s.router.GET("/debug/block", ginHandlerWrapper(pprof.Handler("block").ServeHTTP))
	s.router.GET("/debug/threadcreate", ginHandlerWrapper(pprof.Handler("threadcreate").ServeHTTP))

	// test
	s.router.GET("/health_check/Location", func(c *gin.Context) {
		var req struct {
			ServiceID string `json:"ServiceID"`
		}

		if err := c.Bind(&req); err != nil {
			log.Warn().
				Err(err).
				Msg("select_game_addr request bind failed")

			c.String(http.StatusBadRequest, "bad request:%s", err.Error())
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, "pass!")
	})

	// test transfer message
	// s.router.GET("/transfer", func(c *gin.Context) {
	// 	filter := s.g.cg.selector.ConsistentHashFilter(&msg.Handshake{UserID: "test_user1"})
	// 	entry := filter(nil)
	// 	log.Info().Interface("node", entry).Msg("select game node")
	// 	c.JSON(http.StatusOK, "pass")
	// })

	// select_game_addr
	s.router.POST("/select_game_addr", func(c *gin.Context) {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			timeCounterHistogram.WithLabelValues("/select_game_addr").Observe(v)
		}))
		defer timer.ObserveDuration()

		opsSelectGameCounter.Inc()

		var req struct {
			UserID string `json:"userId"`
		}

		if err := c.Bind(&req); err != nil {
			log.Warn().
				Err(err).
				Msg("select_game_addr request bind failed")

			c.String(http.StatusBadRequest, "bad request:%s", err.Error())
			return
		}

		if user, metadata := s.g.gs.SelectGame(req.UserID); user != nil {
			h := gin.H{
				"userId":        req.UserID,
				"userName":      user.PlayerName,
				"accountId":     user.AccountID,
				"gameId":        metadata["gameId"],
				"publicTcpAddr": metadata["publicTcpAddr"],
				"publicWsAddr":  metadata["publicWsAddr"],
			}
			c.JSON(http.StatusOK, h)

			log.Info().
				Interface("gin.H", h).
				Msg("select_game_addr calling with result")
			return
		}

		c.String(http.StatusBadRequest, fmt.Sprintf("cannot find account by userid<%s>", req.UserID))
	})

	// metrics
	metricsHandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{Registry: prometheus.DefaultRegisterer})
	s.router.GET("/metrics", ginHandlerWrapper(metricsHandler.ServeHTTP))
}

func (s *GinServer) setupHttpsRouter() {
	s.tlsRouter.Use(limit.MaxAllowed(ginConcurrentRequestLimit))
	s.router.Use(gin.LoggerWithWriter(logger.Logger))

	// store_write
	s.tlsRouter.POST("/store_write", func(c *gin.Context) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		if c.Bind(&req) == nil {
			if err := s.g.mi.StoreWrite(req.Key, req.Value); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("store write failed: %s", err.Error()))
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		c.String(http.StatusBadRequest, "bad request")
	})

	// pub_gate_result
	s.tlsRouter.POST("/pub_gate_result", func(c *gin.Context) {
		if err := s.g.GateResult(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "status ok")
	})

	// update_player_exp
	s.tlsRouter.POST("/update_player_exp", func(c *gin.Context) {
		var req struct {
			Id string `json:"id"`
		}

		if c.Bind(&req) == nil {
			id := cast.ToInt64(req.Id)
			r, err := s.g.rpcHandler.CallUpdatePlayerExp(id)
			c.String(http.StatusOK, "UpdatePlayerExp result", r, err)

			// test storage
			//user := NewUserInfo().(*UserInfo)
			//if err := store.GetStore().LoadObject(store.StoreType_User, "_id", id, user); err != nil {
			//logger.Warn(err)
			//}

			//user.UserID = id
			//user.PlayerName = "dudu"
			//if err := store.GetStore().SaveObject(store.StoreType_User, user); err != nil {
			//logger.Warn(err)
			//}

			//user.PlayerLevel++
			//user.PlayerName += "."
			//fields := map[string]any{
			//"player_level": user.PlayerLevel,
			//"player_name":  user.PlayerName,
			//}
			//if err := store.GetStore().SaveFields(store.StoreType_User, user, fields); err != nil {
			//logger.Warn(err)
			//}
		}

	})

	// get_lite_player
	s.tlsRouter.POST("/get_lite_player", func(c *gin.Context) {
		var req struct {
			PlayerId string `json:"playerId"`
		}

		if c.Bind(&req) == nil {
			id := cast.ToInt64(req.PlayerId)
			rep, err := s.g.rpcHandler.CallGetRemotePlayerInfo(id)
			if err == nil {
				c.JSON(http.StatusOK, rep)
				return
			}

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusBadRequest, "request error")
	})

}

func NewGinServer(ctx *cli.Context, g *Gate) *GinServer {
	if ctx.Bool("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &GinServer{
		g:         g,
		router:    gin.Default(),
		tlsRouter: gin.Default(),
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info().Msgf("[GIN-debug] %s %s %s %d", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	s.setupHttpRouter()
	s.setupHttpsRouter()
	return s
}

func (s *GinServer) Main(ctx *cli.Context) error {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Fatal().
					Err(err).
					Msg("GinServer Run() failed")
			}
			exitCh <- err
		})
	}

	s.wg.Wrap(func() {
		utils.CaptureException()
		exitFunc(s.Run(ctx))
	})

	// listen https
	go func() {
		defer utils.CaptureException()

		certPath := ctx.String("cert_path_release")
		keyPath := ctx.String("key_path_release")
		if ctx.Bool("debug") {
			certPath = ctx.String("cert_path_debug")
			keyPath = ctx.String("key_path_debug")
		}

		server := &http.Server{
			Addr:         ctx.String("https_listen_addr"),
			Handler:      s.tlsRouter,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
		}

		if err := server.ListenAndServeTLS(certPath, keyPath); err != nil {
			log.Error().
				Err(err).
				Msg("GinServer RunTLS failed")
			exitCh <- err
		}
	}()

	// listen http
	go func() {
		defer utils.CaptureException()

		server := &http.Server{
			Addr:         ctx.String("http_listen_addr"),
			Handler:      s.router,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Error().
				Err(err).
				Msg("GinServer Run failed")
			exitCh <- err
		}
	}()

	return <-exitCh
}

func (s *GinServer) Run(ctx *cli.Context) error {
	<-ctx.Done()
	log.Info().Msg("GinServer context done...")
	return nil
}

func (s *GinServer) Exit(ctx context.Context) {
	s.wg.Wait()
	log.Info().Msg("gin server exit...")
}
