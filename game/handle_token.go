package game

import (
	logger "github.com/sirupsen/logrus"
	"github.com/yokaiio/yokai_server/internal/define"
	"github.com/yokaiio/yokai_server/internal/transport"
	pbGame "github.com/yokaiio/yokai_server/proto/game"
)

func (m *MsgHandler) handleAddToken(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.ID(),
			"account_name": acct.Name(),
		}).Warn("add token failed")
		return
	}

	if acct.Player() == nil {
		return
	}

	msg, ok := p.Body.(*pbGame.MC_AddToken)
	if !ok {
		logger.Warn("Add Token failed, recv message body error")
		return
	}

	acct.PushWrapHandler(func() {
		err := acct.Player().TokenManager().TokenInc(msg.Type, msg.Value)
		if err != nil {
			logger.Warn("token inc failed:", err)
		}

		acct.Player().TokenManager().Save()

		reply := &pbGame.MS_TokenList{Tokens: make([]*pbGame.Token, 0, define.Token_End)}
		for n := 0; n < define.Token_End; n++ {
			v, err := acct.Player().TokenManager().GetToken(int32(n))
			if err != nil {
				logger.Warn("token get value failed:", err)
				return
			}

			t := &pbGame.Token{
				Type:    v.ID,
				Value:   v.Value,
				MaxHold: v.MaxHold,
			}
			reply.Tokens = append(reply.Tokens, t)
		}
		acct.SendProtoMessage(reply)
	})

}

func (m *MsgHandler) handleQueryTokens(sock transport.Socket, p *transport.Message) {
	acct := m.g.am.GetAccountBySock(sock)
	if acct == nil {
		logger.WithFields(logger.Fields{
			"account_id":   acct.ID(),
			"account_name": acct.Name(),
		}).Warn("query tokens failed")
		return
	}

	if acct.Player() == nil {
		return
	}

	reply := &pbGame.MS_TokenList{Tokens: make([]*pbGame.Token, 0, define.Token_End)}
	for n := 0; n < define.Token_End; n++ {
		v, err := acct.Player().TokenManager().GetToken(int32(n))
		if err != nil {
			logger.Warn("token get value failed:", err)
			return
		}

		t := &pbGame.Token{
			Type:    v.ID,
			Value:   v.Value,
			MaxHold: v.MaxHold,
		}
		reply.Tokens = append(reply.Tokens, t)
	}
	acct.SendProtoMessage(reply)
}
