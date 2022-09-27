# transcode-sros

Golang implementation of the original transcode-sros python version (https://github.com/door7302/transcode-sros).

Usage:
```
$ ./transcode-sros -h
NAME:
   transcode-sros - transcode a router CLI configuration file

USAGE:
   transcode-sros [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config-file value, -f value  Unflatten SROS configuration file (default: "inventory.yml")
   --shorten, -s                  service vprn cli without service and customer name (default: false)
   --debug, -d                    debug mode for troubleshooting (default: false)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)


$ ./transcode-sros -f router-sros.cfg
```

The ***shorten*** option removes "name" and "customer" for every configuration line in a VPRN.

Example result without shorten option:
```
/configure service vprn 100 name "Test VPRN" customer 1 create
/configure service vprn 100 name "Test VPRN" customer 1 description "Test VPRN"
```
Example result with shorten option:
```
/configure service vprn 100 name "Test VPRN" customer 1 create
/configure service vprn 100 description "Test VPRN"
```