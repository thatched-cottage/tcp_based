package base_info

type ServerNode byte

const (
	CentralNode ServerNode = iota
	Node1
	Node2
	Client
)

const ByteLength = 1024 // payload length

const (
	RegisterClientPkg byte = iota
	CommonPkg
	HeartbeatPkg
	RegisterPkg
)

var NodeName = map[ServerNode]string{
	CentralNode: `中心节点`,
	Node1:       `节点1`,
	Node2:       `节点2`,
	Client:      `客户端`,
}
