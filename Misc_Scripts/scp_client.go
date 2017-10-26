package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`

func main() {
	signer, _ := ssh.ParsePrivateKey([]byte(privateKey))
	clientConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "192.243.118.65:29138", clientConfig)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		content := "123456789\n"
		fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
		fmt.Fprintln(w, "C0644", len(content), "testfile1")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00") // transfer end with \x00
		fmt.Fprintln(w, "C0644", len(content), "testfile2")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")
	}()
	if err := session.Run("scp -tr ./"); err != nil {
		panic("Failed to run: " + err.Error())
	}
}

