# `httpelf`: text-powered REST API client

`httpelf` is a text-file based REST API client & manager. It uses
the nascent `.http` file format to represent http(s) requests &
is capable of making pure text-powered client requests.

## General Usage

```
$ httpelf <request>.http
```

Find some sample request files in the 

## Installation/Updation

If you're on ubuntu, just run:

```
git pull
./install.sh
```

This will install golang/httpie if necessary,
build `httpelf` and finally move the binary
into `/usr/bin`

## FW-specific workflow

1. Pull/Clone `httpelf` repo. 
1. Run `./install.sh (will update the binary if code has updated)

1. Clone [APIHub](https://gitlab.com/flyweightgroup/apihub) repo. 
1. Create a folder for your service (ex: `impression_collector`)
1. Start defining `*.http` files for each service specific API. 
1. Use `httpelf file.http` to test the newly defined APIs. 
1. Push the API files once ready into APIHub repo