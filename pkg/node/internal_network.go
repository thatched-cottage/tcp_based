package node

import (
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"net"
)

type InternalNetwork struct {
	rq readQueue
	wq writeQueue
}

func init() {
	ININP.IINP = &InternalNetwork{}
	ININP.IINP.Init()
}

func (this *InternalNetwork) Init() {
	log.Debugf("readQueue init")
	this.rq.q.Init()
	go this.rq.Handle()
	this.wq.q.Init()
	this.wq.close = make(chan interface{}, 1)
}
func (this *InternalNetwork) SendMsg(pkgInfo *base_info.PkgInfo) {
	sendData := pkgInfo.PkgInfoToB()
	this.wq.PushBack(sendData)
}
func (this *InternalNetwork) SendsyncMsg(pkgInfo *base_info.PkgInfo) {
	sendData := pkgInfo.PkgInfoToB()
	this.wq.PushBack(sendData)
}
func (this *InternalNetwork) Run(address string) {
	for {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Debugf("dial error: %s", err)
			continue
		}
		log.Debugf("connect to server ok")
		if conn != nil {
			this.HandleConn(conn)
		}
	}
}
func (this *InternalNetwork) HandleConn(c net.Conn) {
	if err := this.RegisterNode(c); err != nil {
		return
	}
	go this.ReadHandle(c)
	this.wq.writeHandle(c)
}
func (this *InternalNetwork) ReadHandle(c net.Conn) {
	defer c.Close()
	for {
		b := make([]byte, base_info.ByteLenth) //base_info.ByteLenth 包的长度
		n, err := c.Read(b)
		if err != nil {
			this.wq.Close()
			log.Errorf("conn read err:%s", err.Error())
			return
		}
		if n == 0 {
			log.Errorf("conn read msg err: n == 0")
			continue
		}
		log.Debugf("read byte:%v", b)
		switch b[0] {
		case base_info.CommonPkg:
			this.rq.PushBack(&b)
		case base_info.HeartbeatPkg:
		}
	}
}
func (this *InternalNetwork) RegisterNode(c net.Conn) error {
	b := []byte{base_info.RegisterPkg,
		byte(pkg.Server_name),
		byte(base_info.CentralNode)}

	_, err := c.Write(b)
	if err != nil {
		log.Errorf("conn Write err:", err.Error())
		return err
	}
	b = make([]byte, base_info.ByteLenth) //base_info.ByteLenth 包的长度
	_, err = c.Read(b)
	if err != nil {
		log.Errorf("conn Write err:", err.Error())
		return err
	}
	pkgtype := (b)[0]
	switch pkgtype {
	case base_info.RegisterPkg:
		log.Debugf("register node ack")
		log.Debugf("register node success")
	}
	return nil
}
