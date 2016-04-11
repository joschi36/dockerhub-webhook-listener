# Docker Hub Webhook to Rancher
Receives Webhook calls from Docker Hub and Calls In-Service Upgrades in Rancher

**Based on the excellent work of Brian Goff (cpuguy83)**

## Usage
### The Docker Way
Build a Docker image from the included Dockerfile:
```bash
docker build -t myorg/myproject-dockerhub-listener .
```

You **must** supply these environment variables to start the app listening:

| Environment Variable | Purpose |
| -------------------- | ------- |
| `api_key`              | API Key to use when calling the app (supply as request param in Docker Hub config) |
| `rancher_url`          | The URL if your Rancher Server API root (e.g. http://myrancher.example.com:8080/v1/projects/1a2) |
| `rancher_user_key`     | The API Access Key (create via 'API' in the Rancher console) |
| `rancher_secret_key`   | The API Secret Key |
| `rancher_service`      | The name of the rancher service you'd like to upgrade |

Test locally, thus:
```bash
docker run -it -p 8080:8080 \
-e api_key="12345" \
-e rancher_url="http://myrancher.example.com:8080/v1/projects/1a2" \
-e rancher_user_key="1A2B3C4D5E6F1234567A" \
-e rancher_secret_key="fTEb6aGsD65RhNE32g6yTY1F8hCl5rQTz" \
-e rancher_service="myservice" \
-e listen_addr="0.0.0.0:8080" \
myorg/myproject-dockerhub-listener
```

### The Manual Way
Add a config file using `-config-file`
This file should be in INI format and is intended for use with handlers

**Example**
```
[apiKeys]
key = 1a2b3c4d5e6f

[rancher]
url = http://myrancher.example.com:8080/v1/projects/1a2
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
