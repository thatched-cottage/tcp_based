package base_info

import (
	"encoding/binary"
	"github.com/thatched-cottage/tcp_based.git/pkg/command"
	"github.com/thatched-cottage/tcp_based.git/pkg/log"
)

// 对于包的封装，服务器与客户端之前是需要一种协议确定是否可以连接。
type PkgInfo struct {
	// Pkg 的类型
	PkgType byte
	// 发往的目标，确定发到哪一个目标，central node 会根据Target 确定你发往的目标，进行转发。
	Target ServerNode
	// 自己是哪个Node节点，标识自己的属于那个Node方便接收方回包
	Source ServerNode

	// 如果是client发的信息，需要携带ClientId。
	ClientId []byte

	// 携带的信息，command必须需要是在 command 定义过的 ,msg 是具体的信息
	Command command.CommandType
	Msg     []byte
	// B 是数据的byte格式
	B *[]byte
}

func (this *PkgInfo) PkgInfoTargetNodeName() string {
	return NodeName[this.Target]
}
func (this *PkgInfo) PkgInfoSourceNodeName() string {
	return NodeName[this.Source]
}

func BToPkgInfo(b *[]byte) *PkgInfo {
	log.Debugf("byte:%v", *b)
	pkgInfo := &PkgInfo{}
	pkgInfo.B = b
	offset := uint32(0)
	pkgInfo.PkgType = (*b)[offset]
	offset++
	pkgInfo.Source = ServerNode((*b)[offset])
	offset++
	pkgInfo.Target = ServerNode((*b)[offset])
	offset++
	if (*b)[offset] > 0 { // clientId 长度
		ClientIdLength := uint32((*b)[offset])
		offset++
		pkgInfo.ClientId = (*b)[offset : offset+ClientIdLength]
		offset += ClientIdLength
		log.Debugf("this.ClientId:%s ,Sizeof:%d", pkgInfo.ClientId, ClientIdLength)
	} else {
		offset++
	}
	pkgInfo.Command = command.CommandType((*b)[offset])
	log.Debugf("this.Command:%d ", pkgInfo.Command)
	offset++
	MsgLength := binary.BigEndian.Uint32((*b)[offset:])
	offset += 4
	if MsgLength > 0 { // 消息 长度
		log.Debugf("MsgLength:%d", MsgLength)
		pkgInfo.Msg = (*b)[offset : offset+MsgLength]
		offset += MsgLength
	}
	log.Debugf("pkgInfo :%+v", pkgInfo)
	return pkgInfo
}

func (this *PkgInfo) PkgInfoToB() *[]byte {
	log.Debugf("pkgInfo :%+v", this)
	b := &[]byte{}
	*b = append(*b, this.PkgType)      // 0
	*b = append(*b, byte(this.Source)) // 1
	*b = append(*b, byte(this.Target)) // 2
	log.Debugf("this.ClientId:%s ,Sizeof:%d", this.ClientId, byte(len(this.ClientId)))
	*b = append(*b, byte(len(this.ClientId))) // 3
	*b = append(*b, this.ClientId...)         // 4
	*b = append(*b, byte(this.Command))       // 5
	log.Debugf("this.Command:%d ", this.Command)
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, uint32(len(this.Msg)))
	*b = append(*b, b4...)       //6-9
	*b = append(*b, this.Msg...) //10-n
	log.Debugf("byte:%v", *b)
	return b
}
