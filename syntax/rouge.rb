# -*- coding: utf-8 -*- #

module Rouge
  module Lexers
    class Lama2 < RegexLexer
      title     "lama2"
      tag       'Lama2'
      mimetypes 'text/x-lama2'
      filenames '*.l2', '*.lama2'

      state:root do
          rule /(#.*)/, Comment
          rule /(GET|POST|PUT|DELETE|get|post|put|delete)/, Keyword
          rule /(MULTIPART|multipart)/, Keyword
          rule /(https?|\$\{.*\})(.*)/ do
            groups Name::Variable, Punctuation
          end
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

    end
  end
end

