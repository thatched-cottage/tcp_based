package central_node

import (
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"gitee.com/wuxiansheng/tcp_based.git/utils"
	"net"
)

type writeQueue struct {
	q utils.Queue
	base_info.Server_node
	c     *net.Conn
	close chan interface{}
}

func (this *writeQueue) Init() {
	log.Debugf("writeQueue Init")
	this.q.Init()
	this.close = make(chan interface{}, 1)
	log.Debugf("this.q adder:%p", &this.q)
}

func (this *writeQueue) PushBack(data interface{}) {
	log.Debugf("writeQueue PushBack")
	log.Debugf("this.q adder:%p", &this.q)
	this.q.PushBack(data)
}

func (this *writeQueue) Pop() chan interface{} {
	log.Debugf("writeQueue Pop")
	log.Debugf("this.q adder:%p", &this.q)
	return this.q.Pop()
}

func (this *writeQueue) Close() {
	if this.c == nil {
		return
	}
	this.close <- true
}

func (this *writeQueue) HandleConn(c net.Conn, node base_info.Server_node) {
	log.Debugf("HandleConn")
	this.Server_node = node
	this.c = &c
	defer (*this.c).Close()
	for {
		select {
		case <-this.close:
			log.Debugf("ctx Done:")
			this.c = nil
			return
		case i := <-this.Pop():
			log.Debugf("this is %s write queue", base_info.NodeName[(this.Server_node)])
			b := i.(*[]byte)
			n, err := c.Write(*b)
			if err != nil {
				log.Debugf("conn read err:", err.Error())
				return
			}
			if n == 0 {
				log.Debugf("conn read msg err: n == 0")
				continue
			}
		}
	}
	log.Debugf("node close:", base_info.NodeName[(this.Server_node)])
}
