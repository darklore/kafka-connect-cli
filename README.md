# kafka-connect-cli

Command line tool around kafka connect REST interface

```
Command line tool arounc kakfa connect REST interface

Usage:
  kccli [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  connector   Commands for connectors
  help        Help about any command
  plugins     Commands for connector plugin
  version     Show version
  worker      Get a connect worker information

Flags:
  -e, --endpoint string   Kafka connect REST endpoint (default "http://localhost:8083")
  -h, --help              help for kccli

Use "kccli [command] --help" for more information about a command.
```

## Install

```
$ go install github.com/darklore/kafka-connect-cli/cmd/kccli@latest
```
