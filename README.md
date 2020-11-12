# A Simple CLI tool for free5GC

**/!\ Still under development**

![free5gc-cli](https://img.shields.io/badge/Golang-freecli-blue?logo=go)

- [Disclaimer](#disclaimer)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Subscriber management](#subscriber-management)

## Disclaimer

This project provides an unofficial repository for the free5gc project as the official webconsole does not cover all use-cases. free5gc-cli is a simple cli utility to interact with free5gc. I use it mainly to add or delete subscriber to the database when I can't have access to the WEBUI. This CLI is based on the webconsole and rely on free5gc lib. It is also subject to changes and further evolutions.

## Installation

Follow the installation instructions provided by free5gc repository, but instead of clonine the project: `https://github.com/free5gc/free5gc.git`. Clone this forked version of the project `https://github.com/rrajdax/free5gc`.

The compilation and installation procedure of the gNB is the same as the other free5gc core functions, you can compile the functions using 

``` bash
cd ~/free5gc
go build -o bin/freecli -x src/free5gc-cli/freecli.go
```

## Configuration

The gNB `freecli.cfg` configuration file is located in `free5gc/config` folder. A sample is also present into `free5gc-cli/config` folder.

``` yaml
info:
  version: 1.0.0
  description: free5gc-cli initial local configuration

configuration:
  mongodb:
    name: free5gc
    url: mongodb://localhost:27017
```

## Usage

``` bash
cd bin
```

### Subscriber management

Add a subscriber

```bash
./freecli --add <subscriber_configuration.yaml>
```

Remove a subscriber

```bash
./freecli --remove <imsi> <plmn>
```

Update a subscriber

```bash
./freecli --update <imsi> <plmn> <subscriber_configuration.yaml>
```

Get all the subscribers

```bash
./freecli --subscribers
```