package msg

import (
	"server/proto_game"

	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&proto_game.Hello{})
}

// type Hello struct {
// 	Name string
// }
