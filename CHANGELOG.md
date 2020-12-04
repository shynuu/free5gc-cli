# CHANGELOG

## 2020-12-04 - Version 0.4

### New

- Initial Release of GNB Module: Register UE, PDU Session Request, QoS DSCP Marking from gNB to UPF
- Initial Release of QoS Module: QoS DSCP Marking of GTP packet (useful for marking egress traffic from UPF to gNB as UPF does not implement GTP PDU Session Extension Header)

### Fix

- Fix nil pointer when exiting socket
- Fix logs folder creation

### Improvements

- Wiki documentation
- Readme
- Cleaning code for gNB module

## 2020-11-27

### New

#### Network Function Module

- Database commands

#### QoS Module

- Initialize QoS Module

#### gNB Module

- Register and Deregister user tested against bare-metal and compose version of free5gc v3.0.4

### Improvements

- Clean code of subscriber module
- Clean code of gNB module

## 2020-11-26

### New

- Initialize gNB module

### Update

- Update README
- Clean code of subscriber module

## 2020-11-21

Initial release of Freecli