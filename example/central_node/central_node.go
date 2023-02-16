package main

import (
	"github.com/thatched-cottage/tcp_based.git/pkg"
	"github.com/thatched-cottage/tcp_based.git/pkg/ININP"
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
	_ "github.com/thatched-cottage/tcp_based.git/pkg/central_node"
	"time"
)

func main() {
	pkg.Server_name = base_info.CentralNode
	go ININP.IINP.Run(":8888")
	go ININP.IENP.Run(":8889")

	for true {
		time.Sleep(time.Second * 100)
	}
}
