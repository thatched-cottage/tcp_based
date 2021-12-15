package main

import (
	"fmt"
	"gitee.com/wuxiansheng/tcp_based.git/pkg"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/base_info"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/command"
	"gitee.com/wuxiansheng/tcp_based.git/pkg/log"
	"net"
)

func main() {
	pkg.Server_name = base_info.Clinet
	for {
		conn, err := net.Dial("tcp", ":8889")
		if err != nil {
			log.Debugf("dial error: %s", err)
		}

		if conn != nil {
			HandleConn(conn)
		}
	}
}

func HandleConn(c net.Conn) {
	if err := RegisterClient(c); err != nil {
		return
	}
	defer c.Close()
	for {
		pkgInfo := &base_info.PkgInfo{}
		pkgInfo.PkgType = base_info.CommonPkg
		pkgInfo.Source = base_info.Clinet
		pkgInfo.ClientId = []byte("1")
		log.Debugf("请输入发往那个服务器:")
	NODE:
		fmt.Scanln(&pkgInfo.Target)
		if base_info.Node1 != pkgInfo.Target &&
			base_info.Node2 != pkgInfo.Target {
			log.Debugf("输入错误！[%d]", pkgInfo.Target)
			log.Debugf("1：node1.")
			log.Debugf("2：node2.")
			log.Debugf("请重新输入那个服务器:")
			goto NODE
		}

		log.Debugf("请输入协议号:")
	COMMAND:
		fmt.Scanln(&pkgInfo.Command)
		if command.CommandMirror != pkgInfo.Command {
			log.Debugf("输入错误！[%d]", pkgInfo.Command)
			log.Debugf("2：Command Mirror.")
			log.Debugf("请重新协议号:")
			goto COMMAND
		}

		log.Debugf("请输入内容:")
	MSG:
		fmt.Scanln(&pkgInfo.Msg)
		if len(pkgInfo.Msg) == 0 {
			log.Debugf("输入错误！请输入合法长度，请输入内容:")
			goto MSG
		}

		b := pkgInfo.PkgInfoToB()
		n, err := c.Write(*b)
		if err != nil {
			log.Debugf("conn Write err:%s", err.Error())
			return
		}
		if n == 0 {
			log.Debugf("conn Write msg err: n == 0")
			continue
		}
		readHandle(c)
	}
}

func readHandle(c net.Conn) {
	log.Debugf("Waiting to receive package")
	b := make([]byte, base_info.ByteLenth)
	_, err := c.Read(b)
	if err != nil {
		log.Debugf("conn read err:%s", err.Error())
		return
	}
	pkgInof := base_info.BToPkgInfo(&b)
	switch pkgInof.PkgType {
	case base_info.CommonPkg:
		log.Debugf("str : %s", string(pkgInof.Msg))
	}
}

func RegisterClient(c net.Conn) error {
	log.Debugf("Register client")

	pkgInfo := &base_info.PkgInfo{}
	pkgInfo.PkgType = base_info.RegisterClientPkg
	pkgInfo.Source = base_info.Clinet
	pkgInfo.Target = base_info.CentralNode
	pkgInfo.ClientId = []byte("1")

	_, err := c.Write(*pkgInfo.PkgInfoToB())
	if nil != err {
		log.Errorf("Register client send msg err:%s ", err.Error())
		return err
	}

	b := make([]byte, base_info.ByteLenth)
	_, err = c.Read(b)
	if nil != err {
		log.Fatalf("Register client send msg err:%s ", err.Error())
		return err
	}
	log.Debugf("Register client success!!!")
	return nil
}
