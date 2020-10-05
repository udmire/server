package gate

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/client"
	log "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	pbGame "github.com/yokaiio/yokai_server/proto/game"
	pbGate "github.com/yokaiio/yokai_server/proto/gate"
	"github.com/yokaiio/yokai_server/utils"
)

type RpcHandler struct {
	g       *Gate
	gameSrv pbGame.GameService
}

func NewRpcHandler(g *Gate, ucli *cli.Context) *RpcHandler {
	h := &RpcHandler{
		g: g,
		gameSrv: pbGame.NewGameService(
			"yokai_game",
			g.mi.srv.Client(),
		),
	}

	pbGate.RegisterGateServiceHandler(g.mi.srv.Server(), h)

	return h
}

/////////////////////////////////////////////
// rpc call
/////////////////////////////////////////////
func (h *RpcHandler) CallGetRemoteLitePlayer(id int64) (*pbGame.GetRemoteLitePlayerReply, error) {
	req := &pbGame.GetRemoteLitePlayerRequest{Id: id}
	return h.gameSrv.GetRemoteLitePlayer(context.Background(), req, client.WithSelectOption(utils.SectionIDRandSelector(id)))
}

func (h *RpcHandler) CallUpdatePlayerExp(id int64) (*pbGame.UpdatePlayerExpReply, error) {
	req := &pbGame.UpdatePlayerExpRequest{Id: id}
	return h.gameSrv.UpdatePlayerExp(context.Background(), req)
}

/////////////////////////////////////////////
// rpc receive
/////////////////////////////////////////////
func (h *RpcHandler) GetGateStatus(ctx context.Context, req *pbGate.GateEmptyMessage, rsp *pbGate.GetGateStatusReply) error {
	rsp.Status = &pbGate.GateStatus{GateId: int32(h.g.ID), Health: 2}
	return nil
}

func (h *RpcHandler) UpdateUserInfo(ctx context.Context, req *pbGate.UpdateUserInfoRequest, rsp *pbGate.GateEmptyMessage) error {
	defer log.Info().Interface("request", req).Msg("update user info")
	return h.g.gs.UpdateUserInfo(req)
}

func (h *RpcHandler) SyncPlayerInfo(ctx context.Context, req *pbGate.SyncPlayerInfoRequest, rsp *pbGate.SyncPlayerInfoReply) error {
	tm := time.Now()
	defer func() {
		d := time.Since(tm)
		if d > time.Second*2 {
			log.Warn().
				Dur("latency", d).
				Int64("user_id", req.UserId).
				Msg("rpc receive handle timeout")
		}
	}()

	info, err := h.g.gs.getUserInfo(req.UserId)
	if err != nil {
		log.Warn().
			Int64("user_id", req.UserId).
			Err(err).
			Msg("rpc receive handle SyncPlayerInfo failed")
		return fmt.Errorf("RpcHandler.SyncPlayerInfo failed: %w", err)
	}

	rsp.Info = &pbGate.UserInfo{
		UserId:      info.UserID,
		AccountId:   info.AccountID,
		GameId:      int32(info.GameID),
		PlayerId:    info.PlayerID,
		PlayerName:  info.PlayerName,
		PlayerLevel: info.PlayerLevel,
	}

	return nil
}
