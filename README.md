# Docker Hub Webhook to Rancher
Receives Webhook calls from Docker Hub and Calls In-Service Upgrades in Rancher
**Based on the excellent work of Brian Goff (cpuguy83)**

## Usage
Add a config file using `-config-file`
This file should be in INI format and is intended for use with handlers
**Example**
```
[apiKeys]
key = 1a2b3c4d5e6f

[rancher]
url = http://myrancherserver:8080/v1/projects/1e2
userkey = <Rancher API User Key>
secretkey = <Rancher API User Secret>
service = <Rancher Service Name>
```

Register a handler in `handler.go`
You an use "Logger" as a reference to how to set this up.

```bash
cd hub-listener
go build
./hub-listener --listen=0.0.0.0:80 --config-file=config.ini
```

You should use SSL.
To do so ad a `tls` section to your config file, with a `cert` and a `key` file

You should also use authentication.
Right now DockerHub doesn't really support this, but you can use an api key as a
query param.
To handle this, you need to add an `apikeys` section to the config file along
with a list of `key`s
