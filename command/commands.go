package command

import (
	"github.com/bill-rich/consul-sm-fed-svc/command/agent"
	"github.com/mitchellh/cli"
)

func init() {
	Register("agent", func(ui cli.Ui) (cli.Command, error) { return agent.New(ui), nil })
}