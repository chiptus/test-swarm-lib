package main

import (
	"fmt"
	"log"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/stack/options"
	"github.com/docker/cli/cli/command/stack/swarm"
	composetypes "github.com/docker/cli/cli/compose/types"
	"github.com/docker/cli/cli/flags"
)

func main() {
	cli, err := command.NewDockerCli(command.WithStandardStreams())
	if err != nil {
		log.Fatalf("error creating client %s", err)
	}

	err = cli.Initialize(flags.NewClientOptions())
	if err != nil {
		log.Fatalf("error init client %s", err)
	}

	// flagset := pflag.NewFlagSet("new", pflag.ContinueOnError)
	err = swarm.RunDeploy(cli, options.Deploy{ResolveImage: swarm.ResolveImageAlways}, &composetypes.Config{})
	if err != nil {
		log.Fatalf("error deploying %s", err)
	}

	fmt.Println("success!")
}
