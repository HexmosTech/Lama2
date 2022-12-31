# Getting Started

## Installation/Update

### One-line install/update in Linux/MacOS

To install/update *Lama2* and its dependencies automatically, run the following:

```
curl -s https://hexmos.com/lama2/install.sh | bash -s
```

### One-line install/update in Windows

To install/update *Lama2* and its dependencies automatically, run the following as *Administrator*:

```
choco install lama2 --version=1.0.0 --force -y
```

#### (Optional) Import your collections from Postman

Follow [guide](./postman.md)
 to import your existing Postman collections into a Plain-Text
 Lama2 repository.

#### Self update

An easier way to update the binary to latest release is through:

```
l2 -u
```

!!! Note

     Install the [VSCode extension](#from-vs-code) to launch requests from within your editor


### Manual install

#### Step 1: Install HTTPie

*Lama2* depends on [HTTPie for Terminal](https://httpie.io/cli). Use
their official instructions to get the `http` command functional in
your local system. 

#### Step 2: Download & install *Lama2* binary packages

Head over to Lama2 [releases](https://github.com/HexmosTech/Lama2/releases). Check under
the *Assets* head to find various packages. Download the relevant package for your 
operating system and CPU architecture. Once you have the package, run the following:

```
tar --overwrite -xvzf <download>.tar.gz
mv l2 /usr/local/bin
```

### Build from source

Run `make` in the project root. You'll need to have the following tools in your `PATH`:

1. `go` (v1.17+)
1. `golangcli-lint`
1. `gofumpt`

Also, you'll need to install `mkdocs` the first time; for that run these:

```
cd docs/Lama2
poetry install # get poetry from https://python-poetry.org/
```

Once `make` finishes, find the binary at `./build/l2`. Moreover,
you can launch the documentation locally through `make serve`.

Read `makefile` to find other useful helper commands.

## How to use

### From the terminal

Type `l2` into the terminal. You should get something like:

```
Usage:
  l2 [OPTIONS] [LamaAPIFile]

Application Options:
  -o, --output=      Path to output JSON file to store logs, headers and result
  -v, --verbose      Show verbose debug information
  -n, --nocolor      Disable color in httpie output
  -h, --help         Usage help for Lama2
      --version      Print Lama2 binary version

Help Options:
  -h, --help         Show this help message
```

### From VS Code

Find *Lama2 for VSCode* at the [VSCode Marketplace](https://marketplace.visualstudio.com/items?itemName=hexmos.Lama2). The extension requires the *l2* command available (usually at `/usr/local/bin/l2` for Linux/MacOS and `C:\ProgramData\chocolatey\bin` for Windows).

Once the extension is installed, open the command palette (ctrl + shift + p) and search for `Execute current file` to execute the file