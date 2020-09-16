package core

import (
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

const (
	defaultUser   = "root"
	defaultPasswd = "MhxzKhl2015"
)

func GetHosts() []Host {
	return loadHosts(debugboxFilter())
}

func RenderTable() {
	hk := newHostKeeper(debugboxFilter())
	table := tablewriter.NewWriter(logrus.StandardLogger().Out)
	table.SetHeader([]string{"id", "ip", "host_name"})
	hk.hostRange(func(h Host) bool {
		table.Append([]string{h.ID, h.IP, h.Hostname})
		return true
	})
	table.Render()
}

func debugboxFilter() filterFunc {
	return func(h Host) bool { return strings.Contains(h.Hostname, "debugbox") }
}

func strFilter(str string) filterFunc {
	return func(h Host) bool {
		prefix := strings.Contains(h.Hostname, "debugbox")
		switch {
		case StrTo(str).MustInt() != 0:
			return strings.Contains(h.Hostname, str) && prefix
		default:
			return h.ID == str && prefix
		}
	}
}

func Debugbox(str string) {
	h, ok := loadHost(strFilter(str))
	if !ok {
		return
	}
	procErr(sshClient(&h).Interaction().err)
}

func DebugboxStatus(str string) {
	h, ok := loadHost(strFilter(str))
	if !ok {
		return
	}
	client := sshClient(&h)
	procErr(client.err)
	cmds := []string{
		"spadmin upgrader version",
		"spadmin config get global -n super_api_token",
	}
	for _, cmd := range cmds {
		out, err := client.RunCmd(cmd)
		procErr(err)
		logrus.Info(out)
	}
}
