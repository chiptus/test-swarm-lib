## Issues

Both compose and swarm has some difficulties with dependencies

To Install:

### Compose

run

```sh
go get -u github.com/docker/compose-cli@6bfdfa8 # to install v2-beta.4
go mod edit -replace github.com/jaguilar/vt100=github.com/tonistiigi/vt100@v0.0.0-20190402012908-ad4c4a574305
go mod tidy # to install other dependencies
```

### Swarm

run

```sh
go get -u github.com/docker/cli@v20.10.7+incompatible
go get -u github.com/docker/swarmkit@d6592ddefd8a5319aadff74c558b816b1a0b2590
go mod tidy # to install other dependencies
```
