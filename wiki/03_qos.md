This module is used to modify the DSCP field based on the match of some fields

- [Requirements](#requirements)
- [Configuration](#configuration)
- [Commands](#commands)
  - [Mark GTP Packets with DSCP QoS](#mark-gtp-packets-with-dscp-qos)
  - [Flush all QoS rules](#flush-all-qos-rules)

## Requirements

- Linux OS: the gNB module uses NFQUEUE and IPTABLE U32 EXTENSION under the hood

## Configuration

```yaml
info:
  version: 1.0.0
  description: "Freecli - QoS module"

configuration:
  ip:
    - 172.16.10.3
    - 172.16.10.2
  port:
    - 80
    - 8080
    - 443
    - 5060
```

You can reload the module configuration files with the command `qos# configuration reload`

## Commands

### Mark GTP Packets with DSCP QoS

```bash
qos# mark --set-phb <PHB> --destination-ip <GNB_IP> --source-ip <UPF_IP> --teid <TEID> --protocol <PROTOCOL> --destination-port <PORT> --source-port <PORT>
```

Example

```bash
qos# mark --set-phb ef --destination-ip 172.16.10.2 --source-ip 172.16.10.2 --teid 1 --protocol udp --destination-port 8080 --source-port 80
```

### Flush all QoS rules

```bash
qos# flush
```