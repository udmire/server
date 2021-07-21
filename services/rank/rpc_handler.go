package rank

import (
	"context"
	"errors"
	"fmt"
	"time"

	pbGame "e.coding.net/mmstudio/blade/server/proto/server/game"
	pbRank "e.coding.net/mmstudio/blade/server/proto/server/rank"
	"e.coding.net/mmstudio/blade/server/utils"
	"github.com/asim/go-micro/v3/client"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	ErrInvalidGlobalConfig = errors.New("invalid global config")
)

var (
	DefaultRpcTimeout = 5 * time.Second // 默认rpc超时时间
)

type RpcHandler struct {
	m       *Rank
	rankSrv pbRank.RankService
	gameSrv pbGame.GameService
}

func NewRpcHandler(cli *cli.Context, m *Rank) *RpcHandler {
	h := &RpcHandler{
		m: m,
		rankSrv: pbRank.NewRankService(
			"rank",
			m.mi.srv.Client(),
		),
		gameSrv: pbGame.NewGameService(
			"game",
			m.mi.srv.Client(),
		),
	}

	err := pbRank.RegisterRankServiceHandler(m.mi.srv.Server(), h)
	if err != nil {
		log.Fatal().Err(err).Msg("RegisterRankServiceHandler failed")
	}

	return h
}

// 一致性哈希
func (h *RpcHandler) consistentHashCallOption(key string) client.CallOption {
	return client.WithSelectOption(
		utils.ConsistentHashSelector(h.m.cons, key),
	)
}

// 重试次数
func (h *RpcHandler) retries(times int) client.CallOption {
	return client.WithRetries(times)
}

/////////////////////////////////////////////
// rpc call
/////////////////////////////////////////////
func (h *RpcHandler) CallKickRankData(rankId int64, nodeId int32) (*pbRank.KickRankDataRs, error) {
	if rankId == -1 {
		return nil, errors.New("invalid rank data id")
	}

	if nodeId == int32(h.m.ID) {
		return nil, errors.New("same rank node id")
	}

	req := &pbRank.KickRankDataRq{
		RankId:     rankId,
		RankNodeId: nodeId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultRpcTimeout)
	defer cancel()

	return h.rankSrv.KickRankData(
		ctx,
		req,
		client.WithSelectOption(
			utils.SpecificIDSelector(
				fmt.Sprintf("rank-%d", nodeId),
			),
		),
	)
}

/////////////////////////////////////////////
// rpc receive
/////////////////////////////////////////////
// 踢出邮件cache
func (h *RpcHandler) KickRankData(
	ctx context.Context,
	req *pbRank.KickRankDataRq,
	rsp *pbRank.KickRankDataRs,
) error {
	return h.m.manager.KickRankData(req.GetRankId(), req.GetRankNodeId())
}
