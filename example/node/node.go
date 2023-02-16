package main

import (
	"flag"
<<<<<<< HEAD
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	_ "gitee.com/wuxiansheng/tcp_based.git/pkg/node"
)
=======
	"github.com/thatched-cottage/tcp_based.git/pkg"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	_ "github.com/thatched-cottage/tcp_based.git/pkg/node"
)

>>>>>>> c43ea45 (feat: 开始优化)
var nodename = flag.Int("nodename", 0, "node name")

func main() {
	flag.Parse()
<<<<<<< HEAD
	pkg.Server_name = base_info.ServerNode(*nodename)
	log.Debugf("这是%s", base_info.NodeName[pkg.Server_name])
=======
	pkg.ServerName = base_info.ServerNode(*nodename)
	log.Debugf("这是%s", base_info.NodeName[pkg.ServerName])
>>>>>>> c43ea45 (feat: 开始优化)
	ININP.IINP.Run(":8888")
}
