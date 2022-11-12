# **`Lama2`**: Plain-text Powered REST API Client for Teams

![](./banner.png)

Quick links: [How-tos](./howto.md) | [Explanations](./explanation.md)

## About

*Lama2* is a plain-text powered REST API client & manager built for serious engineering teams.
*Lama2* specifies the .l2 syntax for describing APIs, and implements a CLI to execute .l2 files. Engineers can collaborate on `.l2` files using version control. *Lama2* integrates nicely with
IDEs and text-editors as well. Think of *Lama2* as **markdown for APIs.**

## Benefits

1. **Plain-text files:** Store APIs in plain-text `.l2` files. Simple and human-friendly syntax. Learn basics within  minutes!
1. **Simple CLI:** Launch the CLI tool `l2` on `*.l2` files to make REST API requests.
1. **Editor support:** Invoke *Lama2* from your favorite text editor or IDE. Friendly API to build editor extensions.
1. **Longevity & trackability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share API repo with teammates and colleagues to collaborate
1. **Documentation:** Explore tutorials, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a recursive descent parser, based on a clearly specified grammar. Dig into details and implement new syntax (ex: to support `websockets`)

## Example Demo: A POST request

![](demo.gif)

#### `basic_post.l2`:

```bash
POST
https://httpbin.org/post
{
    "hello": "world"
}
```

#### Execution:

```
l2 basic_post.l2
```