package pkg

import (
	"fmt"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/command"
	handle2 "gitee.com/wuxiansheng/tcp_based.git/pkg/handle"
)

type Packet struct {
	Adder string
	Msg   []byte
}

var (
	handles = map[command.CommandType]msgHandle{}
)

func init() {
	handles[command.CommandConn] = handle2.CommandConnHandle
	handles[command.CommandMirror] = handle2.CommandMirrorHandle
}

type msgHandle func(body *[]byte) (*[]byte, error)

func Decode(info *base_info.PkgInfo) (*[]byte, error) {
	commandID := info.Command
	if f, has := handles[info.Command]; has {
		return f(&info.Msg)
	} else {
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}
