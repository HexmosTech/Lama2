# `httpelf`: text-powered REST API client

`httpelf` is a text-file based REST API client & manager. It uses
the nascent `.http` file format to represent http(s) requests &
is capable of making pure text-powered client requests.

[[_TOC_]]

## General Usage

```
$ elf <request>.http
```

Find some sample request files in the 

### Installation

If you're on ubuntu, just run:

```
git clone git@gitlab.com:flyweightgroup/apihub.git 
./install.sh
```

This will install golang/httpie if necessary,
build `httpelf` and finally move the binary
into `/usr/bin`

### Updates

If you're on ubuntu, just run:

```
git pull
./install.sh
```

**Note:** overwrites the existing binary in `/usr/bin`

## FW-specific workflow

1. Pull/Clone `httpelf` repo. 
1. Run `./install.sh (will update the binary if code has updated)

1. Clone [APIHub](https://gitlab.com/flyweightgroup/apihub) repo. 
1. Create a folder for your service (ex: `impression_collector`)
1. Start defining `*.http` files for each service specific API. 
1. Use `elf file.http` to test the newly defined APIs. 
1. Push the API files once ready into APIHub repo

## `.http` format

1. Lines starting with `#` are comments and are ignored altogether
1. Order of directives doesn't matter (experiment putting headers before json or json before url). One can simply "dump" pieces of the query into the file as they emerge in the mind & expect `elf` to just handle it
1. Cookies are specified as a header: `Cookie:'sessionid=foo;another-cookie=bar'`
