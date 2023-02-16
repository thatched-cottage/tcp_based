package main

import (
	"github.com/thatched-cottage/tcp_based.git/pkg"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	_ "github.com/thatched-cottage/tcp_based.git/pkg/node"
)

func main() {
	pkg.Server_name = base_info.Node2
	ININP.IINP.Run(":8888")
}
