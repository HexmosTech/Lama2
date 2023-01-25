# -*- coding: utf-8 -*- #

module Rouge
  module Lexers
    class Lama2 < RegexLexer
      title     "lama2"
      tag       'Lama2'
      mimetypes 'text/x-lama2'
      filenames '*.l2', '*.lama2'

      state:root do
          rule /(^\s*$)/, String::Doc
          rule /(^---$)/, Punctuation
          rule /(^\s*(GET|POST|PUT|DELETE|get|post|put|delete))/, Keyword, :root_lama
          rule /(^\s*#.*)/, Comment, :root_lama
          rule /(?!(^\s*(#|GET|POST|PUT|DELETE|get|post|put|delete)))/, Name::Variable, :root_javascript
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:js_string1__1 do
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:js_string2__1 do
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:js_string3__1 do
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:json__1 do
          rule /(\{)/, Punctuation, :json__1
          rule /(\[)/, Punctuation, :json__2
          rule /(,|\:)/, Punctuation
          rule /(\")/, Punctuation, :json__3
          rule /((?:\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\.[0-9]*)(?:(?:e|E)(?:\+|\-)[0-9]+)?)?)/, Keyword::Type
          rule /(true|false|null)/, Keyword::Type
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:json__2 do
          rule /(\{)/, Punctuation, :json__1
          rule /(\[)/, Punctuation, :json__2
          rule /(,|\:)/, Punctuation
          rule /(\")/, Punctuation, :json__3
          rule /((?:\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\.[0-9]*)(?:(?:e|E)(?:\+|\-)[0-9]+)?)?)/, Keyword::Type
          rule /(true|false|null)/, Keyword::Type
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:json__3 do
          rule /((?:\\u[0-9a-f]{4})|(?:\\["\\\/bfnrt]))/, String::Escape
          rule /([^\"\\\n\r]+)/, String::Char
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:multi_line_comment__1 do
          rule /(\n|\r|\r\n)/, String
          rule /./, Comment
      end

      state:root_javascript do
          rule /(^---$)/, Punctuation
          rule /(\/\*)/, Comment, :multi_line_comment__1
          rule /(\/\/.*$)/, Comment
          rule /(\`)/, Punctuation, :js_string3__1
          rule /(")/, Punctuation, :js_string1__1
          rule /(')/, Punctuation, :js_string2__1
          rule /([.,\\/#!$%\^&\*;:{}=\-_\[\]`~()])/, Punctuation
          rule /(?:\b)(\.[_$a-zA-Z\xa0-\uffff][_$a-zA-Z0-9\xa0-\uffff]*)(\()/ do
            groups Name::Function, Punctuation
          end
          rule /(?:\b)(?:do|if|in|for|let|new|try|var|case|else|enum|eval|false|null|this|true|void|with|break|catch|class|const|super|throw|while|yield|delete|export|import|public|return|static|switch|typeof|default|extends|finally|package|private|continue|debugger|function|arguments|interface|protected|implements|instanceof)(?:\b)/, Keyword::Declaration
          rule /(?:\b)(var|let|const)(?:\b)/, Keyword::Namespace
          rule /(?:\b)(return|yield)(?:\b)/, Keyword::Reserved
          rule /(?:\b)([_$a-zA-Z\xa0-\uffff][_$a-zA-Z0-9\xa0-\uffff]*)(?:\b)/, Name::Variable
          rule /(\+|-|\*|\\/|=|>|<|>=|<=|&|\||%|!|\^|\(|\))/, Punctuation
          rule /(\b\d+)/, Number
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

      state:root_lama do
          rule /(^---$)/, Punctuation
          rule /(#.*)/, Comment
          rule /(GET|POST|PUT|DELETE|get|post|put|delete)/, Keyword
          rule /(MULTIPART|multipart)/, Keyword
          rule /(https?|\$\{.*\})([^\s\n\r]+)(?:\b)/ do
            groups Name::Variable, Punctuation
          end
          rule /(\$\{.*\})(.*)(?:\b)/, Punctuation
          rule /(\{)/, Punctuation, :json__1
          rule /(\[)/, Punctuation, :json__2
          rule /(,|\:)/, Punctuation
          rule /(\")/, Punctuation, :json__3
          rule /((?:\-?(?:0|(?:[1-9][0-9]*)))(?:(?:\.[0-9]*)(?:(?:e|E)(?:\+|\-)[0-9]+)?)?)/, Keyword::Type
          rule /(true|false|null)/, Keyword::Type
          rule /(.*)(:)([^#\n\r]+)/ do
            groups Name::Constant, Punctuation, Literal
          end
          rule /(.*)(=)([^#\n\r]+)/ do
            groups Name::Constant, Punctuation, Literal
          end
          rule /(.*)(@)([^#\n\r]+)/ do
            groups Name::Constant, Punctuation, Literal
          end
          rule /(\n|\r|\r\n)/, String
          rule /./, String
      end

    end
  end
end

