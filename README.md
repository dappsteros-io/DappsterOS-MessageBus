# DappsterOS-MessageBus

[![Go Reference](https://pkg.go.dev/badge/github.com/dappster-io/DappsterOS-MessageBus.svg)](https://pkg.go.dev/github.com/dappster-io/DappsterOS-MessageBus) [![Go Report Card](https://goreportcard.com/badge/github.com/dappster-io/DappsterOS-MessageBus)](https://goreportcard.com/report/github.com/dappster-io/DappsterOS-MessageBus) [![goreleaser](https://github.com/dappster-io/DappsterOS-MessageBus/actions/workflows/release.yml/badge.svg)](https://github.com/dappster-io/DappsterOS-MessageBus/actions/workflows/release.yml) [![codecov](https://codecov.io/gh/dappster-io/DappsterOS-MessageBus/branch/main/graph/badge.svg?token=U4S4ZSZAL9)](https://codecov.io/gh/dappster-io/DappsterOS-MessageBus)

Message bus accepts events and actions from various sources and delivers them to subscribers.

See [openapi.yaml](./api/message_bus/openapi.yaml) for API specification.




## publish api to npm

### edit version in package.json

### run
```bash
yarn

yarn start
```

### publish

Manual publish
```bash
yarn publish
```

Auto publish
```bash 
git push origin dev**
```