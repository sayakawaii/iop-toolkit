/*
# ------------------------------------------------------------
# -- oltCon_test.go
# --
# -- Huang Minghe
# -- 2022-8-29
# ------------------------------------------------------------
*/
package oltConnect

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"omciAnalyzer/dao"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
)

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

var inputChan = make(chan []byte)
var outputChan = make(chan string)
var ctrlChan = make(chan os.Signal)

func Connect(con dao.CollectorConnectionInfo) {

	fmt.Println("eshell")
	// config := &ssh.ClientConfig{
	// 	User: "root",
	// 	Auth: []ssh.AuthMethod{
	// 		ssh.Password("2x2=4"),
	// 	},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	// client, err := ssh.Dial("tcp", "135.251.192.132:923", config)
	// if err != nil {
	// 	log.Fatal("Failed to dial: ", err)
	// }
	config := &ssh.ClientConfig{
		User: con.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(con.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", con.IP, con.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	//ctrl + c
	signal.Notify(ctrlChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	//input pipe
	//session.Stdin = os.Stdin
	// inputChan := make(chan []byte)
	resWriter, _ := session.StdinPipe()
	// go func() {
	// 	for {
	// 		reader := bufio.NewReaderSize(os.Stdin, 65535) // NewReader 默认4096
	// 		data, _, _ := reader.ReadLine()
	// 		inputChan <- append(data, '\n')
	// 	}
	// }()

	go func() {
		for {
			select {
			case <-ctrlChan:
				resWriter.Write([]byte{3, '\n'})
			case data := <-inputChan:
				resWriter.Write(data)
			}
		}
	}()

	//output pipe
	//session.Stdout = os.Stdout
	resReader, _ := session.StdoutPipe()
	scanner := bufio.NewScanner(resReader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, dropCR(data[0:i]), nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}

		// "#"
		if data[len(data)-1] == byte(32) &&
			data[len(data)-2] == byte(35) &&
			data[len(data)-3] == byte(32) {
			return len(data), data, nil
		}

		// Request more data.
		return 0, nil, nil
	})

	go func() {
		for scanner.Scan() {
			if len(scanner.Bytes()) > 2 &&
				scanner.Bytes()[len(scanner.Bytes())-1] == byte(32) &&
				scanner.Bytes()[len(scanner.Bytes())-2] == byte(35) &&
				scanner.Bytes()[len(scanner.Bytes())-3] == byte(32) {
				// fmt.Print(scanner.Text())
				outputChan <- scanner.Text()
			} else {
				// fmt.Println(scanner.Text())
				outputChan <- scanner.Text() + "\n"
			}
		}
	}()

	// start shell
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}
	// fmt.Println(session.Wait())
	for range ctrlChan {
		close(inputChan)
		close(outputChan)
		close(ctrlChan)
		fmt.Println("connection end")
		return
	}
}

func EndConnection() {
	ctrlChan <- syscall.SIGTERM
	time.Sleep(100 * time.Millisecond)
}

func runCmd(cmd []byte, d time.Duration) string {
	// cnt := 0
	var str string
	//start timer
	timer := time.NewTimer(d)
	inputChan <- append(cmd, '\n')
	// return ""
	// for data := range outputChan {
	// 	cnt++
	// 	str += data
	// 	if cnt == 1 {
	// 		//reset timer as 50ms
	// 		timer.Reset(50 * time.Millisecond)
	// 	}
	// }
	for {
		select {
		case data := <-outputChan:
			// case str +=<-outputChan:
			// cnt++
			str += data
			// if cnt == 1 {
			//reset timer as 50ms
			// timer.Stop()
			// timer.Reset(50 * time.Millisecond)
			// }
		case <-timer.C:
			return str
		}
	}
}

func GetSlotInfo() map[uint64]string {
	fmt.Println(runCmd([]byte("pwd"), 500*time.Millisecond))
	res := runCmd([]byte("confd_cli -u techsupport"), 1000*time.Millisecond)
	// fmt.Println(res)
	res = runCmd([]byte("show hardware-state component model-name"), 1000*time.Millisecond)
	// fmt.Println("res:", res)
	if len(res) != 0 {
		var slotInfo = make(map[uint64]string)
		//get one line
		startIndex := 0
		endIndex := 0
		var tmp string
		var str string
		for {
			tmp = res[startIndex:]
			// fmt.Println("tmp:", tmp)
			endIndex = strings.Index(tmp, "\n")
			// fmt.Println("endIndex:", endIndex)
			str = tmp[0:endIndex]
			//get olt slot info
			for i := uint64(1); i < 4; i++ {
				keyStr := "Slot-" + strconv.FormatUint(i, 10) + "_"
				// fmt.Println("str:", str)
				// fmt.Println("keyStr", keyStr)
				if strings.Contains(str, keyStr) { //gpon log
					leftIndex := strings.Index(str, keyStr)
					rightIndex := strings.Index(str, " ")
					slot := str[leftIndex:rightIndex]
					if _, ret := slotInfo[i]; !ret {
						slotInfo[i] = slot
					}
					break
				}
			}
			startIndex += endIndex + 1
			if startIndex >= len(res) {
				break
			}
		}
		runCmd([]byte("exit"), 100*time.Millisecond)
		return slotInfo
	}
	runCmd([]byte("exit"), 100*time.Millisecond)
	return nil
}

var slotIPDef map[uint64]string = map[uint64]string{
	1: "169.254.1.3",
	2: "169.254.1.4",
	3: "169.254.1.5",
	4: "169.254.1.6",
}

func GetSlotIP(slotIndex uint64) string {
	return slotIPDef[slotIndex]
}

func TestCmd() {

	// time.Sleep(1 * 1000 * 1000 * 1000)
	// inputChan <- append([]byte("pwd"), '\n')
	fmt.Println(runCmd([]byte("pwd"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	//login LT
	// inputChan <- append([]byte("ssh -p 2222 root@169.254.1.3"), '\n')
	fmt.Println(runCmd([]byte("ssh -p 2222 root@169.254.1.6"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	// inputChan <- append([]byte("2x2=4"), '\n')
	fmt.Println(runCmd([]byte("2x2=4"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	//login app
	// inputChan <- append([]byte("pwd"), '\n')
	fmt.Println(runCmd([]byte("pwd"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	// inputChan <- append([]byte("cd /isam/slot_default/onumgnt_hypervisor_app/run/"), '\n')
	fmt.Println(runCmd([]byte("cd /isam/slot_default/onumgnt_hypervisor_app/run/"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	// inputChan <- append([]byte("/isam/user/calamares -cnt tty1_ext"), '\n')
	fmt.Println(runCmd([]byte("/isam/user/calamares -cnt tty1_ext"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	//get data
	// inputChan <- append([]byte("onumgnt showOnuData summary"), '\n')
	fmt.Println(runCmd([]byte("onumgnt showOnuData summary"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	//exit app
	// inputChan <- []byte{0x3, 0xd}
	fmt.Println(runCmd([]byte{0x3, 0xd}, 500*time.Millisecond))
	//exit LT
	// inputChan <- append([]byte("exit"), '\n')
	fmt.Println(runCmd([]byte("exit"), 500*time.Millisecond))
	//exit NT
	// inputChan <- append([]byte("exit"), '\n')
	fmt.Println(runCmd([]byte("exit"), 500*time.Millisecond))
	// time.Sleep(1 * 1000 * 1000 * 1000)
	//
	ctrlChan <- syscall.SIGTERM
	time.Sleep(500 * time.Millisecond)
	fmt.Println("TestCmd end")
}
