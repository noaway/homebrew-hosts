package core

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// sshClient func
func sshClient(h *Host) *client {
	if err := h.getAuth(); err != nil {
		return &client{err: err}
	}
	c, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", h.IP), &ssh.ClientConfig{
		User:            h.auth.user,
		Auth:            []ssh.AuthMethod{h.auth.method},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	return &client{c, err}
}

type client struct {
	*ssh.Client
	err error
}

func (c *client) Interaction() *client {
	if c.err != nil {
		return c
	}
	session, err := c.NewSession()
	if err != nil {
		c.err = err
		return c
	}
	defer session.Close()
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		c.err = err
		return c
	}
	defer func() { _ = terminal.Restore(fd, oldState) }()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		c.err = err
		return c
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	procErr(session.RequestPty("xterm", termHeight, termWidth, modes))
	procErr(session.Run("/bin/bash"))
	return c
}
