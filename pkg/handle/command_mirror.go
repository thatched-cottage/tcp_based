package handle

import (
	"fmt"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	"unsafe"
)

type CommandMirrorReq struct {
	Msg string
}

func MarshalCommandMirrorReq(r *CommandConnReq) []byte {
	Len := unsafe.Sizeof(*r)
	s := &S{
		addr: uintptr(unsafe.Pointer(r)),
		cap:  int(Len),
		len:  int(Len),
	}
	return *(*[]byte)(unsafe.Pointer(s))
}

func UnMarshalCommandMirrorReq(d *[]byte) *CommandConnReq {
	fmt.Println("[]byte is : ", d)
	return *(**CommandConnReq)(unsafe.Pointer(&d))
}

type CommandMirrorRes struct {
	Msg string
	err error
}

func CommandMirrorHandle(body *[]byte) (*[]byte, error) {
	log.Debugf("CommandMirrorHandle.Msg:%s", string(*body))
	// 往下层传递，下层返回值转成 byte
	return body, nil
}
