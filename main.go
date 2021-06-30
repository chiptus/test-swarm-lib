package main

import (
	"fmt"
	"log"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/stack/loader"
	"github.com/docker/cli/cli/command/stack/options"
	"github.com/docker/cli/cli/command/stack/swarm"
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

	ns := "test"

	opts := options.Deploy{
		Namespace:    ns,
		ResolveImage: swarm.ResolveImageAlways,
		Composefiles: []string{"./docker-compose.yml"},
	}
	config, err := loader.LoadComposefile(cli, opts)
	if err != nil {
		log.Fatalf("error loading compose file; err=%s", err)
	}

	err = swarm.RunDeploy(cli, opts, config)
	if err != nil {
		log.Fatalf("error deploying; err=%s", err)
	}

	log.Printf("deployed successfully")

	err = swarm.RunRemove(cli, options.Remove{Namespaces: []string{ns}})
	if err != nil {
		log.Fatalf("error removing; err=%s", err)
	}

	log.Printf("removed successfully")

	fmt.Println("success!")
}
