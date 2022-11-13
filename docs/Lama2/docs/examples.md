# Examples

The following examples provide a sampling of
the various types of requests *Lama2* handles
presently. Execute each file as:

```
l2 <file_name>.l2
```

You can also clone the repo and open it up in
VSCode, install the *Lama2* extension and fire
requests from there.

## GET request

```
GET
https://httpbin.org/get
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0000_sample_get.l2)

## JSON POST request

One can dump the JSON body at the end of an `.l2` file to create a POST request:

```
POST
https://httpbin.org/post

{
    "a": "b",
    "c": "d"
}
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0002_sample_post.l2)

## JSON POST in VarJSON format

Make a POST request with JSON body specified
as `key=value`. *Lama2* converts the input into
a corresponding JSON value `{"a": "b", "c": "d"}`. We call the `key=value` format *VarJSON*. This example produces an effect identical to the [previous one](#basic-json-post)

```
POST
https://httpbin.org/post

a=b
c=d
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0001_sample_post_varjson.l2)

## Comments

One can start a comment anywhere in the file
with the `#` character.

```
# Pound symbol signifies a comment
POST
https://httpbin.org/post

a=b # Comments may start at the end of lines as well
c=d

# Comments work even after the payload
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0003_comment.l2)

## Environment Variables: Switch base URL

Specify variables in `l2.env` and then load
them up in the API files. Presently, **the `l2.env` file should reside in the same directory as the `.l2` API file.**

#### l2.env

```
export LOCAL="http://localhost:8000"
export REMOTE="http://httpbin.org"
```

#### env_example.l2
```
POST
${REMOTE}/post

{
    "lorem": "ipsum"
}
```

Get [Source Files](https://github.com/HexmosTech/Lama2/tree/main/examples/0004_env_switch_root)

## Headers

Use `key:value` format to specify headers. **Headers must come before the data section**. 

Specify strings for key/value in three ways:

1. Double quoted (`"hello"`)
1. Single quoted (`'hello'`)
1. Unquoted (`hello`)

```
POST 
https://httpbin.org/post

# HEADERS
X-Parse-Application-Id:'helloworld'
X-Parse-REST-API-Key:"byeworld"

# DATA
a="b"  # double-quoted string
'c'=d  # single-quoted & unquoted strings
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0005_headers_simple.l2)

## Send cookies in header

Headers represent cookies in *Lama2*. Just specify cookie key value pairs separated by
`=` within the header value as shown.
```
POST 
https://httpbin.org/post

# HEADERS
Cookie:"sessionid=foo;another-cookie=bar"

# DATA
hello=world
```
Get [Source File](https://github.com/HexmosTech/Lama2/tree/main/examples/0006_cookies.l2)

## Fill forms & attach files with MULTIPART

Use the `MULTIPART` keyword after the HTTP
verb to enable forms and file attachments.

The data section may contain any number of 
form inputs using the `key=value` syntax.

Following the data section, one can specify
any number of files in the form of `<field_name>@<file_path>`. The file path is
relative to the API file.

```
POST
MULTIPART
http://httpbin.org/post

'X-Parse-Application-Id':hello 
X-Parse-REST-API-Key:"world"

# DATA
first=second

# FILES
myfile@./image.jpeg
```
Get [Source Files](https://github.com/HexmosTech/Lama2/tree/main/examples/0007_multipart_file)

## Image as Base64 encoded JSON field

We can embed images (or other files) as 
base64 strings in JSON using *Lama2*.

First, we define a `PHOTO` variable, loaded
up with the results of the `base64` command.

### l2.env

```
export PHOTO=`base64 image.jpeg`
```

Next, we refer to the `PHOTO` variable in
the API file. Pay special attention to the
quoting mechanism `"'{PHOTO}'"`. **The quoting
must look exactly as shown in the following
template for the request to work correctly.**

### base64_embed.l2

```
POST
http://httpbin.org/post

{
	"imageb64_field": "'${PHOTO}'",
}
```