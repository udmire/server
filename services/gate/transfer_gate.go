package gate

import (
	"encoding/binary"
	"io"

	transfer_gate "e.coding.net/mmstudio/blade/gate"
	"e.coding.net/mmstudio/blade/golib/module"
	"e.coding.net/mmstudio/blade/golib/net/link"
	"e.coding.net/mmstudio/blade/server/transport"
	"e.coding.net/mmstudio/blade/server/utils"
	log "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var transferCodec = &TransferCodec{}

type TransferGate struct {
	cg       *transfer_gate.Gate
	g        *Gate
	selector *TransferSelector
}

func NewTransferGate(ctx *cli.Context, g *Gate) *TransferGate {
	transferGate := &TransferGate{
		g:        g,
		selector: NewTransferSelector(g),
	}

	fp := link.NewProtocol(link.WithProtocolOptionMaxRecv(int(transport.TcpPacketMaxSize)),
		link.WithProtocolOptionMaxSend(int(transport.TcpPacketMaxSize)),
		link.WithProtocolOptionByteOrder(binary.LittleEndian),
	)
	spec := transfer_gate.NewSpec(
		transfer_gate.WithEnableWebSocket(true),
		transfer_gate.WithTransferProvider(func(writer io.ReadWriter) (link.Transporter, error) {
			return fp.NewTransporter(writer)
		}),
		transfer_gate.WithMessageCodec(transferCodec),
		transfer_gate.WithEnableZeroCopy(true),
		transfer_gate.WithEnableXListener(false),
		transfer_gate.WithSelector(transferGate.selector),
		transfer_gate.WithPatchPath("config/patch"),
	)

	cg, err := transfer_gate.New(spec)
	if !utils.ErrCheck(err, "New transfer gate failed", spec) {
		return nil
	}

	transferGate.cg = cg

	spec.Filter = append(spec.Filter, transferGate.selector.ConsistentHashFilter)
	if err != nil {
		log.Fatal().Err(err).Msg("new transfer gate server failed")
		return nil
	}

	return transferGate
}

func (cg *TransferGate) Run(ctx *cli.Context) error {
	module.Run(cg.cg)
	return nil
}
