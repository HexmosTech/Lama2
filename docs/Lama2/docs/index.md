# **`Lama2`**: Free, Lightweight & Collaborative API Client

![](./banner.png)

![Build](https://github.com/HexmosTech/Lama2/actions/workflows/release.yml/badge.svg)
![Test](https://github.com/HexmosTech/Lama2/actions/workflows/testapp.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/HexmosTech/lama2)](https://goreportcard.com/report/github.com/HexmosTech/lama2)

<a href="https://www.producthunt.com/posts/hexmos-lama2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-hexmos&#0045;lama2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=370586&theme=light" alt="Hexmos&#0032;Lama2 - Plain&#0045;Text&#0032;Powered&#0032;REST&#0032;API&#0032;Client&#0032;for&#0032;Teams | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

## The Problem

Most engineering teams have trouble collaborating on APIs or end up paying too much for Postman/Insomnia. Our 100% FOSS product _Lama2_ helps your team collaborate on APIs for free based on git, so that you can save money and avoid friction.

## About

_Lama2_ is a Plain-Text powered REST API client & manager built for serious engineering teams.
_Lama2_ specifies the [`.l2` syntax](tutorials/examples.md) for describing APIs, and implements a [CLI](tutorials/installation.md) to execute `.l2` files. Engineers [collaborate](tutorials/collaboration.md) on `.l2` files using version control. _Lama2_ [integrates](tutorials/editor.md) nicely with
IDEs and text-editors as well. Think of _Lama2_ as [**Markdown for APIs.**](reference/philosophy.md)

## Benefits

1. **Plain-Text files:** Store APIs in the Plain-Text `.l2` API files. [Simple](explanation/l2format.md) and [human-friendly](explanation/syntax.md) syntax. Learn [basics](tutorials/examples.md) within minutes!
1. **Simple CLI:** Launch the [CLI](tutorials/installation.md) tool `l2` on API files to make REST API requests.
1. **Editor support:** Invoke _Lama2_ from your [favorite text editor](tutorials/installation.md#from-vs-code) or IDE. Helpful documentation and tool support to [build editor extensions](tutorials/editor.md) included.
1. **Longevity & track-ability:** Commit `.l2` files to `git` or other version control for long life & change tracking.
1. **Collaboration:** Share [API repo](tutorials/collaboration.md) with teammates and colleagues to collaborate
1. **Code generation:** [Convert](tutorials/codegen.md) `l2` API definitions into your preferred language/library, be it Python, Javascript, Java or many other options.
1. **Powerful chaining:** Chain multiple `l2` requests through embedded [Javascript blocks](tutorials/examples.md#chain-requests-using-javascript)
1. **Documentation:** [Explore](index.md) examples, how-tos, explanations, references, FAQ/RAQs, and diagrams. Documentation is a priority, so that you never have to get lost.
1. **Extensibility:** `.l2` syntax is implemented as a _recursive descent parser_, based on a [formal grammar](reference/grammar.md). [Dig into details](reference/parser.md) and [implement](https://github.com/HexmosTech/Lama2/tree/main/parser) new syntax (ex: to support `websockets`)

!!! tip

    **Coming from postman?** Let *Lama2* help you [convert](tutorials/postman.md) Postman
    data dumps into a  nice  Plain-Text Lama2 repo.

## Terminal Demo: A POST request

![](https://hexmos.com/lama2/demo.gif)

**`basic_post.l2`:**

<pre id="liveapi-element">
POST
https://httpbin.org/post
{
    "hello": "world"
}

---

_METADATA={
    "api_name": "Basic POST",
    "sample_response": {
        "args": {},
        "data": "{\"hello\":\"world\"}",
        "files": {},
        "form": {},
        "headers": {
            "Accept": "*/*",
            "Accept-Encoding": "gzip, deflate",
            "Content-Length": "17",
            "Content-Type": "application/json",
            "Host": "httpbin.org",
            "User-Agent": "python-requests/2.25.1",
            "X-Amzn-Trace-Id": "Root=1-6686d0e6-336395f75a7f8a8b474a7563"
        },
        "json": {
            "hello": "world"
        },
        "origin": "20.235.244.147",
        "url": "https://httpbin.org/post"
    }
}
</pre>

**Execute:**

```
l2 basic_post.l2
```

## VSCode Demo: The same POST request

![](https://hexmos.com/lama2/demo2.gif)

## Community and Support

1. [Full documentation @hexmos.com/lama2](https://hexmos.com/lama2/)
1. Report issues and propose improvements at [Github issues](https://github.com/HexmosTech/Lama2/issues)
1. Discuss with the community at [Discord](https://discord.gg/zTmxXA6F)
