package core

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
)

const (
	EtcHostsPath = "/etc/hosts"
)

type filterFunc func(Host) bool

type Host struct {
	ID       string
	IP       string
	Hostname string
	Alias    string
	auth     *auth
}

func (h *Host) getAuth() error {
	identityPath := ssh_config.Get(h.Hostname, "IdentityFile")
	user := ssh_config.Get(h.Hostname, "User")
	switch {
	case identityPath != "" && user != "":
		key, err := ioutil.ReadFile(expand(identityPath))
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return err
		}
		h.auth = &auth{user, ssh.PublicKeys(signer)}
		return nil
	default:
		h.auth = &auth{defaultUser, ssh.Password(defaultPasswd)}
	}
	return nil
}

// auth struct
type auth struct {
	user   string
	method ssh.AuthMethod
}

func newHostKeeper(fn filterFunc) *hostKeeper {
	hk := &hostKeeper{bucket: make(map[string]*list.Element), list: list.New(), filter: fn}
	if err := hk.readHosts(); err != nil {
		panic(err)
	}
	return hk
}

func loadHosts(fn filterFunc) []Host {
	hk := newHostKeeper(fn)
	list := make([]Host, 0, hk.len())
	hk.hostRange(func(h Host) bool {
		list = append(list, h)
		return true
	})
	return list
}

func loadHost(fn filterFunc) (Host, bool) {
	hk := newHostKeeper(fn)
	list := make([]Host, 0, hk.len())
	hk.hostRange(func(h Host) bool {
		list = append(list, h)
		return true
	})
	if len(list) > 0 {
		return list[0], true
	}
	return Host{}, false
}

type hostKeeper struct {
	bucket map[string]*list.Element
	list   *list.List
	filter filterFunc
}

func (hk *hostKeeper) len() int { return hk.list.Len() }

func (hk *hostKeeper) get(key string) (Host, bool) {
	if elem, ok := hk.bucket[key]; ok {
		return elem.Value.(Host), true
	} else {
		return Host{}, false
	}
}

func (hk *hostKeeper) push(key string, val Host) {
	if _, ok := hk.bucket[key]; ok {
		return
	}
	elem := hk.list.PushBack(val)
	hk.bucket[key] = elem
}

func (hk *hostKeeper) hostRange(block func(Host) bool) {
	for elem := hk.list.Front(); elem != nil; elem = elem.Next() {
		if block(elem.Value.(Host)) {
			continue
		} else {
			return
		}
	}
}

func (hk *hostKeeper) readHosts() error {
	fp, err := os.Open(EtcHostsPath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
loop:
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "#"); i > -1 {
			line = line[:i]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		if net.ParseIP(fields[0]) == nil {
			log.Printf("invalid IP. [address='%v',path='%v']", fields[0], EtcHostsPath)
			continue
		}

		h := Host{}
		for i, s := range fields {
			if s == "broadcasthost" || s == "localhost" {
				continue loop
			}
			switch i {
			case 0:
				h.IP = s
			case 1:
				h.Hostname = s
			case 2:
				h.Alias = s
			}
		}
		h.ID = md5HashBytes([]byte(fmt.Sprintf("%v%v", h.IP, h.Hostname)))[:8]
		if hk.filter != nil && !hk.filter(h) {
			continue loop
		}
		hk.push(h.ID, h)
	}
	return nil
}
