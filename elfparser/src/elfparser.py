#!/usr/bin/env python3

from audioop import mul
from email.quoprimime import unquote
from .parser import *
import json
import unicodedata
import math


class ElfParser(Parser):
    """
    HttpFile ::= HttpVerb Multipart? Url Details?
    Details ::= ( (DataInput Headers?) | (Headers DataInput?) )
    DataInput ::= (Json | Varjson) FileRefs?
    HttpVerb ::= "GET" | "HEAD" | "POST" | "PUT" | "DELETE" | "CONNECT" | "OPTIONS" | "TRACE" | "PATCH"
    Multipart ::= "MULTIPART"
    Varjson ::= (StringType "=" StringType)+ FileRefs*
    Headers ::= (StringType ":" StringType)
    Url ::= "http" "s"? ":://" U>nquoted+
    Json ::= AnyType
    Start ::= AnyType
    AnyType ::= ComplexType | PrimitiveType
    ComplexType ::= List | Map
    List ::= "[" (AnyType ",")* AnyType? "]"
    Map ::= "{" (Pair ",")* Pair? "}"
    Pair ::= (QuotedString | UnquotedString) ":" AnyType
    PrimitiveType ::= Null | Boolean | QuotedString | Number
    Number ::= Integer | Fraction | Exponent
    Integer ::= Digit | Onenine Digits | '-' Digit | '-' Onenine Digits
    Digits ::= Digit | Digit Digits
    Digit ::= '0' | Onenine
    OneNine ::= [1-9]
    Fraction ::= '.' Digits | ""
    Exponent ::= 'E' Sign Digits | 'e' Sign Digits | ""
    Sign ::= '+' | '-' | ""
    """
    
    def __init__(self):
        super().__init__()
        self.context = {}
    
    def eat_whitespace(self):
        is_processing_comment = False

        while self.pos < self.len:
            char = self.text[self.pos + 1]
            if is_processing_comment:
                if char == "\n":
                    is_processing_comment = False
            else:
                if char == "#":
                    is_processing_comment = True
                elif char not in " \f\v\r\t\n":
                    break

            self.pos += 1

    def start(self):
        return self.match("http_file")
    
    def http_file(self):
        verb = self.match("http_verb")
        multipart = self.maybe_match("multipart")
        if multipart is not None:
            self.context["multipart"] = True
        url = self.match("the_url")
        details = self.maybe_match("details")
        
        res = {
            "verb": verb,
            "url": url,
        }
        
        if details is not None:
            res["details"] = details
            
        if multipart is not None:
            res["multipart"] = True
            
            
        return res
    
    def multipart(self):
        return self.keyword("multipart", case_insensitive=True)
        
    def http_verb(self):
        verb_list = {"GET": "get", "HEAD": "head", "POST": "post", "PUT": "put", "DELETE": "delete", "CONNECT": "connect", "OPTIONS": "options", "TRACE": "trace", "PATCH": "patch"}
        
        for k in verb_list.keys():
            r = self.maybe_keyword(k)
            if r is not None:
                return verb_list[r]
            
        for k in verb_list.values():
            r = self.maybe_keyword(k)
            if r is not None:
                return r

        raise ParseError(
            self.pos + 1, "Couldn't find http verb (any of GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCh)", self.text[self.pos + 1]
        )
        
    def details(self):
        res = self.maybe_match('header_data')
        if res is not None:
            return res
        return self.match("data_header")
        
    def data_header(self):
        ip_data =  self.match("data_input")
        op = {"ip_data": ip_data}
        headers = self.maybe_match("headers")
        if headers is not None:
            op["headers"] = headers
        return op

    def header_data(self):
        headers =  self.match("headers")
        op = {"headers": headers}
        ip_data = self.maybe_match("data_input")
        if ip_data is not None:
            op["ip_data"] = ip_data
        return op
    
    def headers(self):
        head_dict = {}
        key, value = self.match('header_pair')
        head_dict[key] = value
        
        while True:
            res = self.maybe_match('header_pair')
            if res is None:
                break
            
            key, value = res
            head_dict[key] = value
            
        return head_dict
    
    def data_input(self):
        res = self.maybe_match("varjson")
        if res is not None:
            return res
        return self.match("any_type")
    
    def varjson(self):
        varjson_dict = {}
        if (not 'multipart' in self.context) or (self.context['multipart'] == False):
            key, value = self.match('varjson_pair')
            varjson_dict[key] = value
        
        while True:
            res = self.maybe_match('varjson_pair')
            if res is None:
                break
            
            key, value = res
            varjson_dict[key] = value
            
        if 'multipart' in self.context and self.context['multipart']:
            while True:
                res = self.maybe_match('files_pair')
                if res is None:
                    break
                
                if not '@files' in varjson_dict:
                    varjson_dict['@files'] = {}
                
                key, value = res
                varjson_dict['@files'][key] = value

        if len(varjson_dict.keys()) == 0:
            raise ParseError(
                self.pos + 1, "No varjson match", self.text[self.pos + 1]
            )
            
        return varjson_dict
        
        
    def the_url(self):
        res = []
        res.append(self.keyword("http", eat_ws_end=False))
        r = self.maybe_char("s")
        if r is not None:
            res.append(r)
        res.append(self.keyword("://", eat_ws_start=False, eat_ws_end=False))
        while True:
            r = self.maybe_char("A-Za-z0-9-._~:/?#[]@!$&'()*+,;%=")
            if r is not None:
                res.append(r)
            else:
                break
            
        return ''.join(res)
    

    def any_type(self):
        return self.match("complex_type", "primitive_type")

    def primitive_type(self):
        return self.match("null", "boolean", "quoted_string", "number")

    def complex_type(self):
        return self.match("list", "map")

    def list(self):
        rv = []

        self.keyword("[")
        comma = False
        while True:
            item = self.maybe_match("any_type")
            try:
                # print("item = ", item.encode('utf-8', 'replace').decode())
                pass
            except:
                # print(item)
                pass
            if item is None:
                break

            rv.append(item)

            if not self.maybe_keyword(","):
                comma = False
                break
            else:
                comma = True

        self.keyword("]")
        # if comma == True:
        #     raise ParseError(
        #         self.pos, "Unnecessary trailing comma found", self.text[self.pos]
        #     )
        return rv

    def map(self):
        rv = {}

        self.keyword("{")
        comma = False
        while True:
            item = self.maybe_match("pair")
            if item is None:
                break

            rv[item[0]] = item[1]

            if not self.maybe_keyword(","):
                comma = False
                break
            else:
                comma = True

        self.keyword("}")
        # if comma == True:
        #     raise ParseError(
        #         self.pos, "Unnecessary trailing comma found", self.text[self.pos]
        #     )

        return rv

    def header_pair(self):
        key = self.maybe_match("quoted_string")
        if key is None:
            key = self.match('unquoted')

        if type(key) is not str:
            raise ParseError(
                self.pos + 1, "Expected string but got number", self.text[self.pos + 1]
            )

        self.keyword(":")
        value = self.maybe_match("quoted_string")
        if value is None:
            value = self.match('unquoted')

        return key, value

    def varjson_pair(self):
        key = self.maybe_match("quoted_string")
        if key is None:
            key = self.match('varjson_unquoted')
            
        if type(key) is not str:
            raise ParseError(
                self.pos + 1, "Expected string but got number", self.text[self.pos + 1]
            )

        self.keyword("=")
        value = self.maybe_match("quoted_string")
        if value is None:
            value = self.match('varjson_unquoted')



        return key, value

    def files_pair(self):
        key = self.maybe_match("quoted_string")
        if key is None:
            key = self.match('files_unquoted')

        if type(key) is not str:
            raise ParseError(
                self.pos + 1, "Expected string but got number", self.text[self.pos + 1]
            )

        self.keyword("@")
        value = self.maybe_match("quoted_string")
        if value is None:
            value = self.match('files_unquoted')

        return key, value

    def pair(self):
        key = self.match("quoted_string")

        if type(key) is not str:
            raise ParseError(
                self.pos + 1, "Expected string but got number", self.text[self.pos + 1]
            )

        self.keyword(":")
        value = self.match("any_type")

        return key, value

    def null(self):
        self.keyword("null")
        return "null"

    def boolean(self):
        boolean = self.keyword("true", "false")
        return boolean[0] == "t"

    def number(self):
        intpart = self.match("integer")
        fracpart = self.match("fraction")
        expart = self.match("exponent")
        if expart != 0:
            return (intpart + fracpart) * math.pow(10, expart)
        else:
            return intpart + fracpart

    def integer(self):
        return int(
            self.match(
                "integer_rule_4", "integer_rule_3", "integer_rule_2", "integer_rule_1"
            )
        )

    def integer_rule_1(self):
        return self.match("digit")

    def integer_rule_2(self):
        s = []
        ionenine = self.match("onenine")
        if ionenine is not None:
            s.append(ionenine)

        idigits = self.match("digits")
        if idigits is not None:
            s.append(idigits)

        return "".join(s)

    def integer_rule_3(self):
        s = []
        nsign = self.char("-")
        if nsign is not None:
            s.append(nsign)
        r = self.match("integer_rule_1")
        s.append(r)
        return "".join(s)

    def integer_rule_4(self):
        s = []
        nsign = self.char("-")
        if nsign is not None:
            s.append(nsign)
        r = self.match("integer_rule_2")
        s.append(r)
        return "".join(s)

    def fraction(self):
        try:
            return float(self.match("fraction_rule_1"))
        except ParseError:
            return 0

    def fraction_rule_1(self):
        s = []
        s.append(self.char("."))
        s.append(self.match("digits"))
        return "".join(s)

    def exponent(self):
        try:
            return self.match("exponent_rule_1")
        except ParseError:
            return 0

    def exponent_rule_1(self):
        self.char("eE")
        esign = self.match("sign")
        multiplier = None
        if esign == "+" or esign == "":
            multiplier = 1
        else:
            multiplier = -1
        return multiplier * int(self.match("digits"))

    def digits(self):
        s = []
        r = self.match("digit")
        s.append(r)
        while True:
            r = self.maybe_match("digit")
            if r is None:
                break
            s.append(r)
        return "".join(s)

    def digit(self):
        s = self.maybe_char("0")
        if not s is None:
            return s

        return self.match("onenine")

    def onenine(self):
        s = self.char("1-9")
        return s

    def sign(self):
        s = self.maybe_char("+-")
        if s is None:
            return ""
        else:
            return s
        
    def string_type(self):
        uq = self.maybe_match('unquoted')
        if uq is not None:
            return uq
        return self.match('quoted')
    
    def varjson_unquoted(self):
        # compared to the standard `unquoted`, we do not match
        # '=' as part of the unquoted string. Why? Varjson pair 
        # is separated by `=` symbol
        # TODO: Avoid duplication by taking in arguments to standard
        # `unquoted` function
        acceptable_chars = '@0-9A-Za-z \t!$%&()*+./;<>?^_`|~-'

        chars = [self.char(acceptable_chars)]

        while True:
            char = self.maybe_char(acceptable_chars)
            if char is None:
                break

            if char in 'Ee.':
                number_type = float

            chars.append(char)

        rv = ''.join(chars).rstrip(' \t')
        return rv

    def files_unquoted(self):
        # compared to the standard `unquoted`, we do not match
        # '=' as part of the unquoted string. Why? Varjson pair 
        # is separated by `=` symbol
        # TODO: Avoid duplication by taking in arguments to standard
        # `unquoted` function
        acceptable_chars = '0-9A-Za-z \t!$%&()*+./;<=>?^_`|~-'

        chars = [self.char(acceptable_chars)]

        while True:
            char = self.maybe_char(acceptable_chars)
            if char is None:
                break

            if char in 'Ee.':
                number_type = float

            chars.append(char)

        rv = ''.join(chars).rstrip(' \t')
        return rv

    def unquoted(self):
        acceptable_chars = '0-9A-Za-z \t!$%&()*+./;<=>?^_`|~-'

        chars = [self.char(acceptable_chars)]

        while True:
            char = self.maybe_char(acceptable_chars)
            if char is None:
                break

            if char in 'Ee.':
                number_type = float

            chars.append(char)

        rv = ''.join(chars).rstrip(' \t')
        return rv

    def quoted_string(self):
        quote = self.char("\"'")
        chars = []

        escape_sequences = {
            "b": "\b",
            "f": "\f",
            "n": "\n",
            "r": "\r",
            "t": "\t",
            "\\": "\\",
            "/": "/",
            quote: f"\{quote}",
        }

        while True:
            char = self.char()
            if char == quote:
                break
            elif char == "\\":
                escape = self.char()
                if escape == "u":
                    code_point = []
                    for i in range(4):
                        code_point.append(self.char("0-9a-fA-F"))

                    chars.append(chr(int("".join(code_point), 16)))
                else:
                    try:
                        chars.append(escape_sequences[escape])
                    except:
                        raise ParseError(
                            self.pos + 1,
                            "Invalid escape sequence",
                            self.text[self.pos + 1],
                        )

            else:
                if unicodedata.category(char) == "Cc":
                    raise ParseError(
                        self.pos + 1,
                        "Unescaped control character",
                        self.text[self.pos + 1],
                    )

                chars.append(char)

        try:
            res = "".join(chars)
        except Exception as e:
            raise e
        return res


if __name__ == "__main__":
    import sys
    from pprint import pprint

    parser = ElfParser()

    try:
        pprint(parser.parse(sys.stdin.read()))
    except ParseError as e:
        print("Error: " + str(e))
