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
	"omciAnalyzer/dao"
	"testing"
)

// func dropCR(data []byte) []byte {
// 	if len(data) > 0 && data[len(data)-1] == '\r' {
// 		return data[0 : len(data)-1]
// 	}
// 	return data
// }

func TestOLTConnection(t *testing.T) {
	// c := new(conContext)
	// c.User = "root"
	// c.Passwd = "2x2=4"
	// c.Addr = "135.251.192.132"
	// c.Port = 923

	// handler := conHandlerDef[1]

	// fmt.Println("connect to: " + c.Addr)
	// client, err := handler.connect(*c)
	// if err != nil {
	// 	return
	// }
	// defer client.Close()

	// session, err := client.NewSession()
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// modes := ssh.TerminalModes{
	// 	ssh.ECHO:          0,     // disable echoing
	// 	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	// 	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	// }

	var con dao.CollectorConnectionInfo
	con.Board = "FGLT-B"
	con.Platform = "LS"
	con.IP = "10.99.76.132"
	con.Port = 923
	con.User = "root"
	con.Password = "2x2=4"
	go Connect(con)
	TestCmd()
	t.Log("Test OLTC onnection end")
}
