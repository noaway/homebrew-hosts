package core

import (
	"regexp"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

const (
	defaultUser   = "root"
	defaultPasswd = "MhxzKhl2015"
)

var re = regexp.MustCompile(`"(.*)"`)

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

func GetSuperToken(h *Host) string {
	out, err := sshClient(h).SuperToken()
	procErr(err)
	return out
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
	c := sshClient(&h)
	procErr(c.err)
	cmds := []func(*client) (string, error){
		(*client).DebugboxVersion,
		(*client).SuperToken,
	}
	for _, cmd := range cmds {
		out, err := cmd(c)
		procErr(err)
		logrus.Info(out)
	}
}

func (c *client) DebugboxVersion() (string, error) {
	cmd := "spadmin upgrader version"
	return c.RunCmd(cmd)
}

func (c *client) SuperToken() (string, error) {
	cmd := "spadmin config get global -n super_api_token"
	out, err := c.RunCmd(cmd)
	if err != nil {
		return "", err
	}
	strs := re.FindStringSubmatch(out)
	if len(strs) > 0 {
		return strings.Replace(strs[0], `"`, "", -1), nil
	}
	return "", nil
}
