from pygments.lexer import RegexLexer, bygroups
from pygments.token import *

import re

__all__=['Lama2Lexer']

class Lama2Lexer(RegexLexer):
    name = 'Lama2'
    aliases = ['lama2']
    filenames = ['*.l2', '*.lama2']
    flags = re.MULTILINE | re.UNICODE

    tokens = {
        'root' : [
            (u'(#.*)', bygroups(Comment)),
            (u'(GET|POST|PUT|DELETE|get|post|put|delete)', bygroups(Keyword)),
            (u'(MULTIPART|multipart)', bygroups(Keyword)),
            (u'(https?|\\$\\{.*\\})(.*)', bygroups(Name.Variable, Punctuation)),
            (u'(\\{)', bygroups(Punctuation), 'json__1'),
            (u'(\\[)', bygroups(Punctuation), 'json__2'),
            (u'(,|\\:)', bygroups(Punctuation)),
            (u'(\\\")', bygroups(Punctuation), 'json__3'),
            (u'((?:\\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\\.[0-9]*)(?:(?:e|E)(?:\\+|\\-)[0-9]+)?)?)', bygroups(Keyword.Type)),
            (u'(true|false|null)', bygroups(Keyword.Type)),
            (u'(.*)(:)([^#\\n\\r]+)', bygroups(Name.Constant, Punctuation, Literal)),
            (u'(.*)(=)([^#\\n\\r]+)', bygroups(Name.Constant, Punctuation, Literal)),
            (u'(.*)(@)([^#\\n\\r]+)', bygroups(Name.Constant, Punctuation, Literal)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'json__1' : [
            (u'(\\{)', bygroups(Punctuation), 'json__1'),
            (u'(\\[)', bygroups(Punctuation), 'json__2'),
            (u'(,|\\:)', bygroups(Punctuation)),
            (u'(\\\")', bygroups(Punctuation), 'json__3'),
            (u'((?:\\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\\.[0-9]*)(?:(?:e|E)(?:\\+|\\-)[0-9]+)?)?)', bygroups(Keyword.Type)),
            (u'(true|false|null)', bygroups(Keyword.Type)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'json__2' : [
            (u'(\\{)', bygroups(Punctuation), 'json__1'),
            (u'(\\[)', bygroups(Punctuation), 'json__2'),
            (u'(,|\\:)', bygroups(Punctuation)),
            (u'(\\\")', bygroups(Punctuation), 'json__3'),
            (u'((?:\\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\\.[0-9]*)(?:(?:e|E)(?:\\+|\\-)[0-9]+)?)?)', bygroups(Keyword.Type)),
            (u'(true|false|null)', bygroups(Keyword.Type)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'json__3' : [
            (u'((?:\\\\u[0-9a-f]{4})|(?:\\\\[\"\\\\/bfnrt]))', bygroups(String.Escape)),
            (u'([^\\\"\\\\\\n\\r]+)', bygroups(String.Char)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ]
    }

