package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

const key = `------RSA PRIVATE KEY-----`

func main() {
	ce := func(err error, msg string) { //在main 函数中定义匿名函数并使用的方法。
		if err != nil {
			log.Fatalf("%s error:%v", msg, err)
		}
	}

	signer, _ := ssh.ParsePrivateKey([]byte(key))

	client, err := ssh.Dial("tcp", "192.243.118.65:29138", &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	ce(err, "dial")

	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", 25, 80, modes)

	ce(err, "request pty")
	err = session.Shell()
	ce(err, "start shell")
	err = session.Wait()
	ce(err, "return")

}

