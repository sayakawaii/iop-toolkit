/*
# ------------------------------------------------------------
# -- conDefine.go
# --
# -- Huang Minghe
# -- 2022-8-29
# ------------------------------------------------------------
*/
package oltConnect

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type conContext struct {
	User   string
	Passwd string
	Addr   string
	Port   uint64
}

type conProxy struct {
	User   string
	Passwd string
	Addr   string
	Port   uint64
}

type conHandler interface {
	connect(c conContext) (*ssh.Client, error)
	exec(client *ssh.Client, cmd string) ([]byte, error)
}

type conHandlerBase struct{}

func (con *conHandlerBase) connect(c conContext) (*ssh.Client, error) {
	// create ssh configuration
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            c.User,
		Auth:            []ssh.AuthMethod{ssh.Password(c.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个不够安全，生产环境不建议使用
		//HostKeyCallback: ssh.FixedHostKey(), // 建议使用这种，目前还没研究出怎么使用[todo]
	}

	// dial连接服务器
	addr := fmt.Sprintf("%s:%d", c.Addr, c.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("connection faild ", err)
		return nil, err
	}

	//defer sshClient.Close()
	return client, nil
}

func (con *conHandlerBase) exec(client *ssh.Client, cmd string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	cmdInfo, err := session.CombinedOutput(cmd)
	if err != nil {
		return nil, err
	}
	return cmdInfo, nil
}

var (
	conHandlerDef = make(map[uint64]conHandler)
)

func conHandlerRegist(key uint64, handler conHandler) {
	conHandlerDef[key] = handler
}
