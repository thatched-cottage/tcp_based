package central_node

import (
	"fmt"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"net"
)

type InternalNetwork struct {
	rq  readQueue
	wqs map[base_info.ServerNode]*writeQueue
}

func init() {
	ININP.IINP = &InternalNetwork{}
	ININP.IINP.Init()
}

func (this *InternalNetwork) Init() {
	log.Debugf("Init")
	this.rq.q.Init()
	go this.rq.Handle()
	this.wqs = map[base_info.ServerNode]*writeQueue{}
	this.wqs[base_info.Node1] = &writeQueue{}
	this.wqs[base_info.Node1].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Node1])
	this.wqs[base_info.Node2] = &writeQueue{}
	this.wqs[base_info.Node2].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Node2])
	this.wqs[base_info.Client] = &writeQueue{}
	this.wqs[base_info.Client].Init()
	log.Debugf("map adder:%p", this.wqs[base_info.Client])
}
func (this *InternalNetwork) SendMsg(pkgInfo *base_info.PkgInfo) {
	this.wqs[pkgInfo.Target].PushBack(pkgInfo.B)
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
func (this *InternalNetwork) ReadHandle(source base_info.ServerNode, c net.Conn) {
	defer c.Close()
	for {
		b := make([]byte, base_info.ByteLength) //base_info.ByteLength 包的长度
		n, err := c.Read(b)
		if err != nil {
			if this.wqs[source] != nil {
				this.wqs[source].Close()
			}
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
func (this *InternalNetwork) RegisterNode(c net.Conn) (base_info.ServerNode, error) {
	b := make([]byte, base_info.ByteLength) //base_info.ByteLength 包的长度
	_, err := c.Read(b)
	if err != nil {
		log.Debugf("conn read err:%s", err.Error())
		return 0, err
	}
	source := base_info.ServerNode(b[1])
	switch b[0] {
	case base_info.RegisterPkg:
		log.Debugf("source:%s", base_info.NodeName[source])
		log.Debugf("map adder:%p", this.wqs[source])
		if this.wqs[source] != nil {
			this.wqs[source].Close()
		}
		go func() {
			this.wqs[source].HandleConn(c, source)
			this.wqs[source].PushBack(&b)
		}()

		log.Debugf("register node:%s", base_info.NodeName[source])
	default:
		log.Errorf("pkg type error")
		return 0, fmt.Errorf("pkg type error")
	}
	return source, nil
}
