package core

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	EtcHostsPath = "/etc/hosts"
)

type Host struct {
	ID       string
	IP       string
	Hostname string
	Alias    string
}

func newHostKeeper(fn func(h Host) bool) *hostKeeper {
	hk := &hostKeeper{bucket: make(map[string]*list.Element), list: list.New(), filterFunc: fn}
	if err := hk.readHosts(); err != nil {
		panic(err)
	}
	return hk
}

type hostKeeper struct {
	bucket     map[string]*list.Element
	list       *list.List
	filterFunc func(h Host) bool
}

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
		if !filter(h) {
			continue loop
		}
		h.ID = md5HashBytes([]byte(fmt.Sprintf("%v%v", h.IP, h.Hostname)))[:8]
		hk.push(h.ID, h)
	}
	return nil
}
