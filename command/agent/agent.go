package agent

import (
	"flag"
	"fmt"
	"github.com/bill-rich/consul-sm-fed-svc/listener"
	"github.com/bill-rich/consul-sm-fed-svc/resource_listener"
	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
	"time"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

func (c *cmd) init() {
	// Fill in flags
}

type cmd struct{
	UI    cli.Ui
	flags *flag.FlagSet
	http  *flags.HTTPFlags
	help  string

	authMethodName string
	description    string
	selector       string
	bindType       string
	bindName       string

	showMeta bool
}

func (c *cmd) Run(args []string) int {
	return c.run(args)
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

func (c *cmd) run(args []string) int {
	// c.flags.Parse(args)
	rootCACerts := []string{}
	caCert := "/home/hrich/go/src/github.com/mesh-federation/fsd-server-example/cert.pem"
	peerCert := "/home/hrich/go/src/github.com/mesh-federation/fsd-server-example/cert.pem"
	peerKey := "/home/hrich/go/src/github.com/mesh-federation/fsd-server-example/key.pem"
	serverAddr := ":8001"
	serverPort := uint32(8001)
	insecureSkipVerify := true

	ch := make(chan bool)
	go resource_listener.Start(rootCACerts, caCert, peerCert, peerKey, serverPort, ch)
	<-ch
	time.Sleep(time.Second*3)
	go listener.Register(rootCACerts, caCert, peerCert, peerKey, serverAddr, insecureSkipVerify, ch)
	<-ch
	<-ch
	fmt.Println("test")
	// rootCACerts []string, caCert string, peerCert string, peerKey string,
	//	serverAddr string, insecureSkipVerify bool
	return 0
}

const synopsis = "Start the consul-sm-fed-svc"
const help = `
Usage: consul-sm-fed-svc start
`

