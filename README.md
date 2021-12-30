# 🧝 `httpelf`: text-powered REST API client

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
git clone git@gitlab.com:flyweightgroup/httpelf.git
./install.sh
```

This will install golang/httpie if necessary,
build `httpelf` and finally move the binary
into `/usr/bin` (`/usr/local/bin` in macOS)

### Updates

If you're on ubuntu, just run:

```
git pull
./install.sh
```

**Note:** overwrites the existing binary in `/usr/bin` (`/usr/local/bin` in macOS)

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
1. Two request modes are supported presently: `JSON` & `MULTIPART`
    1. `JSON` is the default mode for elf. 
    1. The word `MULTIPART` anywhere in the request file will trigger the `MULTIPART` mode
1. `MULTIPART` mode supports `file@./relative/path` syntax for file uploads
    1. The paths are always *relative* to the `.http` file
    1. An example can be found in `sample_http/fabi_register.http`
