*Lama2* is capable of converting `.l2` files into functional
code of your preferred language and library.

For example, to generate code in *python requests*, use the following command:

```
l2 -c python.requests myfile.l2
```

To pick the default library in a language, omit the library as follows:


```
l2 -c python myfile.l2
```

## Languages and libraries supported

*  shell
    *  curl (default)
    *  httpie
    *  wget
*  powershell
    *  webrequest (default)
    *  restmethod
*  ocaml
    *  cohttp (default)
*  csharp
    *  restsharp (default)
    *  httpclient
*  r
    *  httr (default)
*  php
    *  curl (default)
    *  guzzle
    *  http1
    *  http2
*  ruby
    *  native (default)
*  clojure
    *  clj_http (default)
*  java
    *  unirest (default)
    *  asynchttp
    *  nethttp
    *  okhttp
*  http
    *  1.1 (default)
    *  http1.1
*  swift
    *  nsurlsession (default)
*  node
    *  native (default)
    *  request
    *  unirest
    *  axios
    *  fetch
*  c
    *  libcurl (default)
*  go
    *  native (default)
*  python
    *  python3 (default)
    *  requests
*  kotlin
    *  okhttp (default)
*  javascript
    *  xhr (default)
    *  axios
    *  fetch
    *  jquery
*  objc
    *  nsurlsession (default)
