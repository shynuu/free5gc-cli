<p align="center">
<img width="500" alt="image" src="https://user-images.githubusercontent.com/41422704/100396438-44cb8f80-3045-11eb-95d6-d2d5c84e3740.png">
</p>

<p align="center">
<a href="https://github.com/srajdax/free5gc-cli/releases"><img src="https://img.shields.io/badge/Freecli%205G-v0.2-blue?logo=go" alt="Freecli 5G"/></a>
<a href="https://github.com/free5gc/free5gc"><img src="https://img.shields.io/badge/Tested-free5gc%20v3.0.4-red" alt="Free5GC v3.0.4"/></a>
<img src="https://img.shields.io/badge/OS-Linux-g" alt="OS Linux"/>
<a href="https://github.com/Srajdax/free5gc-cli/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-Apache%202-lightgrey" alt="Apache 2 License"/></a>
</p>

![freecli-interactive-cli](https://user-images.githubusercontent.com/41422704/99889610-220d3580-2c57-11eb-9133-f4a1daaa9258.gif)

Freecli 5G is an interactive cli utility to test and experiment free5gc.

## Features

- Subscriber: manage the subscriber of free5gc in a fashion way
- 5G gNB: emulate a 5G gNB: Registration of UE, PDU Session Request, Provide linux level GTP encapsulation allowing to send and receive real traffic to/from the UPF, QoS DSCP Marking
- QoS: apply DSCP PHB to packets
- NF: manage the CN Network Functions and interact with the database

Read the [WIKI](https://github.com/Srajdax/free5gc-cli/wiki) for more documentation on the CLI and each module

## To avoid any confusion

Although this project uses free5gc libraries, it is not part of the official free5gc project. 
I've developed it to simplify my 5G deployments/experiments and thought it was nice to share it. 

All contributions are welcome !

## Acknowledgments

Thanks to the free5gc team for their efforts, their lib and their open source 5G Core Network