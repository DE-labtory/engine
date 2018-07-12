package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"net"

	"github.com/it-chain/it-chain-Engine/cmd/icode"
	"github.com/it-chain/it-chain-Engine/cmd/peer"
	"github.com/it-chain/it-chain-Engine/conf"
	"github.com/urfave/cli"
)

func PrintLogo() {
	fmt.Println(`
	___  _________               ________  ___  ___  ________  ___  ________
	|\  \|\___   ___\            |\   ____\|\  \|\  \|\   __  \|\  \|\   ___  \
	\ \  \|___ \  \_|____________\ \  \___|\ \  \\\  \ \  \|\  \ \  \ \  \\ \  \
	 \ \  \   \ \  \|\____________\ \  \    \ \   __  \ \   __  \ \  \ \  \\ \  \
	  \ \  \   \ \  \|____________|\ \  \____\ \  \ \  \ \  \ \  \ \  \ \  \\ \  \
           \ \__\   \ \__\              \ \_______\ \__\ \__\ \__\ \__\ \__\ \__\\ \__\
	    \|__|    \|__|               \|_______|\|__|\|__|\|__|\|__|\|__|\|__| \|__|
	`)
}

func main() {

	app := cli.NewApp()
	app.Name = "it-chain"
	app.Version = "0.1.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "it-chain",
			Email: "it-chain@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config",
			Usage: "name for config",
		},
	}
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, peer.PeerCmd())
	app.Commands = append(app.Commands, icode.IcodeCmd())
	app.Action = func(c *cli.Context) error {
		configName := c.String("config")
		conf.SetConfigName(configName)
		return start()
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start() error {
	configuration := conf.GetConfiguration()
	ln, err := net.Listen("tcp", configuration.Common.NodeIp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on %q: %s\n", conf.GetConfiguration().GrpcGateway.Ip, err)
		return err
	}
	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on %q: %s\n", conf.GetConfiguration().GrpcGateway.Ip, err)
		return err
	}

	initGateway()
	initTxPool()
	initIcode()
	initPeer()
	return nil
}

func initGateway() error {
	return nil
}
func initIcode() error {
	return nil
}
func initPeer() error {
	return nil
}
func initTxPool() error {
	return nil
}
func initConsensus() error {
	return nil
}
