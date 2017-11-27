package gate

import (
	"server/game"
	"server/msg"
	"server/proto_game"
)

func init() {
	msg.Processor.SetRouter(&proto_game.Hello{}, game.ChanRPC)
}
