package main

import (
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/ININP"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	_ "gitee.com/wuxiansheng/tcp_based.git/pkg/node"
)

func main() {
	pkg.Server_name = base_info.Node1
	ININP.IINP.Run(":8888")
}
