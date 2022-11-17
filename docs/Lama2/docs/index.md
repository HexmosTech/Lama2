# **`Lama2`**: Plain-text Powered REST API Client for Teams

![](./banner.png)


![Build](https://github.com/HexmosTech/Lama2/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/HexmosTech/Lama2/actions/workflows/testapp.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/HexmosTech/lama2)](https://goreportcard.com/report/github.com/HexmosTech/lama2)

Quick links: [How-tos](./howto.md) | [Explanations](./explanation.md)

## About

*Lama2* is a plain-text powered REST API client & manager built for serious engineering teams.
*Lama2* specifies the .l2 syntax for describing APIs, and implements a CLI to execute .l2 files. Engineers can collaborate on `.l2` files using version control. *Lama2* integrates nicely with
IDEs and text-editors as well. Think of *Lama2* as **Markdown for APIs.**

## Benefits

1. **Plain-text files:** Store APIs in the plain-text `.l2` API files. Simple and human-friendly syntax. Learn basics within  minutes!
1. **Simple CLI:** Launch the CLI tool `l2` on API files to make REST API requests.
1. **Editor support:** Invoke *Lama2* from your favorite text editor or IDE. Helpful documentation and tool support to build editor extensions included.
1. **Longevity & track-ability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share API repo with teammates and colleagues to collaborate
1. **Documentation:** Explore tutorials, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a *recursive descent parser*, based on a clearly specified grammar. Dig into details and implement new syntax (ex: to support `websockets`)


## Terminal Demo: A POST request

![](demo.gif)

**`basic_post.l2`:**

```bash
POST
https://httpbin.org/post
{
    "hello": "world"
}
```

**Execute:**

```
l2 basic_post.l2
```

## VSCode Demo: The same POST request

![](demo2.gif)
