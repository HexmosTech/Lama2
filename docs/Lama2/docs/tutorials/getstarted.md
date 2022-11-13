# Getting Started

## Installation

### One-line install (Recommended in Linux/MacOS)

To install *Lama2* and its dependencies automatically, run the following:

```
curl -s https://raw.githubusercontent.com/HexmosTech/Lama2/main/install.sh | bash -s
```

### Manual install

#### Step 1: Install Httpie

*Lama2* depends on [Httpie for Terminal](https://httpie.io/cli). Use
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


## How to use

### From the terminal

Type `l2` into the terminal. You should get something like:

```
Usage:
  l2 [OPTIONS] [LamaAPIFile]

Application Options:
  -v, --verbose      Show verbose debug information
  -p, --prettify     Prettify specified .lama file
  -s, --sort         Sort specification into recommended order
  -n, --nocolor      Disable color in httpie output
  -h, --help         Usage help for Lama2

Help Options:
  -h, --help         Show this help message
```

### From VS Code

TODO