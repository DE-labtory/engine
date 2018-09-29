/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"
	"os"
	"time"

	"github.com/it-chain/iLogger"

	"github.com/it-chain/engine/cmd/connection"
	"github.com/it-chain/engine/cmd/ivm"
	"github.com/it-chain/engine/cmd/on"
	"github.com/it-chain/engine/common"
	"github.com/it-chain/engine/conf"
	"github.com/urfave/cli"
)

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
			Value: "",
			Usage: "name for config",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "set debug mode",
		},
	}
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, ivm.IcodeCmd())
	app.Commands = append(app.Commands, connection.Cmd())
	app.Before = func(c *cli.Context) error {
		if configPath := c.String("config"); configPath != "" {
			absPath, err := common.RelativeToAbsolutePath(configPath)
			if err != nil {
				return err
			}
			conf.SetConfigPath(absPath)
		}

		if c.Bool("debug") {
			iLogger.SetToDebug()
		}

		return nil
	}
	app.Action = on.Action
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
