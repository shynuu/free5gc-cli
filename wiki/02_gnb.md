This module is used to emulate a gNB and UE functionalities

- [Requirements](#requirements)
- [Implementend Features](#implementend-features)
- [3GPP Procedures implemented](#3gpp-procedures-implemented)
- [Configuration](#configuration)
- [Commands](#commands)
  - [Registering a UE on the PLMN](#registering-a-ue-on-the-plmn)
  - [PDU Session request](#pdu-session-request)
  - [Mark PDU Session with DSCP](#mark-pdu-session-with-dscp)
  - [Flush all the QoS rules associated to PDU Sessions](#flush-all-the-qos-rules-associated-to-pdu-sessions)

## Requirements

- Free5gc running with the AMF and UPF reacheable
- Linux OS: The gNB uses TUN interface and iptables
- Launch Freecli using sudo privileges

## Implementend Features

- Registration of UE, PDU Session Request
- Provide IP level GTP encapsulation allows to send and receive real traffic to the UPF
- QoS DSCP Marking

## 3GPP Procedures implemented

3GPP procedures implemented and tested

- [x] UE Registration
- [ ] UE De-Registration
- [x] PDU Session Request
- [ ] PDU Session Release

## Configuration

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
    ipv4Addr: "172.16.20.2"
    port: 9487
  gtpInterface:
    ipv4Addr: "172.16.20.3"
    port: 2152
  ueSubnet: "60.60.0.0"
  security:
    networkName: 5G:mnc093.mcc208.3gppnetwork.org
    k: 5122250214c33e723a5dd523fc145fc0
    opc: 981d464c7c52eb6e5036234984ad0bcf
    op: c9e8763286b5b9ffbdf56e1297d0887b
    sqn: 16f3b3f70fc2
  # Used for autocompletion
  snssai:
    - sst: 1
      sd: "112233"
    - sst: 1
      sd: "010203"
  # Used for autocompletion
  dnn:
    - internet
    - voip
  # Used for autocompletion
  ue:
    - supi: imsi-2089300000000
      plmn: 20893
    - supi: imsi-2089300000001
      plmn: 20893
    - supi: imsi-2089300000002
      plmn: 20893
    - supi: imsi-2089300000003
      plmn: 20893
  # Specify the linux TUN interface to be created
  tun: tun_gnb
```

When loading the gNB module, a TUN interface will be created and under the name and the IP address specified in the configuration file (here `tun_gnb` and `172.16.20.3`).

After launching the module, you can check that your interface has been created by typing the command `ip addr`.

This interface acts as a router interface for the UE and each packet incoming onto this interface would be encapsulated into a GTP tunnel to the UPF.

Schema below represents the instanciation of the configuration file above.

```      
                                                                              AMF
                                                                     +------------------+
     UE                                                              |                  |
+-----------+                                           amf_interface|                  |
|           |                                           +------------+  172.16.10.2     |
|           |                                           |            |  38412           |
|           |                                           |            |                  |
| 60.60.0.1 +-----+              gNB                    |            |                  |
|           |     |      +------------------+           |            |                  |
+-----------+     |      |                  |           |            |                  |
  Docker/VM       |      |     172.16.20.2  | ng-ran    |            |                  |
                  |      |            9487  +-----------+            |                  |
                  +------+                  |                        +------------------+
                         |                  |                             Docker/VM
                         |                  |
                         |                  |                                 UPF
                         |     172.16.20.3  | tun_gnb                +------------------+
                         |            2152  +------------+           |                  |
                         |                  |            |           |                  |
                         +------------------+            +-----------+  172.16.10.2     |              Data
                                 VM                     upf_interface|  2152            |             Network
                                                                     |                  |          +-----------+
                                                                     |                  |          |           |
                                                                     |                  |          |           |
                                                                     |     60.60.0.101  +----------+           |
                                                                     |                  |          |           |
                                                                     |                  |          |           |
                                                                     +------------------+          +-----------+
                                                                          Docker/VM
```

## Commands

### Registering a UE on the PLMN

```bash
gnb# user register --user <supi> 
```

Example

```bash
gnb# user register --user imsi-2089300000001
gnb# 2020-12-03T23:22:41Z [INFO][Freecli][gNB Module] Successfully register user imsi-2089300000001 on the network
```

### PDU Session request

A user must be registered before running this command

```bash
gnb# user pdu-session request --user <supi> --snssai <snssai> --dnn <dnn>
```

Example

```bash
gnb# user pdu-session request --user imsi-2089300000001 --snssai 01010203 --dnn internet
gnb# 2020-12-03T23:22:41Z [INFO][Freecli][gNB Module] UE Information: IP address of UE 60.60.0.1
gnb# 2020-12-03T23:22:41Z [INFO][Freecli][gNB Module] UE Information: TEID 1
gnb# 2020-12-03T23:22:41Z [INFO][Freecli][gNB Module] UE Information: PDU Session ID 10
gnb# 2020-12-03T23:22:41Z [INFO][Freecli][gNB Module] Successfully Established PDU Session for user imsi-2089300000001 with snssai 01010203 and dnn internet
```

### Mark PDU Session with DSCP

Note: Source port is optionnal

```bash
gnb# user pdu-session qos add --set-phb <PHB> --session <sessionID> --protocol <PROTOCOL> --destination-port <PORT> --source-port <PORT>
```

Example without source port

```bash
gnb# user pdu-session qos add --set-phb ef --session imsi-2089300000001/1 --protocol tcp --destination-port 80
```

### Flush all the QoS rules associated to PDU Sessions

```bash
gnb# user pdu-session qos flush
```

