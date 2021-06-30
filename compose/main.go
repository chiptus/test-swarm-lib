package main

import (
	"bytes"
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
	errStream := &bytes.Buffer{}
	outputStream := &bytes.Buffer{}

	dockercli, err := command.NewDockerCli(command.WithErrorStream(errStream), command.WithOutputStream(outputStream))
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

	err = service.Up(ctx, project, api.UpOptions{})
	if err != nil {
		log.Fatalf("error deploying; err=%s", err)
	}

	log.Printf("deployed successfully")

	err = service.Down(ctx, projectName, api.DownOptions{})
	if err != nil {
		log.Fatalf("error removing; err=%s", err)
	}

	errOutput := errStream.String()
	if errOutput != "" {
		log.Fatalf("failed: %s", errOutput)
	}

	log.Println("success!")
	log.Println(outputStream.String())

}
