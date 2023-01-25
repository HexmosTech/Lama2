Find in this page an informal description of
the rules for authoring `.l2` API files. This
document expects some familiarity with *Lama2*.

To quickly get started with *Lama2*, head over
to [Examples](../tutorials/examples.md). 

On the
other hand, if you are a developer and wish to
learn more about the formal grammar underlying
*l2*, visit the [Grammar](../reference/grammar.md) 
section.

### Comments start with `#`

Lines starting with `#` are comments and hence ignored altogether

### All HTTP Verbs supported - including the common GET/POST/PUT

Fully supported: `GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH`

### JSON is the default submission type, but MULTIPART is supported too


#### `varjson` is a simpler syntax to specify flat JSONs

`varjson` values are defined as follows:

```
hello=world
foo=bar
```

The above results in a JSON submission of the form:

```
{
	"hello": "world",
	"foo": "bar"
}
```

#### Nested JSON can simply be dumped at the end of the document

The JSON recognition engine is quite lenient. It can deal with 
minor errors in the format (such as having single quotes instead
of double quotes, trailing garbage, or an extra comma after the
last element in an array,). 

```
POST
https://httpbin.org/post

{
    "a": "b",
    "c": "d"
}
```



#### MULTIPART allows both file uploads & the usual fields

Example:

```
POST
MULTIPART
http://localhost:8000/register
userid=lince5
file@./helloworld.jpg
```

!!! Note

    The *file path is relative to the request file.*

###  Cookies are sent as headers

Cookies are specified in a `Cookie` header as follows:

```
Cookie:'sessionid=foo;another-cookie=bar'
```

### Environments variables/commands can be defined in `<requests_dir>/l2.env`

By default, *l2* looks for a `l2.env` file in the same directory as the given
request file directory. Example `l2.env`:

```
export PHOTO=`base64 aadhaarlarge.jpg`
export AHOST="http://localhost:8000"
```

#### The environment file can load results of commands

Use the backtick notation `\`command\`` to place the results of
commands into environment variables:

```
export PHOTO=`base64 image.jpeg`
```

One can load the `PHOTO` variable in API files.

### Chain requests through Javascript blocks

*Lama2* supports plain Javascript (JS) blocks
as a glue for manipulating responses and passing on
values to later stages. At a higher
level, a chain of requests may look like:

```
Javascript 1
---
L2 Request 1
---
Javscript 2
---
L2 Request 2
```

The triple-dash (`---`) separator is mandatory. The special
variable `result` contains the response from previous stages.

For example, in the above case, `Javascript 2` can access the response from `L2 Request 1` through the `result` variable.

Learn more about request chaining in [Examples](../tutorials/examples.md#chain-requests-using-javascript).