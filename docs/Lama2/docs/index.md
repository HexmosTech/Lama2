# **`Lama2`**: Plain-text Powered REST API Client for Teams

![](./banner.png)


![Build](https://github.com/HexmosTech/Lama2/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/HexmosTech/Lama2/actions/workflows/testapp.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/HexmosTech/lama2)](https://goreportcard.com/report/github.com/HexmosTech/lama2)

Quick links: [How-tos](./howto.md) | [Explanations](./explanation.md)

## About

*Lama2* is a plain-text powered REST API client & manager built for serious engineering teams.
*Lama2* specifies the [`.l2` syntax](tutorials/examples.html) for describing APIs, and implements a [CLI](tutorials/installation.html) to execute `.l2` files. Engineers can [collaborate](tutorials/collaboration.html) on `.l2` files using version control. *Lama2* [integrates](tutorials/editor.html) nicely with
IDEs and text-editors as well. Think of *Lama2* as [**markdown for APIs.**](reference/philosophy.html)
## Benefits

1. **Plain-text files:** Store APIs in the plain-text `.l2` API files. [Simple](explanation/l2format.html) and [human-friendly](explanation/syntax.html) syntax. Learn [basics](tutorials/examples.html) within  minutes!
1. **Simple CLI:** Launch the [CLI](tutorials/installation.html) tool `l2` on API files to make REST API requests.
1. **Editor support:** Invoke *Lama2* from your [favorite text editor](tutorials/installation.html#from-vs-code) or IDE. Helpful documentation and tool support to build editor extensions included.
1. **Longevity & track-ability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share [API repo](tutorials/collaboration.html) with teammates and colleagues to collaborate
1. **Documentation:** [Explore](index.html) examples, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a *recursive descent parser*, based on a [formal grammar](reference/grammar.html). [Dig into details](reference/parser.html) and [implement](https://github.com/HexmosTech/Lama2/tree/main/parser) new syntax (ex: to support `websockets`)


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

## Community and Support

1. [Full documentation @hexmos.com/lama2](https://hexmos.com/lama2/)
1. Report issues and propose improvements at [Github issues](https://github.com/HexmosTech/Lama2/issues)
1. Discuss with the community at [Discord](https://discord.gg/zTmxXA6F)
