package central_node

import (
	"fmt"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	"net"
	"sync"
)

type ExtranetNetwork struct {
	rq  readQueue
	wqs sync.Map
}

func init() {
	ININP.IENP = &ExtranetNetwork{}
	ININP.IENP.Init()
}

func (this *ExtranetNetwork) Init() {
	log.Debugf("Init")
	this.rq.q.Init()
	go this.rq.Handle()
	this.wqs = sync.Map{}
}
func (this *ExtranetNetwork) SendMsg(pkgInfo *base_info.PkgInfo) {
	log.Debugf("SendMsg to %s", pkgInfo.PkgInfoTargetNodeName())
	wq, ok := this.wqs.Load(string(pkgInfo.ClientId))
	if ok == true {
		wq.(*writeQueue).PushBack(pkgInfo.B)
	} else {
		log.Debugf("The terminal is disconnected")
	}
}
func (this *ExtranetNetwork) SendsyncMsg(pkgInfo *base_info.PkgInfo) {

}
func (this *ExtranetNetwork) Run(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Debugf("listen error:%v", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Debugf("accept error:%v", err)
			break
		}
		this.HandleConn(c)
	}
}
func (this *ExtranetNetwork) HandleConn(c net.Conn) {
	clientId, err := this.RegisterNode(c)
	if err != nil {
		c.Close()
		return
	}
	go this.ReadHandle(clientId, c)
}
func (this *ExtranetNetwork) ReadHandle(clientId []byte, c net.Conn) {
	defer c.Close()
	for {
		b := make([]byte, base_info.ByteLength) //base_info.ByteLength 包的长度
		n, err := c.Read(b)
		log.Debugf("read byte:%v", b)
		if err != nil {
			wq, ok := this.wqs.Load(string(clientId))
			if ok == true {
				wq.(*writeQueue).Close()
			}
			log.Debugf("conn read err:%s", err.Error())
			return
		}
		if n == 0 {
			log.Debugf("conn read msg err: n == 0")
			continue
		}

		pkgtype := (b)[0]
		switch pkgtype {
		case base_info.CommonPkg:
			this.rq.PushBack(&b)
		case base_info.HeartbeatPkg:
		default:
		}
	}
}
func (this *ExtranetNetwork) RegisterNode(c net.Conn) ([]byte, error) {
	b := make([]byte, base_info.ByteLength) //base_info.ByteLength 包的长度
	_, err := c.Read(b)
	if err != nil {
		log.Debugf("conn read err:%s", err.Error())
		return nil, err
	}
	pkgInfo := base_info.BToPkgInfo(&b)
	switch pkgInfo.PkgType {
	case base_info.RegisterClientPkg:
		wqi, ok := this.wqs.Load(string(pkgInfo.ClientId))
		if ok == true {
			log.Debugf("map adder:%p", wqi)
			wqi.(*writeQueue).Close()
		}
		wq := &writeQueue{}
		this.wqs.Store(string(pkgInfo.ClientId), wq)
		wq.Init()
		log.Debugf("map adder:%p", &wq)
		go func() {
			wq.HandleConn(c, pkgInfo.Source)
		}()
		wq.PushBack(&b)

		log.Debugf("register client id :%s ", pkgInfo.ClientId)
	default:
		log.Errorf("pkg type error")
		return nil, fmt.Errorf("pkg type error")
	}
	return pkgInfo.ClientId, nil
}
