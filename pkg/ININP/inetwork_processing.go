package ININP

import (
	"github.com/thatched-cottage/tcp_based.git/pkg/base_info"
)

type INetworkProcessing interface {
	SendMsg(pkgInfo *base_info.PkgInfo)
	SendsyncMsg(pkgInfo *base_info.PkgInfo)
	Run(address string)
	Init()
}

var IINP INetworkProcessing // 内部
var IENP INetworkProcessing // 外部
