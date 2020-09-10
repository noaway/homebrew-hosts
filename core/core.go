package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/kevinburke/ssh_config"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	defaultUser   = "root"
	defaultPasswd = "MhxzKhl2015"
)

func GetHosts() []Host {
	hostList := []Host{}
	newHostKeeper(filter).hostRange(func(h Host) bool {
		hostList = append(hostList, h)
		return true
	})
	return hostList
}

func RenderTable() {
	hk := newHostKeeper(filter)
	table := tablewriter.NewWriter(logrus.StandardLogger().Out)
	table.SetHeader([]string{"id", "ip", "host_name"})
	hk.hostRange(func(h Host) bool {
		table.Append([]string{h.ID, h.IP, h.Hostname})
		return true
	})
	table.Render()
}

func filter(h Host) bool { return strings.Contains(h.Hostname, "debugbox") }

func getAuth(hostname string) (string, ssh.AuthMethod) {
	identityPath := ssh_config.Get(hostname, "IdentityFile")
	user := ssh_config.Get(hostname, "User")
	key, err := ioutil.ReadFile(expand(identityPath))
	procErr(err)
	signer, err := ssh.ParsePrivateKey(key)
	procErr(err)
	if identityPath != "" && user != "" {
		return user, ssh.PublicKeys(signer)
	}
	return defaultUser, ssh.Password(defaultPasswd)
}

func Debugbox(id string) {
	hk := newHostKeeper(filter)
	h, ok := hk.get(id)
	if !ok {
		return
	}
	user, auth := getAuth(h.Hostname)
	addr := fmt.Sprintf("%v:22", h.IP)
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{auth},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	procErr(err)
	defer client.Close()
	session, err := client.NewSession()
	procErr(err)
	defer session.Close()
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	procErr(err)
	defer func() { _ = terminal.Restore(fd, oldState) }()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	procErr(err)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	procErr(session.RequestPty("xterm", termHeight, termWidth, modes))
	procErr(session.Run("/bin/bash"))
}
