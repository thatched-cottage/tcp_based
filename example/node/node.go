package main

import (
	"flag"
	"github.com/thatched-cottage/tcp_based.git/pkg"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
	_ "github.com/thatched-cottage/tcp_based.git/pkg/node"
)

var nodename = flag.Int("nodename", 0, "node name")

func main() {
	flag.Parse()

	pkg.ServerName = base_info.ServerNode(*nodename)
	log.Debugf("这是%s", base_info.NodeName[pkg.ServerName])
	ININP.IINP.Run(":8888")
}
