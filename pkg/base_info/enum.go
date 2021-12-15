package base_info

type Server_node byte

const (
	CentralNode Server_node = iota
	Node1
	Node2
	Clinet
)

const ByteLenth = 1024 // 包最大

const (
	RegisterClientPkg byte = iota
	CommonPkg
	HeartbeatPkg
	RegisterPkg
)

var NodeName = map[Server_node]string{
	CentralNode: `中心节点`,
	Node1:       `节点1`,
	Node2:       `节点2`,
	Clinet:      `客户端`,
}
