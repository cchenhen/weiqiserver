package internal

import (
	"reflect"
	"server/proto_game"

	"github.com/golang/protobuf/proto"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&proto_game.Hello{}, testHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func testHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*proto_game.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("hello cc:%v", m.Name)

	cc := "Mr ZhenZheng gay gay!"
	a.WriteMsg(&proto_game.Hello{
		Name: proto.String(cc),
	})
}
