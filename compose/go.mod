module github.com/chiptus/test-swarm-lib/compose

go 1.16

require (
	github.com/compose-spec/compose-go v0.0.0-20210616101828-4e2e7e7def77
	github.com/docker/cli v20.10.7+incompatible
	github.com/docker/compose-cli v1.0.18-0.20210621100149-6bfdfa894731
)

replace github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
