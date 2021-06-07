package gate

import (
	"encoding/binary"
	"errors"
	"hash/crc32"

	"e.coding.net/mmstudio/blade/golib/encoding2"
	"e.coding.net/mmstudio/blade/server/utils"
	"github.com/valyala/bytebufferpool"
	"google.golang.org/protobuf/proto"
)

var (
	ErrCodecInvalidType = errors.New("invalid codec type")
)

type TransferCodec struct {
	encoding2.Codec
}

func (c *TransferCodec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(proto.Message)
	if !ok {
		return nil, ErrCodecInvalidType
	}

	body, err := proto.Marshal(p)
	if !utils.ErrCheck(err, "proto marshal failed when CommonCodec.Marshal", p) {
		return nil, err
	}

	protoName := p.ProtoReflect().Descriptor().Name()
	var nameCrc uint32 = crc32.ChecksumIEEE([]byte(protoName))
	buffer := bytebufferpool.Get()
	defer bytebufferpool.Put(buffer)

	_ = binary.Write(buffer, binary.LittleEndian, uint32(nameCrc))
	_, _ = buffer.Write(body)

	return buffer.Bytes(), nil
}

func (c *TransferCodec) Unmarshal(data []byte, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return ErrCodecInvalidType
	}

	return proto.Unmarshal(data[4:], p)
}

func (c *TransferCodec) Name() string {
	return "transfer_codec"
}
