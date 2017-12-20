# NAP

[![Build Status](https://api.travis-ci.org/JetMuffin/nap.svg?branch=master)](https://travis-ci.org/JetMuffin/nap)
[![Coverage Status](https://coveralls.io/repos/github/JetMuffin/nap/badge.svg?branch=master)](https://coveralls.io/github/JetMuffin/nap?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/JetMuffin/nap)](https://goreportcard.com/report/github.com/JetMuffin/nap)

**UNDER CONSTRUCTION**

## Features

- [x] Application management
- [ ] Workloads management
- [ ] Auto-scaling
- [ ] Rolling update
- [ ] Health checks
- [ ] User Authorization
- [ ] Docker compose support
- [x] Task web console
- [ ] IP-Per-Task
- [ ] Multiple schedule stategy

## Getting Start

### Building NAP

Note that the master branch is the development branch and checkout your git head to a release branch.

Prerequisites:

* Go >= 1.7.1

Fetch source codes from github:

```
$ go get -u -v github.com/JetMuffin/nap
$ cd $GOPATH/src/github.com/JetMuffin/nap
```

Run `make`, which creates the NAP binary at `bin/nap`.


### Start master

```
$ ./bin/nap master -c conf/nap.toml
```

Use command `./bin/nap --help` to see full usage.

## Trouble Shooting

TBD

## Contact

TBD

## License

MIT License