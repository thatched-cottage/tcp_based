package node

import (
	"github.com/thatched-cottage/tcp_based.git/pkg"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	"github.com/thatched-cottage/tcp_based.git/utils"
)

// 包的数据组成
// byte |source|target|msg
type readQueue struct {
	q utils.Queue
}

func (this *readQueue) PushBack(data interface{}) {
	log.Debugf("readQueue PushBack")
	this.q.PushBack(data)
}

func (this *readQueue) Handle() {
	log.Debugf("readQueue rqQueueHandle")
	for {
		i := <-this.q.Pop()
		pkgInfo := base_info.BToPkgInfo(i.(*[]byte))
		res, err := pkg.Decode(pkgInfo)
		if err != nil {
			log.Errorf("err:%v", err.Error())
			return
		}
		pkgInfo.Msg = *res
		pkgInfo.Source, pkgInfo.Target = pkgInfo.Target, pkgInfo.Source
		ININP.IINP.SendMsg(pkgInfo)
	}
}
