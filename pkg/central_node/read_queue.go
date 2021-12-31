package central_node

import (
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"gitee.com/wuxiansheng/tcp_based.git/utils"
)

type readQueue struct {
	q utils.Queue
}

func (this *readQueue) PushBack(data interface{}) {
	this.q.PushBack(data)
}

func (this *readQueue) Handle() {
	for {
		req := <-this.q.Pop()
		b := req.(*[]byte)
		pkgInfo := base_info.BToPkgInfo(b)
		log.Debugf("send Msg %s", pkgInfo.PkgInfoTargetNodeName())
		switch pkgInfo.Target {
		case base_info.Node1, base_info.Node2:
			ININP.IINP.SendMsg(pkgInfo)
		case base_info.Client:
			ININP.IENP.SendMsg(pkgInfo)
		}
	}
}
