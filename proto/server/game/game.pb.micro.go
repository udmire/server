// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/game/game.proto

package game

import (
	fmt "fmt"
	_ "github.com/east-eden/server/proto/global"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GameService service

func NewGameServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GameService service

type GameService interface {
	GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, opts ...client.CallOption) (*GetRemotePlayerInfoRs, error)
	KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, opts ...client.CallOption) (*KickAccountOfflineRs, error)
	ExpirePlayerMail(ctx context.Context, in *ExpirePlayerMailRq, opts ...client.CallOption) (*ExpirePlayerMailRs, error)
	// test
	UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, opts ...client.CallOption) (*UpdatePlayerExpReply, error)
}

type gameService struct {
	c    client.Client
	name string
}

func NewGameService(name string, c client.Client) GameService {
	return &gameService{
		c:    c,
		name: name,
	}
}

func (c *gameService) GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, opts ...client.CallOption) (*GetRemotePlayerInfoRs, error) {
	req := c.c.NewRequest(c.name, "GameService.GetRemotePlayerInfo", in)
	out := new(GetRemotePlayerInfoRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, opts ...client.CallOption) (*KickAccountOfflineRs, error) {
	req := c.c.NewRequest(c.name, "GameService.KickAccountOffline", in)
	out := new(KickAccountOfflineRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) ExpirePlayerMail(ctx context.Context, in *ExpirePlayerMailRq, opts ...client.CallOption) (*ExpirePlayerMailRs, error) {
	req := c.c.NewRequest(c.name, "GameService.ExpirePlayerMail", in)
	out := new(ExpirePlayerMailRs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameService) UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, opts ...client.CallOption) (*UpdatePlayerExpReply, error) {
	req := c.c.NewRequest(c.name, "GameService.UpdatePlayerExp", in)
	out := new(UpdatePlayerExpReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameService service

type GameServiceHandler interface {
	GetRemotePlayerInfo(context.Context, *GetRemotePlayerInfoRq, *GetRemotePlayerInfoRs) error
	KickAccountOffline(context.Context, *KickAccountOfflineRq, *KickAccountOfflineRs) error
	ExpirePlayerMail(context.Context, *ExpirePlayerMailRq, *ExpirePlayerMailRs) error
	// test
	UpdatePlayerExp(context.Context, *UpdatePlayerExpRequest, *UpdatePlayerExpReply) error
}

func RegisterGameServiceHandler(s server.Server, hdlr GameServiceHandler, opts ...server.HandlerOption) error {
	type gameService interface {
		GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, out *GetRemotePlayerInfoRs) error
		KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, out *KickAccountOfflineRs) error
		ExpirePlayerMail(ctx context.Context, in *ExpirePlayerMailRq, out *ExpirePlayerMailRs) error
		UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, out *UpdatePlayerExpReply) error
	}
	type GameService struct {
		gameService
	}
	h := &gameServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GameService{h}, opts...))
}

type gameServiceHandler struct {
	GameServiceHandler
}

func (h *gameServiceHandler) GetRemotePlayerInfo(ctx context.Context, in *GetRemotePlayerInfoRq, out *GetRemotePlayerInfoRs) error {
	return h.GameServiceHandler.GetRemotePlayerInfo(ctx, in, out)
}

func (h *gameServiceHandler) KickAccountOffline(ctx context.Context, in *KickAccountOfflineRq, out *KickAccountOfflineRs) error {
	return h.GameServiceHandler.KickAccountOffline(ctx, in, out)
}

func (h *gameServiceHandler) ExpirePlayerMail(ctx context.Context, in *ExpirePlayerMailRq, out *ExpirePlayerMailRs) error {
	return h.GameServiceHandler.ExpirePlayerMail(ctx, in, out)
}

func (h *gameServiceHandler) UpdatePlayerExp(ctx context.Context, in *UpdatePlayerExpRequest, out *UpdatePlayerExpReply) error {
	return h.GameServiceHandler.UpdatePlayerExp(ctx, in, out)
}
