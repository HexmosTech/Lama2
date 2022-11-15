# `Lama2`: Plain-text Powered REST API Client


<div align="center">
<img src="banner.png" width="80%" />

![Build](https://github.com/HexmosTech/Lama2/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/HexmosTech/Lama2/actions/workflows/testapp.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/HexmosTech/lama2)](https://goreportcard.com/report/github.com/HexmosTech/lama2)

*Quick Links: [Installation](#installation) | [Demo](#terminal-demo-a-post-request) | [Full Documentation](https://hexmos.com/lama2)*
</div>


## About

*Lama2* is a plain-text powered REST API client & manager built for serious engineering teams.
*Lama2* specifies the .l2 syntax for describing APIs, and implements a CLI to execute .l2 files. Engineers can collaborate on `.l2` files using version control. *Lama2* integrates nicely with
IDEs and text-editors as well. Think of *Lama2* as **markdown for APIs.**

## Benefits 

1. **Plain-text files:** Store APIs in the plain-text `.l2` API files. Simple and human-friendly syntax. Learn basics within  minutes!
1. **Simple CLI:** Launch the CLI tool `l2` on API files to make REST API requests.
1. **Editor support:** Invoke *Lama2* from your favorite text editor or IDE. Helpful documentation and tool support to build editor extensions included.
1. **Longevity & track-ability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share API repo with teammates and colleagues to collaborate
1. **Documentation:** Explore tutorials, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a *recursive descent parser*, based on a clearly specified grammar. Dig into details and implement new syntax (ex: to support `websockets`)

## Terminal Demo: A POST request

![](docs/Lama2/docs/demo.gif)

## VSCode Demo: The same POST request

![](docs/Lama2/docs/demo2.gif)

## Installation

```
curl -s https://raw.githubusercontent.com/HexmosTech/Lama2/main/install.sh | bash -s
```

Learn more about other installation methods from [our docs](https://hexmos.com/lama2/tutorials/installation.html).

### VSCode Extension

Find *Lama2 for VSCode* at the [VSCode Marketplace](https://marketplace.visualstudio.com/items?itemName=hexmos.Lama2). The extension requires the *l2* command available (usually at `/usr/local/bin/l2`).

## Design Philosophy

1. Delegate subtasks to open and mature tools:
    1. API Collaboration - git
    1. API Organization - OS file system
    1. API Editing - VSCode/IDEs
    1. API Requests - HTTPie
1. Start a language formalization effort from the beginning (learning from Markdown history)
1. Keep language syntax simple
1. Invest into good documentation
1. Strive to keep components decoupled




## General Usage

Simply type any of the following to get options (option support is in progress):

```
elf
elf -h 
elf --help
```

Find some sample request files in the `sample_http` folder.

### Prettification

To fix messed up JSONs in place:

```
elf -p file.l2
```

Prettify all files in a directory (useful for large API collections such as FW-specific
APIHub):
```
find . -type f -name "*.l2" -exec elf -p {} \
```

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
1. Start defining `*.l2` files for each service specific API. 
1. Use `elf file.l2` to test the newly defined APIs. 
1. Push the API files once ready into APIHub repo


## Development

### Testing

```
go test -v
```