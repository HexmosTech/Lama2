# 🧝 `httpelf`: text-powered REST API client

`httpelf` is a text-file based REST API client & manager. It uses
the nascent `.http` file format to represent http(s) requests &
is capable of making pure text-powered client requests.

[[_TOC_]]

## General Usage

```
$ elf <request>.http
```

Find some sample request files in the 

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
1. Start defining `*.http` files for each service specific API. 
1. Use `elf file.http` to test the newly defined APIs. 
1. Push the API files once ready into APIHub repo

## `.http` format

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

#### Nested JSON can simply be dumped into the document at any place

Check the `sample_http` directory for numerous examples of this format.

####  Cookies are sent as headers

Cookies are specified in a `Cookie` header as follows:

```
Cookie:'sessionid=foo;another-cookie=bar'
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

Note that *file path is relative to the request file.*

### Environments variables/commands can be defined in `<requests_dir>/elf.env`

By default, `elf` looks for a `elf.env` file in the same directory as the given
request file directory. Example `elf.env`:

```
export PHOTO=`base64 aadhaarlarge.jpg`
export AHOST="http://localhost:8000"
```

#### Base64 of images can be obtained through an environment command

In the previous section, we defined the `PHOTO` variable as a shell command.
The results of the command can be used in a requests file as follows:

```
POST
${AHOST}/aadhaarscanjson

{
	"aadhaarfileb64": "'${PHOTO}'",
	"coords": "130,327,539,331,564,607,107,617",
	"uemail": "shrijith.sv@gmail.com"
}
```

Note that within a nested JSON, it is **madatory* to use single quotes around
a variable expansion. This is present limitation of the `elf` system. It is 
ideal to avoid this syntax, however given time/energy constraints we are not 
picking up the issue. Please use the above as a template in the interim.


### Find examples in `sample_http` folder

Please explore `sample_http` folder to see some samples, and get a feel for
how things are defined withing `.http` files.