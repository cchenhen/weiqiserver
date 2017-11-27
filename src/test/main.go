// package main

// import (
// 	"encoding/binary"
// 	"net"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:3563")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Hello 消息（JSON 格式）
// 	// 对应游戏服务器 Hello 消息结构体
// 	data := []byte(`{
// 		"Hello": {
// 			"Name": "cc_test"
// 		}
// 	}`)

// 	// len + data
// 	m := make([]byte, 2+len(data))

// 	// 默认使用大端序
// 	binary.BigEndian.PutUint16(m, uint16(len(data)))

// 	copy(m[2:], data)

// 	// 发送消息
// 	conn.Write(m)
// }

//------------------test_json--------------------

//------------------test_proto-------------------
package main

import (

	// 辅助库
	"encoding/binary"
	"net"

	"github.com/golang/protobuf/proto"
	// test.pb.go 的路径
	"server/proto_game"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	test := &proto_game.Hello{
		// 使用辅助函数设置域的值
		Name: proto.String("handsome cc"),
	}
	data, err := proto.Marshal(test)
	if err != nil {
		panic(err)
	}
	m := make([]byte, 2+len(data))
	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
}
