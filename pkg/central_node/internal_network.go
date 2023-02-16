package central_node

import (
	"fmt"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	"net"
)

type InternalNetwork struct {
	rq  readQueue
	wqs map[base_info.Server_node]*writeQueue
}

func init() {
	ININP.IINP = &InternalNetwork{}
	ININP.IINP.Init()
}

func (this *InternalNetwork) Init() {
	log.Debugf("Init")
	this.rq.q.Init()
	go this.rq.Handle()
	this.wqs = map[base_info.Server_node]*writeQueue{}
	this.wqs[base_info.Node1] = &writeQueue{}
	this.wqs[base_info.Node1].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Node1])
	this.wqs[base_info.Node2] = &writeQueue{}
	this.wqs[base_info.Node2].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Node2])
	this.wqs[base_info.Clinet] = &writeQueue{}
	this.wqs[base_info.Clinet].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Clinet])
}
func (this *InternalNetwork) SendMsg(pkgInfo *base_info.PkgInfo) {
	this.wqs[base_info.Server_node(pkgInfo.Target)].PushBack(pkgInfo.B)
}
func (this *InternalNetwork) SendsyncMsg(pkgInfo *base_info.PkgInfo) {
	//log.Debugf("readQueue SendMsg, param msg:%s ,target: ", *msg, base_info.GetNodeName(target))
}
func (this *InternalNetwork) Run(address string) {
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
func (this *InternalNetwork) HandleConn(c net.Conn) {
	source, err := this.RegisterNode(c)
	if err != nil {
		c.Close()
		return
	}
	go this.ReadHandle(source, c)
}
func (this *InternalNetwork) ReadHandle(source byte, c net.Conn) {
	defer c.Close()
	for {
		b := make([]byte, base_info.ByteLenth) //base_info.ByteLenth 包的长度
		n, err := c.Read(b)
		if err != nil {
			this.wqs[base_info.Server_node(source)].Close()
			log.Debugf("conn read err:%s", err.Error())
			return
		}
		if n == 0 {
			log.Debugf("conn read msg err: n == 0")
			continue
		}

		log.Debugf("read byte:%v", b)
		switch b[0] {
		case base_info.CommonPkg:
			this.rq.PushBack(&b)
		case base_info.HeartbeatPkg:
		default:
		}
	}
}
func (this *InternalNetwork) RegisterNode(c net.Conn) (byte, error) {
	b := make([]byte, base_info.ByteLenth) //base_info.ByteLenth 包的长度
	_, err := c.Read(b)
	if err != nil {
		log.Debugf("conn read err:%s", err.Error())
		return 0, err
	}
	source := b[1]
	switch b[0] {
	case base_info.RegisterPkg, base_info.RegisterClientPkg:
		log.Debugf("source:%s", base_info.NodeName[base_info.Server_node(source)])
		log.Debugf("map adder:%p", this.wqs[base_info.Server_node(source)])
		this.wqs[base_info.Server_node(source)].Close()
		go func() {
			this.wqs[base_info.Server_node(source)].HandleConn(c, base_info.Server_node(source))
		}()
		this.wqs[base_info.Server_node(source)].PushBack(&b)
		log.Debugf("register node:%s", base_info.NodeName[base_info.Server_node(source)])
	default:
		log.Errorf("pkg type error")
		return 0, fmt.Errorf("pkg type error")
	}
	return source, nil
}
