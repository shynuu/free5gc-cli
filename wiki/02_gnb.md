**UNDER DEVELOPEMENT**

This module is used to emulate a gNB and UE functionalities

- [Requirements](#requirements)
- [Configuration file](#configuration-file)
- [3GPP Procedures implemented](#3gpp-procedures-implemented)

## Requirements

- Free5gc running and more specifically mongo service
- Linux OS: the gNB module uses NFQUEUE and IPTABLE U32 EXTENSION under the hood

## Configuration file

It has 1 configuration file

gnb.yaml holds all the necessary information to emulate the gNB.

```yaml
info:
  version: 1.0.0
  description: "Freecli - gNB module"

configuration:
  amfInterface:
    ipv4Addr: "172.16.10.2"
    port: 38412
  upfInterface:
    ipv4Addr: "172.16.10.2"
    port: 2152
  ngranInterface:
    ipv4Addr: "172.16.10.3"
    port: 9487
  gtpInterface:
    ipv4Addr: "172.16.10.3"
    port: 2152
  ueSubnet: "60.60.0.0/16"
  security:
    networkName: 5G:mnc093.mcc208.3gppnetwork.org
    k: 5122250214c33e723a5dd523fc145fc0
    opc: 981d464c7c52eb6e5036234984ad0bcf
    op: c9e8763286b5b9ffbdf56e1297d0887b
    sqn: 16f3b3f70fc2
  snssai:
    - sst: 1
      sd: "112233"
    - sst: 1
      sd: "010203"
  ue:
    - supi: imsi-2089300007487
      plmn: 20893
    - supi: imsi-2089300007486
      plmn: 20893
```

You can reload the module configuration files with the command `gnb# configuration reload`

## 3GPP Procedures implemented

The following 3GPP 5G procedures are currently implemented:

- UE Registration
- UE De-Registration
- PDU Session Establishment
- PDU Session Release