# Freecli - An interactive CLI for free5GC

![free5gc-cli](https://img.shields.io/badge/Freecli-5G-blue?logo=go)
![free5gc-3.0.4](https://img.shields.io/badge/Tested-free5gc%20v3.0.4-red)


![freecli-interactive-cli](https://user-images.githubusercontent.com/41422704/99889610-220d3580-2c57-11eb-9133-f4a1daaa9258.gif)


- [Description](#description)
- [Download and build](#download-and-build)
- [Configuration](#configuration)
- [Run Freecli](#run-freecli)
- [Modules](#modules)
  - [Subscriber module](#subscriber-module)
    - [Add a fixed number of subscribers](#add-a-fixed-number-of-subscribers)
    - [Add a subscriber with a specific supi to a plmn](#add-a-subscriber-with-a-specific-supi-to-a-plmn)
    - [List all the subscribers](#list-all-the-subscribers)
    - [Delete a subscriber from a plmn](#delete-a-subscriber-from-a-plmn)
    - [Flush the database from every subscriber](#flush-the-database-from-every-subscriber)
  - [gNB module](#gnb-module)
- [Acknowledgment](#acknowledgment)

## Description

Freecli is an interactive cli utility to manage free5gc. It currently includes the following modules:

- Subscriber: manage the subscriber of free5gc in a fashion way
- 5G gNB: emulate a 5G gNB **/!\ Still under development**
- QoS: apply DSCP field to packets **/!\ Still under development**

## Download and build

To compile into a binary file

``` bash
git clone https://github.com/Srajdax/free5gc-cli.git
cd free5gc-cli
go mod download
go build -o freecli -x freecli.go
```

or to run it like an interpreter would do

``` bash
git clone https://github.com/Srajdax/free5gc-cli.git
cd free5gc-cli
go run freecli.go
```

## Configuration

Freecli uses free5gc lib under the hood with some modifications. Each freecli module is located into the `module` folder. Each module has its own configuration files located in the `config` folder. Each module is described below.

## Run Freecli

Each module is independent in freecli, at launch you can load the module by typing their name, e.g. to load the subscriber module and access all its functionalities:

```bash
./freecli

# Load the module
freecli> subscriber
2020-11-21T23:22:41Z [INFO][Freecli][Freecli] Loading subscriber module...
2020-11-21T23:22:41Z [INFO][Freecli][Subscriber Module] Successfully load module configuration config/subscriber.yaml
2020-11-21T23:22:41Z [INFO][Freecli][Subscriber Module] Successfully load ue configuration config/subscriber_ue.yaml

subscriber# The module is successfuly loaded

# Exit the module
subscriber# exit
2020-11-22T00:41:12+01:00 [INFO][Freecli][Freecli] Exiting Module...
```

## Modules

### Subscriber module

This module is used to add, delete and flush subscribers of free5gc database. It has 2 configuration files. `subscriber.yaml` holds the configuration of the module and `subscriber_ue.yaml` holds the global configuration of the subscribers to add.

You can reload the module with the command `configuration reload`

#### Add a fixed number of subscribers

```
subscriber# user random --range <number_to_add> --plmn <plmnid>
```

example
```
subscriber# user random --range 10 --plmn 20893
```

#### Add a subscriber with a specific supi to a plmn

```
subscriber# user register --supi <supi> --plmn <plmnid>
```

example
```
subscriber# user register --supi imsi-2089300000013 --plmn 20893
```

#### List all the subscribers

```
subscriber# user list
```

Note: this will also populate the auto-complete for the delete command

#### Delete a subscriber from a plmn

```
subscriber# user delete <supi>/<plmnId>
```

example
```
subscriber# user delete imsi-2089300000000/20893
```

#### Flush the database from every subscriber

```
subscriber# user flush
```

### gNB module

Currently under development

## Acknowledgment

Thanks to the free5gc team for their efforts and their open source 5G Core Network