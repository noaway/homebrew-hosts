package main

import (
	"fmt"
	"os"

	"github.com/noaway/celery/core"
	"github.com/noaway/celery/postman"
	"github.com/spf13/cobra"
)

func cmdGenPostmanEnv() *cobra.Command {
	var out = ""
	cmd := &cobra.Command{
		Use:   "gen_postman_env",
		Short: "run gen_postman_env",
		Run: func(_ *cobra.Command, args []string) {
			hosts := core.GetHosts()
			for _, host := range hosts {
				env := postman.PostmanEnv{
					Name: host.Hostname,
					Values: []postman.PostmanValue{
						{
							Key:     "host",
							Value:   fmt.Sprintf("%v:8107", host.IP),
							Enabled: true,
						},
						{
							Key:     "token",
							Value:   "",
							Enabled: true,
						},
						{
							Key:     "data_host",
							Value:   fmt.Sprintf("%v:8106", host.IP),
							Enabled: true,
						},
					},
				}
				if out == "" {
					fmt.Println(env.JsonString())
					continue
				}
				f, err := os.Create(fmt.Sprintf("%v/%v.json", out, env.Name))
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer f.Close()
				bytes, err := env.JsonBytes()
				if err != nil {
					fmt.Println(err)
					continue
				}
				_, _ = f.Write(bytes)
			}
		},
	}

	cmd.Flags().StringVarP(&out, "out", "o", "", "out")
	return cmd
}

func hosts(_ *cobra.Command, args []string) {
	core.RenderTable()
}

func cmdExec() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "run exec",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) > 0 {
				core.Debugbox(args[0])
			}
		},
	}
	return cmd
}

func main() {
	cmdRoot := &cobra.Command{Use: "hosts", Run: hosts, Version: "0.1.3"}
	cmdRoot.AddCommand(cmdGenPostmanEnv())
	cmdRoot.AddCommand(cmdExec())
	_ = cmdRoot.Execute()
}
