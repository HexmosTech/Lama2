# `httpelf`: text-powered REST API client

`httpelf` is a text-file based REST API client & manager. It uses
the nascent `.http` file format to represent http(s) requests &
is capable of making pure text-powered client requests.

## Usage

```
$ httpelf <request>.http
```

Find some sample request files in the 

## Installation

### Step 1: Install `httpie`

In Ubuntu, this can be done via:

```
sudo snap install httpie
```

### Step 2: Get the `httpelf` binary

Simply download the binary in the `./build/` directory and place it 
in your preferred `bin` folder, such as `/usr/bin` or `~/.local/bin`.

## Build

Take a look at `~/.build.sh`; it installs the necessary dependencies
and invokes go build to generate the binary