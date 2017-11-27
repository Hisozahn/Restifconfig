# Restifconfig

Application provides information about network interfaces on server side and sends that information to client side.
Project consist of two parts
* ifconfig-service - server listening requests (on localhost:55555 by default)
* ifconfig-cli - (command line interface) - client application for requesting above-mentioned server

### Prerequisites

This app was built and tested only on Ubuntu operating system.

If you want to run this application on your system, you need
* docker (https://www.docker.com/) 

If you want to build this project from source code, you need
* go (https://golang.org/)
* make (https://www.gnu.org/software/make/)


### Running

Follow simple steps

```
Pull latest version of server and client

docker pull hisozahn/ifconfig-service:latest
docker pull hisozahn/ifconfig-cli:latest

Run server:

docker run -d --rm -network=host hisozahn/ifconfig-service 'address to listen'

Run client (list available interfaces on server side)

docker run -it --rm -network=host hisozahn/ifconfig-cli 'your' 'arguments' 'here'
```

Note -network=host above. If you don't specify network for docker then containers will not be able to communicate with each other.
Server application accept 1 argument as address to listen to. It will listen localhost:55555 if no arguments are provided.
Client interface information available by -help flag.

### Installing

Prepare go environment variables

```
export GOROOT=/usr/lib/go-1.9
export GOPATH="$HOME"/gopath
```

Create directories for sources and imported packages

```
mkdir -p "$HOME"/gopath/src/github.com
```

Clone git repository

```
cd "$HOME"/gopath/src/github.com
git clone https://github.com/Hisozahn/Restifconfig.git Hisozahn/Restifconfig
```

Make client and server binaries

```
cd "$HOME"/gopath/src/github.com/Hisozahn/Restifconfig
make
```

Example scenario

```
cd "$HOME"/gopath/src/github.com/Hisozahn/Restifconfig/bin
./service 'localhost:11111' &
./cli --server localhost --port 11111 list

Will show list of available network interfaces on your system
```

## Running the tests

Just do

```
cd "$HOME"/gopath/src/github.com/Hisozahn/Restifconfig
go test ./...
```

## Authors

* **Igor Romanov** - *Initial work* - [Hisozahn](https://github.com/Hisozahn)

See also the list of [contributors](https://github.com/Hisozahn/Restifconfig/contributors) who participated in this project.
