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
            (u'(^\\s*$)', bygroups(String.Doc)),
            (u'(^---$)', bygroups(Punctuation)),
            (u'(^\\s*(GET|POST|PUT|DELETE|get|post|put|delete))', bygroups(Keyword), 'root_lama'),
            (u'(^\\s*#.*)', bygroups(Comment), 'root_lama'),
            (u'(?!(^\\s*(#|GET|POST|PUT|DELETE|get|post|put|delete)))', bygroups(Name.Variable), 'root_javascript'),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'js_string1__1' : [
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'js_string2__1' : [
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'js_string3__1' : [
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
        ], 
        'multi_line_comment__1' : [
            ('(\n|\r|\r\n)', String),
            ('.', Comment),
        ], 
        'root_javascript' : [
            (u'(^---$)', bygroups(Punctuation)),
            (u'(/\\*)', bygroups(Comment), 'multi_line_comment__1'),
            (u'(//.*$)', bygroups(Comment)),
            (u'(\\`)', bygroups(Punctuation), 'js_string3__1'),
            (u'(\")', bygroups(Punctuation), 'js_string1__1'),
            (u'(\')', bygroups(Punctuation), 'js_string2__1'),
            (u'([.,\\/#!$%\\^&\\*;:{}=\\-_\\[\\]`~()])', bygroups(Punctuation)),
            (u'(?:\\b)(\\.[_$a-zA-Z\\xa0-\uffff][_$a-zA-Z0-9\\xa0-\uffff]*)(\\()', bygroups(Name.Function, Punctuation)),
            (u'(?:\\b)(?:do|if|in|for|let|new|try|var|case|else|enum|eval|false|null|this|true|void|with|break|catch|class|const|super|throw|while|yield|delete|export|import|public|return|static|switch|typeof|default|extends|finally|package|private|continue|debugger|function|arguments|interface|protected|implements|instanceof)(?:\\b)', bygroups(Keyword.Declaration)),
            (u'(?:\\b)(var|let|const)(?:\\b)', bygroups(Keyword.Namespace)),
            (u'(?:\\b)(return|yield)(?:\\b)', bygroups(Keyword.Reserved)),
            (u'(?:\\b)([_$a-zA-Z\\xa0-\uffff][_$a-zA-Z0-9\\xa0-\uffff]*)(?:\\b)', bygroups(Name.Variable)),
            (u'(\\+|-|\\*|\\/|=|>|<|>=|<=|&|\\||%|!|\\^|\\(|\\))', bygroups(Punctuation)),
            (u'(\\b\\d+)', bygroups(Number)),
            ('(\n|\r|\r\n)', String),
            ('.', String),
        ], 
        'root_lama' : [
            (u'(^---$)', bygroups(Punctuation)),
            (u'(#.*)', bygroups(Comment)),
            (u'(GET|POST|PUT|DELETE|get|post|put|delete)', bygroups(Keyword)),
            (u'(MULTIPART|multipart)', bygroups(Keyword)),
            (u'(https?|\\$\\{.*\\})([^\\s\\n\\r]+)(?:\\b)', bygroups(Name.Variable, Punctuation)),
            (u'(\\$\\{.*\\})(.*)(?:\\b)', bygroups(Punctuation)),
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
        ]
    }

