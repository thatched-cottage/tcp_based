package handle

import (
	"fmt"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"unsafe"
)

type CommandConnReq struct {
	Msg string
}

func MarshalCommandConnReq(r *CommandConnReq) []byte {
	Len := unsafe.Sizeof(*r)
	s := &S{
		addr: uintptr(unsafe.Pointer(r)),
		cap:  int(Len),
		len:  int(Len),
	}
	return *(*[]byte)(unsafe.Pointer(s))
}

func UnMarshalCommandConnReq(d *[]byte) *CommandConnReq {
	fmt.Println("[]byte is : ", d)
	return *(**CommandConnReq)(unsafe.Pointer(&d))
}

type CommandConnRes struct {
	Msg string
	err error
}

func CommandConnHandle(body *[]byte) (*[]byte, error) {
	var req *CommandConnReq = UnMarshalCommandConnReq(body)
	log.Debugf("CommandConnReq.Msg:", req.Msg)
	// 往下层传递，下层返回值转成 byte
	return &[]byte{}, nil
}
