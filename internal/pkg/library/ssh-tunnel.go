package library

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var (
	Username = "elim"
	Password = "permisiGan"
	Host     = "103.24.51.13"
	Port     = "61122"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(address string) (net.Conn, error) {
	return self.client.Dial("tcp", address)
}

func ConnectSSHTunnel() (sshcon *ssh.Client, dialer *ViaSSHDialer, err error) {
	// ssh tunner
	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            Username,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	// if sshPass != "" {
	sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
		return Password, nil
	}))
	// }

	sshcon, errSSH := ssh.Dial("tcp", fmt.Sprintf("%s:%s", Host, Port), sshConfig)
	if errSSH != nil {
		return nil, nil, errSSH
	} else {
		fmt.Println("SSH Connected")
	}

	dialer = &ViaSSHDialer{sshcon}
	return
}
