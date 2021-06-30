package main

import (
	"fmt"
	"log"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/stack"
	"github.com/docker/cli/cli/command/stack/options"
	"github.com/docker/cli/cli/compose/types"
	"github.com/docker/cli/cli/flags"
	"github.com/spf13/pflag"
)

func main() {
	cli, err := command.NewDockerCli(command.WithStandardStreams())
	if err != nil {
		log.Fatal(err)
	}

	cli.Initialize(flags.NewClientOptions())

	flagset := pflag.NewFlagSet("new", pflag.ContinueOnError)
	err = stack.RunDeploy(cli, flagset, &types.Config{}, command.OrchestratorSwarm, options.Deploy{})
	if err != nil {
		log.Printf("error... %s", err)
	}

	fmt.Println("success!")
}
