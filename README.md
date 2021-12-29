# `httpelf`: text-powered REST API client

`httpelf` is a text-file based REST API client & manager. It uses
the nascent `.http` file format to represent http(s) requests &
is capable of making pure text-powered client requests.

## Usage

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
