## Useful Options 
The `l2` command provides some helpful options for
extension developers. The options are:

1. `--nocolor` or `-n` disables colored output in HTTPie;
IDE/Editor plugin.
1. `--output=<target.json` or `-o` writes a structured JSON
output to the target file. The following is the content
structure:
```json
{
    "logs": ...,
    "headers": ...,
    "body": ...
}
```
Right now, all the three values are strings. In the future, we may transform the values further to provide a more parse-friendly structure.

## The command
Combining the above two options, we get:

```bash
l2 -n -o /tmp/lama2.json my_api.l2
```

The command mentioned above disables HTTPie colors,
writes the whole transaction to a structured JSON, 
while also printing details into `stdout`.

The extension author can simply read the file, and
display the contents to users appropriately. For an
example, see [Lama2 for VSCode](https://github.com/HexmosTech/Lama2Code)
(also see [Marketplace page](https://marketplace.visualstudio.com/items?itemName=hexmos.Lama2)).