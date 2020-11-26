
This module is used for subscriber management. 

- [Requirements](#requirements)
- [Configuration files](#configuration-files)
- [Commands](#commands)
  - [Add a fixed number of subscribers](#add-a-fixed-number-of-subscribers)
  - [Add a subscriber with a specific supi to a plmn](#add-a-subscriber-with-a-specific-supi-to-a-plmn)
  - [List all the subscribers](#list-all-the-subscribers)
  - [Delete a subscriber from a plmn](#delete-a-subscriber-from-a-plmn)
  - [Flush the database from every subscriber](#flush-the-database-from-every-subscriber)

## Requirements

- Free5gc running and more specifically mongo service

## Configuration files


It has 2 configuration files.

`subscriber.yaml` holds the configuration of the module
```yaml
info:
  version: 1.0.0
  description: "Freecli - subscriber module"

configuration:
  mongodb:
    name: free5gc
    url: mongodb://172.16.100.30:27018

plmn:
  value:
    - 20893
    - 20810
```

`subscriber_ue.yaml` holds the global configuration of the subscribers to add
```yaml
ueId: imsi-2089300007487
servingPlmnId: 20893

AuthenticationSubscription:
  authenticationMethod: 5G_AKA
  authenticationManagementField: 8000
  milenage:

[...]

      snssai:
        sst: 1
        sd: 112233
      smPolicyDnnData:
        internet:
          dnn: internet
```


You can reload the module configuration files with the command `subscriber# configuration reload`


## Commands

### Add a fixed number of subscribers

```
subscriber# user random --range <number_to_add> --plmn <plmnid>
```

example
```
subscriber# user random --range 10 --plmn 20893
```

### Add a subscriber with a specific supi to a plmn

```
subscriber# user register --supi <supi> --plmn <plmnid>
```

example
```
subscriber# user register --supi imsi-2089300000013 --plmn 20893
```

### List all the subscribers

```
subscriber# user list
```

Note: this will also populate the auto-complete for the delete command

### Delete a subscriber from a plmn

```
subscriber# user delete <supi>/<plmnId>
```

example
```
subscriber# user delete imsi-2089300000000/20893
```

### Flush the database from every subscriber

```
subscriber# user flush
```
