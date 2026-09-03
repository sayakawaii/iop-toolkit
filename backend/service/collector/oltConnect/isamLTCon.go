/*
# ------------------------------------------------------------
# -- isamLTCon.go
# --
# -- Huang Minghe
# -- 2022-8-29
# ------------------------------------------------------------
*/
package oltConnect

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	IsamLT = 2
)

type conIsamLT struct {
	proxy conProxy
	c     conContext
}

func (con *conIsamLT) connect(c conContext) (*ssh.Client, error) {
	con.c = c
	fmt.Println("proxy: ", con.proxy)
	// create ssh configuration
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            con.proxy.User,
		Auth:            []ssh.AuthMethod{ssh.Password(con.proxy.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个不够安全，生产环境不建议使用
		//HostKeyCallback: ssh.FixedHostKey(), // 建议使用这种，目前还没研究出怎么使用[todo]
	}

	// dial连接服务器
	addr := fmt.Sprintf("%s:%d", con.proxy.Addr, con.proxy.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("connection faild ", err)
		return nil, err
	}

	//defer sshClient.Close()
	return client, nil
}

func (con *conIsamLT) exec(client *ssh.Client, cmd string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	cmdStr := "octopus STDIO " + con.c.Addr + ":udp:" + strconv.FormatUint(con.c.Port, 10)
	fmt.Println("cmdStr: ", cmdStr)
	cmdInfo, err := session.CombinedOutput(cmdStr)
	if err != nil {
		panic(err)
	}
	err = session.Signal(ssh.SIGPIPE)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	cmdStr = "shell"
	cmdInfo, err = session.CombinedOutput(cmdStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	cmdStr = "nt"
	cmdInfo, err = session.CombinedOutput(cmdStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	cmdInfo, err = session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	cmdStr = "exit"
	cmdInfo, err = session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	err = session.Signal(ssh.SIGQUIT)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	cmdStr = "q"
	cmdInfo, err = session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
	return cmdInfo, nil
}

func init() {
	con := new(conIsamLT)
	con.proxy = conProxy{User: "mingheh", Passwd: "UIOP_nsb!@#$1234", Addr: "135.251.206.244", Port: 22}
	conHandlerRegist(IsamLT, con)
}
