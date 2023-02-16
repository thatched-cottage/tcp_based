package pkg

import (
	"fmt"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/command"
	handle2 "github.com/thatched-cottage/tcp_based.git/pkg/handle"
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
