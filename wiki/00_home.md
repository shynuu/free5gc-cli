# Welcome to Freecli Wiki

Freecli is an interactive cli utility to manage and test free5gc. It provides auto-completion and a variety of modules

- [Hardware and Software requirements](#hardware-and-software-requirements)
- [Download and build](#download-and-build)
- [Configuration of Freecli](#configuration-of-freecli)
- [Run Freecli](#run-freecli)
- [Modules documentation](#modules-documentation)

## Hardware and Software requirements

This project is tested against an Ubuntu 18.04 LTS VM with the linux kernel 5.0.23-generic. Check each module documentation for additional software requirements.

## Download and build

To compile into a binary file

``` bash
git clone https://github.com/Srajdax/free5gc-cli.git
cd free5gc-cli
go mod download
go build -o freecli -x freecli.go
```

Run it in sudo mode, QoS and gNB modules requires sudo privileges

```bash
cd freecli
sudo ./freecli
```

## Configuration of Freecli

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

## Modules documentation

Modules are under developement and continious improvements but every command listed into the documentation has been tested, is stable and is safe to use.

- [Subscriber Module - manage the subscriber of free5gc](https://github.com/Srajdax/free5gc-cli/wiki/Subscriber-Module)
- [gNB Module - emulate a 5G gNB](https://github.com/Srajdax/free5gc-cli/wiki/gNB-Module)
- [QoS Module - apply DSCP PHB to packets](https://github.com/Srajdax/free5gc-cli/wiki/QoS-Module)
- [NF Module - manage the CN Network Functions and interact with the database](https://github.com/Srajdax/free5gc-cli/wiki/NF-Module)