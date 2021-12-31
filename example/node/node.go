package main

import (
	"flag"
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	_ "gitee.com/wuxiansheng/tcp_based.git/pkg/node"
)
var nodename = flag.Int("nodename", 0, "node name")

func main() {
	flag.Parse()
	pkg.Server_name = base_info.ServerNode(*nodename)
	log.Debugf("这是%s", base_info.NodeName[pkg.Server_name])
	ININP.IINP.Run(":8888")
}
