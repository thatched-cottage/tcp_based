package node

import (
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"gitee.com/wuxiansheng/tcp_based.git/utils"
	"net"
)

type writeQueue struct {
	q     utils.Queue
	c     *net.Conn
	close chan interface{}
}

func (this *writeQueue) PushBack(data interface{}) {
	this.q.PushBack(data)
}

func (this *writeQueue) Pop() chan interface{} {
	return this.q.Pop()
}

func (this *writeQueue) Close() {
	if this.c == nil {
		return
	}
	this.close <- true
}

func (this *writeQueue) writeHandle(c net.Conn) {
	this.c = &c
	defer c.Close()
	for {
		select {
		case <-this.close:
			log.Debugf("ctx Done:")
			this.c = nil
			return
		case i := <-this.Pop():
			log.Debugf("send Msg %s ", base_info.NodeName[(pkg.Server_name)])
			b := i.(*[]byte)
			n, err := c.Write(*b)
			if err != nil {
				log.Errorf("conn Write err:", err.Error())
				return
			}
			if n == 0 {
				log.Errorf("conn Write msg err: n == 0")
				continue
			}
		}
	}
}
