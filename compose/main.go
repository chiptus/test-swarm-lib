package main

import (
	"context"
	"log"

	"github.com/compose-spec/compose-go/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/compose-cli/pkg/api"
	"github.com/docker/compose-cli/pkg/compose"
)

func main() {

	dockercli, err := command.NewDockerCli(command.WithStandardStreams())
	if err != nil {
		log.Fatalf("error creating client %s", err)
	}

	err = dockercli.Initialize(flags.NewClientOptions())
	if err != nil {
		log.Fatalf("error init client %s", err)
	}

	service := compose.NewComposeService(dockercli.Client(), &configfile.ConfigFile{})
	projectName := "compose-test"
	projectOpts, err := cli.NewProjectOptions([]string{"./docker-compose.yml"}, cli.WithName(projectName))
	if err != nil {
		log.Fatalf("failed creating project options: %s", err)
	}

	project, err := cli.ProjectFromOptions(projectOpts)
	if err != nil {
		log.Fatalf("error loading files %s", err)
	}

	ctx := context.Background()

	// consumer := formatter.NewLogConsumer(ctx, os.Stdout, true, true)

	createOpt := api.CreateOptions{}

	err = service.Up(ctx, project, api.UpOptions{Start: api.StartOptions{}, Create: createOpt})
	if err != nil {
		log.Fatalf("error deploying; err=%s", err)
	}

	log.Printf("deployed successfully")

	err = service.Down(ctx, projectName, api.DownOptions{})
	if err != nil {
		log.Fatalf("error removing; err=%s", err)
	}

	log.Printf("removed successfully")

}