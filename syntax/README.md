# Syntax Highlighting for `Lama2`

We use [Iro](https://eeyo.io/iro/documentation/) for generating
syntax grammars for multiple text editors in one go.

## Source Grammar

See [Lama2.iro](./lama2.iro) for the source syntax grammar.

## Generated Grammars

### VSCode

Source file: [textmate.plist](./textmate.plist)

Iro generates a Textmate syntax file in the `Plist` format. We
must convert the `Plist` [file](./textmate.plist) into JSON using [wtools](https://wtools.io/convert-plist-to-json). If the converted output is wrapped up in a list `[]`, then remove that part. After removing
unnecessary junk, we will have a nice syntax JSON of the form `{}`. See the VSCode plugin, particularly [package.json](https://github.com/HexmosTech/Lama2Code/blob/main/package.json#L38) and [syntaxes/lama2.tmGrammar.json](https://github.com/HexmosTech/Lama2Code/blob/main/syntaxes/lama2.tmGrammar.json)

### Other

- [Atom](./atom)
- [Pygments](./pygments.py)
- [Rouge](./rouge.rb)
- [CSS](./style.css)
- [Sublime3](./sublime3)

## Using Iro

For interative usage, head over to [iro web](https://eeyo.io/iro/). 

Post the source file [lama2.iro](./lama2.iro) into the editor.

Load up a sample lama2 doc, and hit the play button.

### Iro learning resources

Following are some other useful resources related to Iro:

- Open source CLI version: https://github.com/c272/iro4cli
- Tutorial from author: https://medium.com/@model_train/creating-universal-syntax-highlighters-with-iro-549501698fd2
- Iro Documentation: https://eeyo.io/iro/documentation/