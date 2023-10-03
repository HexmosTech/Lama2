# `Lama2`: Free, Lightweight & Collaborative API Client


<div align="center">
<img src="https://hexmos.com/lama2/banner.png" width="80%" />

![Build](https://github.com/HexmosTech/Lama2/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/HexmosTech/Lama2/actions/workflows/testapp.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/HexmosTech/lama2)](https://goreportcard.com/report/github.com/HexmosTech/lama2)

<a href="https://www.producthunt.com/posts/hexmos-lama2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-hexmos&#0045;lama2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=370586&theme=light" alt="Hexmos&#0032;Lama2 - Plain&#0045;Text&#0032;Powered&#0032;REST&#0032;API&#0032;Client&#0032;for&#0032;Teams | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

*Quick Links: [Installation](#installationupdate) | [Demo](#terminal-demo-a-post-request) | [Full Documentation](https://hexmos.com/lama2)*
</div>

> **Note**
> The frontend part of the Lama2 project lives in [Lama2Code](https://github.com/HexmosTech/Lama2Code). The
> extension tends to be updated along with the Lama2 core in this repo.

## The Problem

Most engineering teams have trouble collaborating on APIs or end up paying too much for Postman/Insomnia. Our 100% FOSS product *Lama2* helps your team collaborate on APIs for free based on git, so that you can save money and avoid friction.

## About

*Lama2* is a Plain-Text powered REST API client & manager built for serious engineering teams.
*Lama2* specifies the [`.l2` syntax](https://hexmos.com/lama2/tutorials/examples.html) for describing APIs, and implements a [CLI](https://hexmos.com/lama2/tutorials/installation.html) to execute `.l2` files. Engineers [collaborate](https://hexmos.com/lama2/tutorials/collaboration.html) on `.l2` files using version control. *Lama2* [integrates](https://hexmos.com/lama2/tutorials/editor.html) nicely with
IDEs and text-editors as well. Think of *Lama2* as [**Markdown for APIs.**](https://hexmos.com/lama2/reference/philosophy.html)

## Benefits 

1. **Plain-Text files:** Store APIs in the Plain-Text `.l2` API files. [Simple](https://hexmos.com/lama2/explanation/l2format.html) and [human-friendly](https://hexmos.com/lama2/explanation/syntax.html) syntax. Learn [basics](https://hexmos.com/lama2/tutorials/examples.html) within  minutes!
1. **Simple CLI:** Launch the [CLI](https://hexmos.com/lama2/tutorials/installation.html) tool `l2` on API files to make REST API requests.
1. **Editor support:** Invoke *Lama2* from your [favorite text editor](https://hexmos.com/lama2/tutorials/installation.html#from-vs-code) or IDE. Helpful documentation and tool support to [build editor extensions](https://hexmos.com/lama2/tutorials/editor.html) included.
1. **Longevity & track-ability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share [API repo](https://hexmos.com/lama2/tutorials/collaboration.html) with teammates and colleagues to collaborate
1. **Code generation:** [Convert](https://hexmos.com/lama2/tutorials/codegen.html) `l2` API definitions into your preferred language/library, be it Python, Javascript, Java or many other options.
1. **Powerful chaining:** Chain multiple `l2` requests through embedded [Javascript blocks](https://hexmos.com/lama2/tutorials/examples.html#chain-requests-using-javascript)
1. **Documentation:** [Explore](https://hexmos.com/lama2/index.html) examples, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a *recursive descent parser*, based on a [formal grammar](https://hexmos.com/lama2/reference/grammar.html). [Dig into details](https://hexmos.com/lama2/reference/parser.html) and [implement](https://github.com/HexmosTech/Lama2/tree/main/parser) new syntax (ex: to support `websockets`)

> :bulb: **Coming from Postman?** Let *Lama2* help you [convert](https://hexmos.com/lama2/tutorials/postman.html) Postman
data dumps into a  nice  Plain-Text Lama2 repo.

## Terminal Demo: A POST request

![](https://hexmos.com/lama2/demo.gif)

## VSCode Demo: The same POST request

![](https://hexmos.com/lama2/demo2.gif)

## Installation/Update

For **Linux/MacOS** systems, run the following command to either install or update *Lama2*:

```
curl -s https://hexmos.com/lama2/install.sh | bash -s
```

For **Windows** systems, run the following command on powershell as an **Administrator** to either install or update *Lama2*: 

```
choco install lama2 --version=1.0.0 --force -y
```

Learn more about other installation methods from [our docs](https://hexmos.com/lama2/tutorials/installation.html).

### (Optional) Import your collections from Postman

Follow [guide](https://hexmos.com/lama2/tutorials/postman.html)
 to import your existing Postman collections into a Plain-Text
 Lama2 repository.
 
### Self update

An easier way to update the binary for **Linux/MacOS** to latest release is through:

```
l2 -u
```


### VSCode Extension

Find *Lama2 for VSCode* at the [VSCode Marketplace](https://marketplace.visualstudio.com/items?itemName=hexmos.Lama2). The extension requires the *l2* command available (usually at `/usr/local/bin/l2`).

## Design Philosophy

Lama2 takes inspiration from Markdown. Think of our approach as Markdown for APIs.

In particular, we strive to:

1. Delegate subtasks to mature and preferably open tools:
    1. API Collaboration - git
    1. API Organization - OS file system + IDEs
    1. API Editing - VSCode/IDEs
    1. API Requests - HTTPie ([httpie-go](https://github.com/nojima/httpie-go), to be specific)
    1. Request Chaining - Embedded Javascript
1. Start a language formalization effort from the beginning (learning from Markdown history)
1. Keep language syntax simple
1. Invest into good documentation; however, ensure users can perform competently without referencing documentation as much as possible.
1. Strive to keep components decoupled


## Community and Support

1. [Full documentation @hexmos.com/lama2](https://hexmos.com/lama2/)
1. Report issues and propose improvements at [Github issues](https://github.com/HexmosTech/Lama2/issues)
1. Discuss with the community at [Discord](https://discord.gg/zTmxXA6F) and [Github Discussions](https://github.com/HexmosTech/Lama2/discussions)


## Acknowledgements

1. The excellent VSCode ecosystem helping us develop our [VSCode extension](https://github.com/HexmosTech/Lama2Code)
1. httpie-go: A golang version of httpie which we have forked and integrated with Lama2. [Source](https://github.com/HexmosTech/httpie-go)
1. httpsnippet: A HAR Request to multi-language converter; we use a custom fork of [httpsnippet](https://github.com/HexmosTech/httpsnippet) for code generation from l2 files.
1. goja: A pure golang based JS interpreter with great, correct implementation of ECMA5. We use [goja](https://github.com/dop251/goja) for JS processor blocks in multi-stage API files.