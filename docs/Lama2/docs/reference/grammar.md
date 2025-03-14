# Grammar

What follows is a rough rendition of the *Lama2* grammar, followed by
a visual exploration of the grammar in the railroad diagram format (thanks to
[Railroad Diagram Generator](https://bottlecaps.de/rr/ui))

## EBNF Description

```
Lama2File       ::= (Proceessor Separator)? Requestor (Separator Processor Separator Requestor)*
Separator       ::= `^---$`
Processor       ::= `(?!(get|post|head|put|delete|connect|trace|patch))`
Requestor       ::= HTTPVerb Multipart? TheURL Details?
HTTPVerb        ::= "get" | "head" | "post" | "put" | "delete" | "connect" | "trace" | "patch"
Multipart       ::= "multipart"
TheURL          ::= "http" "s"? "://" [A-Za-z0-9-._~:/?#[@!$&'()*+,;%=]+ /* ws: explicit */
Details         ::= HeaderData | DataHeader

HeaderData      ::= Headers DataInput?
DataHeader      ::= DataInput Headers?
Headers         ::= HeaderPair HeaderPair*
HeaderPair      ::= (QuotedString | Unquoted) ":" (QuotedString | Unquoted)
DataInput       ::= VarJSON | JSONType

VarJSON         ::= VarJSONPair VarJSONPair* FilesPair?
VarJSONPair     ::= (QuotedString | VarJSONUnquoted) "=" (QuotedString | VarJSONUnquoted)
FilesPair       ::= FilesPair FilesPair*
FilesPair       ::= (QuotedString | FilesUnquoted) "@" (QuotedString | FilesUnquoted)
VarJSONUnquoted ::= [@0-9A-Za-z \t!$%&()*+./;<>?^_`|~-]+ /* ws: explicit */
FilesUnquoted   ::= [0-9A-Za-z \t!$%&()*+./;<>?^_`|~-]+ /* ws: explicit */
QuotedString    ::= ['"] Char* ['"]
Unquoted        ::= [0-9A-Za-z \t!$%&()*+./;<=>?^_`|~-]+ /* ws: explicit */


JSONType        ::= ComplexType | PrimitiveType
ComplexType     ::= List | Map
PrimitiveType   ::= Null | Boolean | QuotedString | Number
Map             ::= "{" Pair? (Pair ",")* "}"
List            ::= "[" JSONType? (JSONType ",")* "]"
Pair            ::= QuotedString ":" JSONType

Boolean         ::= "true" | "false"
Null            ::= "null"


Number          ::= Integer Fraction? Exponent?
Exponent        ::= [eE] Sign? Digits
Fraction        ::= "." Digits
Integer         ::= IntegerRule4 | IntegerRule3 | IntegerRule2 | IntegerRule1

IntegerRule1    ::= Digit
IntegerRule2    ::= OneNine Digits
IntegerRule3    ::= Sign IntegerRule1
IntegerRule4    ::= Sign IntegerRule2

Digits          ::= Digit Digit*
Digit           ::= "0"? OneNine
OneNine         ::= [1-9]
Sign            ::= [-+]
```

## Railroad Diagram

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml">
   <head>
      <meta http-equiv="Content-Type" content="application/xhtml+xml; charset=UTF-8" />
      <meta name="generator" content="Railroad Diagram Generator 1.67" />
      <style type="text/css">
    ::-moz-selection
    {
      color: #FFFCF0;
      background: #0F0C00;
    }
    ::selection
    {
      color: #FFFCF0;
      background: #0F0C00;
    }
    .ebnf a, .grammar a
    {
      text-decoration: none;
    }
    .ebnf a:hover, .grammar a:hover
    {
      color: #050400;
      text-decoration: underline;
    }
    .signature
    {
      color: #806600;
      font-size: 11px;
      text-align: right;
    }
    body
    {
      font: normal 12px Verdana, sans-serif;
      color: #0F0C00;
      background: #FFFCF0;
    }
    a:link, a:visited
    {
      color: #0F0C00;
    }
    a:link.signature, a:visited.signature
    {
      color: #806600;
    }
    a.button, #tabs li a
    {
      padding: 0.25em 0.5em;
      border: 1px solid #806600;
      background: #F1E8C6;
      color: #806600;
      text-decoration: none;
      font-weight: bold;
    }
    a.button:hover, #tabs li a:hover
    {
      color: #050400;
      background: #FFF6D1;
      border-color: #050400;
    }
    #tabs
    {
      padding: 3px 10px;
      margin-left: 0;
      margin-top: 58px;
      border-bottom: 1px solid #0F0C00;
    }
    #tabs li
    {
      list-style: none;
      margin-left: 5px;
      display: inline;
    }
    #tabs li a
    {
      border-bottom: 1px solid #0F0C00;
    }
    #tabs li a.active
    {
      color: #0F0C00;
      background: #FFFCF0;
      border-color: #0F0C00;
      border-bottom: 1px solid #FFFCF0;
      outline: none;
    }
    #divs div
    {
      display: none;
      overflow:auto;
    }
    #divs div.active
    {
      display: block;
    }
    #text
    {
      border-color: #806600;
      background: #FFFEFA;
      color: #050400;
    }
    .small
    {
      vertical-align: top;
      text-align: right;
      font-size: 9px;
      font-weight: normal;
      line-height: 120%;
    }
    td.small
    {
      padding-top: 0px;
    }
    .hidden
    {
      visibility: hidden;
    }
    td:hover .hidden
    {
      visibility: visible;
    }
    div.download
    {
      display: none;
      background: #FFFCF0;
      position: absolute;
      right: 34px;
      top: 94px;
      padding: 10px;
      border: 1px dotted #0F0C00;
    }
    #divs div.ebnf, .ebnf code
    {
      display: block;
      padding: 10px;
      background: #FFF6D1;
      width: 992px;
    }
    #divs div.grammar
    {
      display: block;
      padding-left: 16px;
      padding-top: 2px;
      padding-bottom: 2px;
      background: #FFF6D1;
    }
    pre
    {
      margin: 0px;
    }
    .ebnf div
    {
      padding-left: 13ch;
      text-indent: -13ch;
    }
    .ebnf code, .grammar code, textarea, pre
    {
      font:12px SFMono-Regular,Consolas,Liberation Mono,Menlo,Courier,monospace;
    }
    tr.option-line td:first-child
    {
      text-align: right
    }
    tr.option-text td
    {
      padding-bottom: 10px
    }
    table.palette
    {
      border-top: 1px solid #050400;
      border-right: 1px solid #050400;
      margin-bottom: 4px
    }
    td.palette
    {
      border-bottom: 1px solid #050400;
      border-left: 1px solid #050400;
    }
    a.palette
    {
      padding: 2px 3px 2px 10px;
      text-decoration: none;
    }
    .palette
    {
      -webkit-user-select: none;
      -khtml-user-select: none;
      -moz-user-select: none;
      -o-user-select: none;
      -ms-user-select: none;
    }
  </style><svg xmlns="http://www.w3.org/2000/svg">
         <defs>
            <style type="text/css">
    @namespace "http://www.w3.org/2000/svg";
    .line                 {fill: none; stroke: #332900; stroke-width: 1;}
    .bold-line            {stroke: #141000; shape-rendering: crispEdges; stroke-width: 2;}
    .thin-line            {stroke: #1F1800; shape-rendering: crispEdges}
    .filled               {fill: #332900; stroke: none;}
    text.terminal         {font-family: Verdana, Sans-serif;
                            font-size: 12px;
                            fill: #141000;
                            font-weight: bold;
                          }
    text.nonterminal      {font-family: Verdana, Sans-serif;
                            font-size: 12px;
                            fill: #1A1400;
                            font-weight: normal;
                          }
    text.regexp           {font-family: Verdana, Sans-serif;
                            font-size: 12px;
                            fill: #1F1800;
                            font-weight: normal;
                          }
    rect, circle, polygon {fill: #332900; stroke: #332900;}
    rect.terminal         {fill: #FFDB4D; stroke: #332900; stroke-width: 1;}
    rect.nonterminal      {fill: #FFEC9E; stroke: #332900; stroke-width: 1;}
    rect.text             {fill: none; stroke: none;}
    polygon.regexp        {fill: #FFF4C7; stroke: #332900; stroke-width: 1;}
  </style>
         </defs></svg></head>
   <body>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Lama2File">Lama2File:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="637" height="113">
         <polygon points="9 61 1 57 1 65"/>
         <polygon points="17 61 9 57 9 65"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Proceessor" xlink:title="Proceessor">
            <rect x="51" y="79" width="90" height="32"/>
            <rect x="49" y="77" width="90" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="97">Proceessor</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Separator" xlink:title="Separator">
            <rect x="161" y="79" width="82" height="32"/>
            <rect x="159" y="77" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="169" y="97">Separator</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Requestor" xlink:title="Requestor">
            <rect x="303" y="47" width="84" height="32"/>
            <rect x="301" y="45" width="84" height="32" class="nonterminal"/>
            <text class="nonterminal" x="311" y="65">Requestor</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Separator" xlink:title="Separator">
            <rect x="303" y="3" width="82" height="32"/>
            <rect x="301" y="1" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="311" y="21">Separator</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Processor" xlink:title="Processor">
            <rect x="405" y="3" width="82" height="32"/>
            <rect x="403" y="1" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="413" y="21">Processor</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Separator" xlink:title="Separator">
            <rect x="507" y="3" width="82" height="32"/>
            <rect x="505" y="1" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="515" y="21">Separator</text></a><path class="line" d="m17 61 h2 m20 0 h10 m0 0 h202 m-232 0 h20 m212 0 h20 m-252 0 q10 0 10 10 m232 0 q0 -10 10 -10 m-242 10 v12 m232 0 v-12 m-232 12 q0 10 10 10 m212 0 q10 0 10 -10 m-222 10 h10 m90 0 h10 m0 0 h10 m82 0 h10 m40 -32 h10 m84 0 h10 m0 0 h202 m-326 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -24 q0 -10 10 -10 m306 44 l20 0 m-20 0 q10 0 10 -10 l0 -24 q0 -10 -10 -10 m-306 0 h10 m82 0 h10 m0 0 h10 m82 0 h10 m0 0 h10 m82 0 h10 m23 44 h-3"/>
         <polygon points="627 61 635 57 635 65"/>
         <polygon points="627 61 619 57 619 65"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Lama2File" title="Lama2File">Lama2File</a></div>
               <div>         ::= ( <a href="#Proceessor" title="Proceessor">Proceessor</a> <a href="#Separator" title="Separator">Separator</a> )? <a href="#Requestor" title="Requestor">Requestor</a> ( <a href="#Separator" title="Separator">Separator</a> <a href="#Processor" title="Processor">Processor</a> <a href="#Separator" title="Separator">Separator</a> <a href="#Requestor" title="Requestor">Requestor</a> )*</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">no references</p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Separator">Separator:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="97" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="31" y="3" width="38" height="32" rx="10"/>
         <rect x="29" y="1" width="38" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="21">---</text>
         <path class="line" d="m17 17 h2 m0 0 h10 m38 0 h10 m3 0 h-3"/>
         <polygon points="87 17 95 13 95 21"/>
         <polygon points="87 17 79 13 79 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Separator" title="Separator">Separator</a></div>
               <div>         ::= '---'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Lama2File" title="Lama2File">Lama2File</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Processor">Processor:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="437" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="31" y="3" width="378" height="32" rx="10"/>
         <rect x="29" y="1" width="378" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="21">(?!(get|post|head|put|delete|connect|trace|patch))</text>
         <path class="line" d="m17 17 h2 m0 0 h10 m378 0 h10 m3 0 h-3"/>
         <polygon points="427 17 435 13 435 21"/>
         <polygon points="427 17 419 13 419 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Processor" title="Processor">Processor</a></div>
               <div>         ::= '(?!(get|post|head|put|delete|connect|trace|patch))'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Lama2File" title="Lama2File">Lama2File</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Requestor">Requestor:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="485" height="69">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#HTTPVerb" xlink:title="HTTPVerb">
            <rect x="31" y="3" width="80" height="32"/>
            <rect x="29" y="1" width="80" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">HTTPVerb</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Multipart" xlink:title="Multipart">
            <rect x="151" y="35" width="76" height="32"/>
            <rect x="149" y="33" width="76" height="32" class="nonterminal"/>
            <text class="nonterminal" x="159" y="53">Multipart</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#TheURL" xlink:title="TheURL">
            <rect x="267" y="3" width="66" height="32"/>
            <rect x="265" y="1" width="66" height="32" class="nonterminal"/>
            <text class="nonterminal" x="275" y="21">TheURL</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Details" xlink:title="Details">
            <rect x="373" y="35" width="64" height="32"/>
            <rect x="371" y="33" width="64" height="32" class="nonterminal"/>
            <text class="nonterminal" x="381" y="53">Details</text></a><path class="line" d="m17 17 h2 m0 0 h10 m80 0 h10 m20 0 h10 m0 0 h86 m-116 0 h20 m96 0 h20 m-136 0 q10 0 10 10 m116 0 q0 -10 10 -10 m-126 10 v12 m116 0 v-12 m-116 12 q0 10 10 10 m96 0 q10 0 10 -10 m-106 10 h10 m76 0 h10 m20 -32 h10 m66 0 h10 m20 0 h10 m0 0 h74 m-104 0 h20 m84 0 h20 m-124 0 q10 0 10 10 m104 0 q0 -10 10 -10 m-114 10 v12 m104 0 v-12 m-104 12 q0 10 10 10 m84 0 q10 0 10 -10 m-94 10 h10 m64 0 h10 m23 -32 h-3"/>
         <polygon points="475 17 483 13 483 21"/>
         <polygon points="475 17 467 13 467 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Requestor" title="Requestor">Requestor</a></div>
               <div>         ::= <a href="#HTTPVerb" title="HTTPVerb">HTTPVerb</a> <a href="#Multipart" title="Multipart">Multipart</a>? <a href="#TheURL" title="TheURL">TheURL</a> <a href="#Details" title="Details">Details</a>?</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Lama2File" title="Lama2File">Lama2File</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="HTTPVerb">HTTPVerb:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="171" height="345">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="51" y="3" width="42" height="32" rx="10"/>
         <rect x="49" y="1" width="42" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="21">get</text>
         <rect x="51" y="47" width="52" height="32" rx="10"/>
         <rect x="49" y="45" width="52" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="65">head</text>
         <rect x="51" y="91" width="48" height="32" rx="10"/>
         <rect x="49" y="89" width="48" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="109">post</text>
         <rect x="51" y="135" width="42" height="32" rx="10"/>
         <rect x="49" y="133" width="42" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="153">put</text>
         <rect x="51" y="179" width="62" height="32" rx="10"/>
         <rect x="49" y="177" width="62" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="197">delete</text>
         <rect x="51" y="223" width="72" height="32" rx="10"/>
         <rect x="49" y="221" width="72" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="241">connect</text>
         <rect x="51" y="267" width="54" height="32" rx="10"/>
         <rect x="49" y="265" width="54" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="285">trace</text>
         <rect x="51" y="311" width="56" height="32" rx="10"/>
         <rect x="49" y="309" width="56" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="329">patch</text>
         <path class="line" d="m17 17 h2 m20 0 h10 m42 0 h10 m0 0 h30 m-112 0 h20 m92 0 h20 m-132 0 q10 0 10 10 m112 0 q0 -10 10 -10 m-122 10 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m52 0 h10 m0 0 h20 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m48 0 h10 m0 0 h24 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m42 0 h10 m0 0 h30 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m62 0 h10 m0 0 h10 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m72 0 h10 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m54 0 h10 m0 0 h18 m-102 -10 v20 m112 0 v-20 m-112 20 v24 m112 0 v-24 m-112 24 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m56 0 h10 m0 0 h16 m23 -308 h-3"/>
         <polygon points="161 17 169 13 169 21"/>
         <polygon points="161 17 153 13 153 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#HTTPVerb" title="HTTPVerb">HTTPVerb</a> ::= 'get'</div>
               <div>           | 'head'</div>
               <div>           | 'post'</div>
               <div>           | 'put'</div>
               <div>           | 'delete'</div>
               <div>           | 'connect'</div>
               <div>           | 'trace'</div>
               <div>           | 'patch'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Requestor" title="Requestor">Requestor</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Multipart">Multipart:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="141" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="31" y="3" width="82" height="32" rx="10"/>
         <rect x="29" y="1" width="82" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="21">multipart</text>
         <path class="line" d="m17 17 h2 m0 0 h10 m82 0 h10 m3 0 h-3"/>
         <polygon points="131 17 139 13 139 21"/>
         <polygon points="131 17 123 13 123 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Multipart" title="Multipart">Multipart</a></div>
               <div>         ::= 'multipart'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Requestor" title="Requestor">Requestor</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="TheURL">TheURL:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="417" height="1109">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <rect x="31" y="19" width="48" height="32" rx="10"/>
         <rect x="29" y="17" width="48" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="37">http</text>
         <rect x="119" y="51" width="26" height="32" rx="10"/>
         <rect x="117" y="49" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="127" y="69">s</text>
         <rect x="185" y="19" width="42" height="32" rx="10"/>
         <rect x="183" y="17" width="42" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="193" y="37">://</text>
         <polygon points="287 35 294 19 342 19 349 35 342 51 294 51"/>
         <polygon points="285 33 292 17 340 17 347 33 340 49 292 49" class="regexp"/>
         <text class="regexp" x="300" y="37">[A-Z]</text>
         <polygon points="287 79 294 63 340 63 347 79 340 95 294 95"/>
         <polygon points="285 77 292 61 338 61 345 77 338 93 292 93" class="regexp"/>
         <text class="regexp" x="300" y="81">[a-z]</text>
         <polygon points="287 123 294 107 342 107 349 123 342 139 294 139"/>
         <polygon points="285 121 292 105 340 105 347 121 340 137 292 137" class="regexp"/>
         <text class="regexp" x="300" y="125">[0-9]</text>
         <rect x="287" y="151" width="26" height="32" rx="10"/>
         <rect x="285" y="149" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="169">-</text>
         <rect x="287" y="195" width="24" height="32" rx="10"/>
         <rect x="285" y="193" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="213">.</text>
         <rect x="287" y="239" width="28" height="32" rx="10"/>
         <rect x="285" y="237" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="257">_</text>
         <rect x="287" y="283" width="30" height="32" rx="10"/>
         <rect x="285" y="281" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="301">~</text>
         <rect x="287" y="327" width="24" height="32" rx="10"/>
         <rect x="285" y="325" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="345">:</text>
         <rect x="287" y="371" width="28" height="32" rx="10"/>
         <rect x="285" y="369" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="389">/</text>
         <rect x="287" y="415" width="26" height="32" rx="10"/>
         <rect x="285" y="413" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="433">?</text>
         <rect x="287" y="459" width="30" height="32" rx="10"/>
         <rect x="285" y="457" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="477">#</text>
         <rect x="287" y="503" width="26" height="32" rx="10"/>
         <rect x="285" y="501" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="521">[</text>
         <rect x="287" y="547" width="32" height="32" rx="10"/>
         <rect x="285" y="545" width="32" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="565">@</text>
         <rect x="287" y="591" width="24" height="32" rx="10"/>
         <rect x="285" y="589" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="609">!</text>
         <rect x="287" y="635" width="28" height="32" rx="10"/>
         <rect x="285" y="633" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="653">$</text>
         <rect x="287" y="679" width="30" height="32" rx="10"/>
         <rect x="285" y="677" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="697">&amp;</text>
         <rect x="287" y="723" width="24" height="32" rx="10"/>
         <rect x="285" y="721" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="741">'</text>
         <rect x="287" y="767" width="26" height="32" rx="10"/>
         <rect x="285" y="765" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="785">(</text>
         <rect x="287" y="811" width="26" height="32" rx="10"/>
         <rect x="285" y="809" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="829">)</text>
         <rect x="287" y="855" width="28" height="32" rx="10"/>
         <rect x="285" y="853" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="873">*</text>
         <rect x="287" y="899" width="30" height="32" rx="10"/>
         <rect x="285" y="897" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="917">+</text>
         <rect x="287" y="943" width="24" height="32" rx="10"/>
         <rect x="285" y="941" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="961">,</text>
         <rect x="287" y="987" width="24" height="32" rx="10"/>
         <rect x="285" y="985" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="1005">;</text>
         <rect x="287" y="1031" width="34" height="32" rx="10"/>
         <rect x="285" y="1029" width="34" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="1049">%</text>
         <rect x="287" y="1075" width="30" height="32" rx="10"/>
         <rect x="285" y="1073" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="1093">=</text>
         <path class="line" d="m17 33 h2 m0 0 h10 m48 0 h10 m20 0 h10 m0 0 h36 m-66 0 h20 m46 0 h20 m-86 0 q10 0 10 10 m66 0 q0 -10 10 -10 m-76 10 v12 m66 0 v-12 m-66 12 q0 10 10 10 m46 0 q10 0 10 -10 m-56 10 h10 m26 0 h10 m20 -32 h10 m42 0 h10 m40 0 h10 m62 0 h10 m-102 0 h20 m82 0 h20 m-122 0 q10 0 10 10 m102 0 q0 -10 10 -10 m-112 10 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m60 0 h10 m0 0 h2 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m62 0 h10 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m32 0 h10 m0 0 h30 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m34 0 h10 m0 0 h28 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-122 -1056 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m122 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-122 0 h10 m0 0 h112 m23 32 h-3"/>
         <polygon points="407 33 415 29 415 37"/>
         <polygon points="407 33 399 29 399 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#TheURL" title="TheURL">TheURL</a>   ::= 'http' 's'? '://' [A-Za-z0-9._~:/?[@!$&amp;'()*+,;%=#x2D#x23]+</div>
               <div>          /* ws: explicit */</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Requestor" title="Requestor">Requestor</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Details">Details:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="195" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#HeaderData" xlink:title="HeaderData">
            <rect x="51" y="3" width="96" height="32"/>
            <rect x="49" y="1" width="96" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">HeaderData</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#DataHeader" xlink:title="DataHeader">
            <rect x="51" y="47" width="96" height="32"/>
            <rect x="49" y="45" width="96" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">DataHeader</text></a><path class="line" d="m17 17 h2 m20 0 h10 m96 0 h10 m-136 0 h20 m116 0 h20 m-156 0 q10 0 10 10 m136 0 q0 -10 10 -10 m-146 10 v24 m136 0 v-24 m-136 24 q0 10 10 10 m116 0 q10 0 10 -10 m-126 10 h10 m96 0 h10 m23 -44 h-3"/>
         <polygon points="185 17 193 13 193 21"/>
         <polygon points="185 17 177 13 177 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Details" title="Details">Details</a>  ::= <a href="#HeaderData" title="HeaderData">HeaderData</a></div>
               <div>           | <a href="#DataHeader" title="DataHeader">DataHeader</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Requestor" title="Requestor">Requestor</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="HeaderData">HeaderData:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="275" height="69">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Headers" xlink:title="Headers">
            <rect x="31" y="3" width="72" height="32"/>
            <rect x="29" y="1" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">Headers</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#DataInput" xlink:title="DataInput">
            <rect x="143" y="35" width="84" height="32"/>
            <rect x="141" y="33" width="84" height="32" class="nonterminal"/>
            <text class="nonterminal" x="151" y="53">DataInput</text></a><path class="line" d="m17 17 h2 m0 0 h10 m72 0 h10 m20 0 h10 m0 0 h94 m-124 0 h20 m104 0 h20 m-144 0 q10 0 10 10 m124 0 q0 -10 10 -10 m-134 10 v12 m124 0 v-12 m-124 12 q0 10 10 10 m104 0 q10 0 10 -10 m-114 10 h10 m84 0 h10 m23 -32 h-3"/>
         <polygon points="265 17 273 13 273 21"/>
         <polygon points="265 17 257 13 257 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#HeaderData" title="HeaderData">HeaderData</a></div>
               <div>         ::= <a href="#Headers" title="Headers">Headers</a> <a href="#DataInput" title="DataInput">DataInput</a>?</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Details" title="Details">Details</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="DataHeader">DataHeader:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="275" height="69">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#DataInput" xlink:title="DataInput">
            <rect x="31" y="3" width="84" height="32"/>
            <rect x="29" y="1" width="84" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">DataInput</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Headers" xlink:title="Headers">
            <rect x="155" y="35" width="72" height="32"/>
            <rect x="153" y="33" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="163" y="53">Headers</text></a><path class="line" d="m17 17 h2 m0 0 h10 m84 0 h10 m20 0 h10 m0 0 h82 m-112 0 h20 m92 0 h20 m-132 0 q10 0 10 10 m112 0 q0 -10 10 -10 m-122 10 v12 m112 0 v-12 m-112 12 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m72 0 h10 m23 -32 h-3"/>
         <polygon points="265 17 273 13 273 21"/>
         <polygon points="265 17 257 13 257 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#DataHeader" title="DataHeader">DataHeader</a></div>
               <div>         ::= <a href="#DataInput" title="DataInput">DataInput</a> <a href="#Headers" title="Headers">Headers</a>?</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Details" title="Details">Details</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Headers">Headers:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="189" height="53">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#HeaderPair" xlink:title="HeaderPair">
            <rect x="51" y="19" width="90" height="32"/>
            <rect x="49" y="17" width="90" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="37">HeaderPair</text></a><path class="line" d="m17 33 h2 m20 0 h10 m90 0 h10 m-130 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m110 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-110 0 h10 m0 0 h100 m23 32 h-3"/>
         <polygon points="179 33 187 29 187 37"/>
         <polygon points="179 33 171 29 171 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Headers" title="Headers">Headers</a>  ::= <a href="#HeaderPair" title="HeaderPair">HeaderPair</a> <a href="#HeaderPair" title="HeaderPair">HeaderPair</a>*</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#DataHeader" title="DataHeader">DataHeader</a></li>
            <li><a href="#HeaderData" title="HeaderData">HeaderData</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="HeaderPair">HeaderPair:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="411" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="51" y="3" width="104" height="32"/>
            <rect x="49" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Unquoted" xlink:title="Unquoted">
            <rect x="51" y="47" width="82" height="32"/>
            <rect x="49" y="45" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">Unquoted</text></a><rect x="195" y="3" width="24" height="32" rx="10"/>
         <rect x="193" y="1" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="203" y="21">:</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="259" y="3" width="104" height="32"/>
            <rect x="257" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="267" y="21">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Unquoted" xlink:title="Unquoted">
            <rect x="259" y="47" width="82" height="32"/>
            <rect x="257" y="45" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="267" y="65">Unquoted</text></a><path class="line" d="m17 17 h2 m20 0 h10 m104 0 h10 m-144 0 h20 m124 0 h20 m-164 0 q10 0 10 10 m144 0 q0 -10 10 -10 m-154 10 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m82 0 h10 m0 0 h22 m20 -44 h10 m24 0 h10 m20 0 h10 m104 0 h10 m-144 0 h20 m124 0 h20 m-164 0 q10 0 10 10 m144 0 q0 -10 10 -10 m-154 10 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m82 0 h10 m0 0 h22 m23 -44 h-3"/>
         <polygon points="401 17 409 13 409 21"/>
         <polygon points="401 17 393 13 393 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#HeaderPair" title="HeaderPair">HeaderPair</a></div>
               <div>         ::= ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#Unquoted" title="Unquoted">Unquoted</a> ) ':' ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#Unquoted" title="Unquoted">Unquoted</a> )</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Headers" title="Headers">Headers</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="DataInput">DataInput:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="181" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#VarJSON" xlink:title="VarJSON">
            <rect x="51" y="3" width="72" height="32"/>
            <rect x="49" y="1" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">VarJSON</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#JSONType" xlink:title="JSONType">
            <rect x="51" y="47" width="82" height="32"/>
            <rect x="49" y="45" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">JSONType</text></a><path class="line" d="m17 17 h2 m20 0 h10 m72 0 h10 m0 0 h10 m-122 0 h20 m102 0 h20 m-142 0 q10 0 10 10 m122 0 q0 -10 10 -10 m-132 10 v24 m122 0 v-24 m-122 24 q0 10 10 10 m102 0 q10 0 10 -10 m-112 10 h10 m82 0 h10 m23 -44 h-3"/>
         <polygon points="171 17 179 13 179 21"/>
         <polygon points="171 17 163 13 163 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#DataInput" title="DataInput">DataInput</a></div>
               <div>         ::= <a href="#VarJSON" title="VarJSON">VarJSON</a></div>
               <div>           | <a href="#JSONType" title="JSONType">JSONType</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#DataHeader" title="DataHeader">DataHeader</a></li>
            <li><a href="#HeaderData" title="HeaderData">HeaderData</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="VarJSON">VarJSON:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="329" height="85">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#VarJSONPair" xlink:title="VarJSONPair">
            <rect x="51" y="19" width="98" height="32"/>
            <rect x="49" y="17" width="98" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="37">VarJSONPair</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#FilesPair" xlink:title="FilesPair">
            <rect x="209" y="51" width="72" height="32"/>
            <rect x="207" y="49" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="217" y="69">FilesPair</text></a><path class="line" d="m17 33 h2 m20 0 h10 m98 0 h10 m-138 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m118 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-118 0 h10 m0 0 h108 m40 32 h10 m0 0 h82 m-112 0 h20 m92 0 h20 m-132 0 q10 0 10 10 m112 0 q0 -10 10 -10 m-122 10 v12 m112 0 v-12 m-112 12 q0 10 10 10 m92 0 q10 0 10 -10 m-102 10 h10 m72 0 h10 m23 -32 h-3"/>
         <polygon points="319 33 327 29 327 37"/>
         <polygon points="319 33 311 29 311 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#VarJSON" title="VarJSON">VarJSON</a>  ::= <a href="#VarJSONPair" title="VarJSONPair">VarJSONPair</a> <a href="#VarJSONPair" title="VarJSONPair">VarJSONPair</a>* <a href="#FilesPair" title="FilesPair">FilesPair</a>?</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#DataInput" title="DataInput">DataInput</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="VarJSONPair">VarJSONPair:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="477" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="51" y="3" width="104" height="32"/>
            <rect x="49" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#VarJSONUnquoted" xlink:title="VarJSONUnquoted">
            <rect x="51" y="47" width="134" height="32"/>
            <rect x="49" y="45" width="134" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">VarJSONUnquoted</text></a><rect x="225" y="3" width="30" height="32" rx="10"/>
         <rect x="223" y="1" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="233" y="21">=</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="295" y="3" width="104" height="32"/>
            <rect x="293" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="303" y="21">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#VarJSONUnquoted" xlink:title="VarJSONUnquoted">
            <rect x="295" y="47" width="134" height="32"/>
            <rect x="293" y="45" width="134" height="32" class="nonterminal"/>
            <text class="nonterminal" x="303" y="65">VarJSONUnquoted</text></a><path class="line" d="m17 17 h2 m20 0 h10 m104 0 h10 m0 0 h30 m-174 0 h20 m154 0 h20 m-194 0 q10 0 10 10 m174 0 q0 -10 10 -10 m-184 10 v24 m174 0 v-24 m-174 24 q0 10 10 10 m154 0 q10 0 10 -10 m-164 10 h10 m134 0 h10 m20 -44 h10 m30 0 h10 m20 0 h10 m104 0 h10 m0 0 h30 m-174 0 h20 m154 0 h20 m-194 0 q10 0 10 10 m174 0 q0 -10 10 -10 m-184 10 v24 m174 0 v-24 m-174 24 q0 10 10 10 m154 0 q10 0 10 -10 m-164 10 h10 m134 0 h10 m23 -44 h-3"/>
         <polygon points="467 17 475 13 475 21"/>
         <polygon points="467 17 459 13 459 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#VarJSONPair" title="VarJSONPair">VarJSONPair</a></div>
               <div>         ::= ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#VarJSONUnquoted" title="VarJSONUnquoted">VarJSONUnquoted</a> ) '=' ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#VarJSONUnquoted" title="VarJSONUnquoted">VarJSONUnquoted</a> )</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#VarJSON" title="VarJSON">VarJSON</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="FilesPair">FilesPair:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="471" height="141">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#FilesPair" xlink:title="FilesPair">
            <rect x="71" y="19" width="72" height="32"/>
            <rect x="69" y="17" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="79" y="37">FilesPair</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="71" y="63" width="104" height="32"/>
            <rect x="69" y="61" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="79" y="81">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#FilesUnquoted" xlink:title="FilesUnquoted">
            <rect x="71" y="107" width="110" height="32"/>
            <rect x="69" y="105" width="110" height="32" class="nonterminal"/>
            <text class="nonterminal" x="79" y="125">FilesUnquoted</text></a><rect x="221" y="63" width="32" height="32" rx="10"/>
         <rect x="219" y="61" width="32" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="229" y="81">@</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="293" y="63" width="104" height="32"/>
            <rect x="291" y="61" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="301" y="81">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#FilesUnquoted" xlink:title="FilesUnquoted">
            <rect x="293" y="107" width="110" height="32"/>
            <rect x="291" y="105" width="110" height="32" class="nonterminal"/>
            <text class="nonterminal" x="301" y="125">FilesUnquoted</text></a><path class="line" d="m17 33 h2 m40 0 h10 m72 0 h10 m-112 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m92 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-92 0 h10 m0 0 h82 m20 32 h260 m-412 0 h20 m392 0 h20 m-432 0 q10 0 10 10 m412 0 q0 -10 10 -10 m-422 10 v24 m412 0 v-24 m-412 24 q0 10 10 10 m392 0 q10 0 10 -10 m-382 10 h10 m104 0 h10 m0 0 h6 m-150 0 h20 m130 0 h20 m-170 0 q10 0 10 10 m150 0 q0 -10 10 -10 m-160 10 v24 m150 0 v-24 m-150 24 q0 10 10 10 m130 0 q10 0 10 -10 m-140 10 h10 m110 0 h10 m20 -44 h10 m32 0 h10 m20 0 h10 m104 0 h10 m0 0 h6 m-150 0 h20 m130 0 h20 m-170 0 q10 0 10 10 m150 0 q0 -10 10 -10 m-160 10 v24 m150 0 v-24 m-150 24 q0 10 10 10 m130 0 q10 0 10 -10 m-140 10 h10 m110 0 h10 m43 -88 h-3"/>
         <polygon points="461 33 469 29 469 37"/>
         <polygon points="461 33 453 29 453 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#FilesPair" title="FilesPair">FilesPair</a></div>
               <div>         ::= <a href="#FilesPair" title="FilesPair">FilesPair</a> <a href="#FilesPair" title="FilesPair">FilesPair</a>*</div>
               <div>           | ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#FilesUnquoted" title="FilesUnquoted">FilesUnquoted</a> ) '@' ( <a href="#QuotedString" title="QuotedString">QuotedString</a> | <a href="#FilesUnquoted" title="FilesUnquoted">FilesUnquoted</a> )</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#FilesPair" title="FilesPair">FilesPair</a></li>
            <li><a href="#VarJSON" title="VarJSON">VarJSON</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="VarJSONUnquoted">VarJSONUnquoted:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="201" height="1197">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <rect x="71" y="19" width="32" height="32" rx="10"/>
         <rect x="69" y="17" width="32" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="37">@</text>
         <polygon points="71 79 78 63 126 63 133 79 126 95 78 95"/>
         <polygon points="69 77 76 61 124 61 131 77 124 93 76 93" class="regexp"/>
         <text class="regexp" x="84" y="81">[0-9]</text>
         <polygon points="71 123 78 107 126 107 133 123 126 139 78 139"/>
         <polygon points="69 121 76 105 124 105 131 121 124 137 76 137" class="regexp"/>
         <text class="regexp" x="84" y="125">[A-Z]</text>
         <polygon points="71 167 78 151 124 151 131 167 124 183 78 183"/>
         <polygon points="69 165 76 149 122 149 129 165 122 181 76 181" class="regexp"/>
         <text class="regexp" x="84" y="169">[a-z]</text>
         <rect x="71" y="195" width="24" height="32" rx="10"/>
         <rect x="69" y="193" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="213"/>
         <rect x="71" y="239" width="28" height="32" rx="10"/>
         <rect x="69" y="237" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="257">\</text>
         <rect x="71" y="283" width="24" height="32" rx="10"/>
         <rect x="69" y="281" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="301">t</text>
         <rect x="71" y="327" width="24" height="32" rx="10"/>
         <rect x="69" y="325" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="345">!</text>
         <rect x="71" y="371" width="28" height="32" rx="10"/>
         <rect x="69" y="369" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="389">$</text>
         <rect x="71" y="415" width="34" height="32" rx="10"/>
         <rect x="69" y="413" width="34" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="433">%</text>
         <rect x="71" y="459" width="30" height="32" rx="10"/>
         <rect x="69" y="457" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="477">&amp;</text>
         <rect x="71" y="503" width="26" height="32" rx="10"/>
         <rect x="69" y="501" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="521">(</text>
         <rect x="71" y="547" width="26" height="32" rx="10"/>
         <rect x="69" y="545" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="565">)</text>
         <rect x="71" y="591" width="28" height="32" rx="10"/>
         <rect x="69" y="589" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="609">*</text>
         <rect x="71" y="635" width="30" height="32" rx="10"/>
         <rect x="69" y="633" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="653">+</text>
         <rect x="71" y="679" width="24" height="32" rx="10"/>
         <rect x="69" y="677" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="697">.</text>
         <rect x="71" y="723" width="28" height="32" rx="10"/>
         <rect x="69" y="721" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="741">/</text>
         <rect x="71" y="767" width="24" height="32" rx="10"/>
         <rect x="69" y="765" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="785">;</text>
         <rect x="71" y="811" width="30" height="32" rx="10"/>
         <rect x="69" y="809" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="829">&lt;</text>
         <rect x="71" y="855" width="30" height="32" rx="10"/>
         <rect x="69" y="853" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="873">&gt;</text>
         <rect x="71" y="899" width="26" height="32" rx="10"/>
         <rect x="69" y="897" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="917">?</text>
         <rect x="71" y="943" width="30" height="32" rx="10"/>
         <rect x="69" y="941" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="961">^</text>
         <rect x="71" y="987" width="28" height="32" rx="10"/>
         <rect x="69" y="985" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1005">_</text>
         <rect x="71" y="1031" width="28" height="32" rx="10"/>
         <rect x="69" y="1029" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1049">`</text>
         <rect x="71" y="1075" width="26" height="32" rx="10"/>
         <rect x="69" y="1073" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1093">|</text>
         <rect x="71" y="1119" width="30" height="32" rx="10"/>
         <rect x="69" y="1117" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1137">~</text>
         <rect x="71" y="1163" width="26" height="32" rx="10"/>
         <rect x="69" y="1161" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1181">-</text>
         <path class="line" d="m17 33 h2 m40 0 h10 m32 0 h10 m0 0 h30 m-102 0 h20 m82 0 h20 m-122 0 q10 0 10 10 m102 0 q0 -10 10 -10 m-112 10 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m62 0 h10 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m62 0 h10 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m60 0 h10 m0 0 h2 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m34 0 h10 m0 0 h28 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-122 -1144 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m122 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-122 0 h10 m0 0 h112 m23 32 h-3"/>
         <polygon points="191 33 199 29 199 37"/>
         <polygon points="191 33 183 29 183 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#VarJSONUnquoted" title="VarJSONUnquoted">VarJSONUnquoted</a></div>
               <div>         ::= [@0-9A-Za-z \t!$%&amp;()*+./;&lt;&gt;?^_`|~#x2D]+</div>
               <div>          /* ws: explicit */</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#VarJSONPair" title="VarJSONPair">VarJSONPair</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="FilesUnquoted">FilesUnquoted:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="201" height="1153">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <polygon points="71 35 78 19 126 19 133 35 126 51 78 51"/>
         <polygon points="69 33 76 17 124 17 131 33 124 49 76 49" class="regexp"/>
         <text class="regexp" x="84" y="37">[0-9]</text>
         <polygon points="71 79 78 63 126 63 133 79 126 95 78 95"/>
         <polygon points="69 77 76 61 124 61 131 77 124 93 76 93" class="regexp"/>
         <text class="regexp" x="84" y="81">[A-Z]</text>
         <polygon points="71 123 78 107 124 107 131 123 124 139 78 139"/>
         <polygon points="69 121 76 105 122 105 129 121 122 137 76 137" class="regexp"/>
         <text class="regexp" x="84" y="125">[a-z]</text>
         <rect x="71" y="151" width="24" height="32" rx="10"/>
         <rect x="69" y="149" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="169"/>
         <rect x="71" y="195" width="28" height="32" rx="10"/>
         <rect x="69" y="193" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="213">\</text>
         <rect x="71" y="239" width="24" height="32" rx="10"/>
         <rect x="69" y="237" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="257">t</text>
         <rect x="71" y="283" width="24" height="32" rx="10"/>
         <rect x="69" y="281" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="301">!</text>
         <rect x="71" y="327" width="28" height="32" rx="10"/>
         <rect x="69" y="325" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="345">$</text>
         <rect x="71" y="371" width="34" height="32" rx="10"/>
         <rect x="69" y="369" width="34" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="389">%</text>
         <rect x="71" y="415" width="30" height="32" rx="10"/>
         <rect x="69" y="413" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="433">&amp;</text>
         <rect x="71" y="459" width="26" height="32" rx="10"/>
         <rect x="69" y="457" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="477">(</text>
         <rect x="71" y="503" width="26" height="32" rx="10"/>
         <rect x="69" y="501" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="521">)</text>
         <rect x="71" y="547" width="28" height="32" rx="10"/>
         <rect x="69" y="545" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="565">*</text>
         <rect x="71" y="591" width="30" height="32" rx="10"/>
         <rect x="69" y="589" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="609">+</text>
         <rect x="71" y="635" width="24" height="32" rx="10"/>
         <rect x="69" y="633" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="653">.</text>
         <rect x="71" y="679" width="28" height="32" rx="10"/>
         <rect x="69" y="677" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="697">/</text>
         <rect x="71" y="723" width="24" height="32" rx="10"/>
         <rect x="69" y="721" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="741">;</text>
         <rect x="71" y="767" width="30" height="32" rx="10"/>
         <rect x="69" y="765" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="785">&lt;</text>
         <rect x="71" y="811" width="30" height="32" rx="10"/>
         <rect x="69" y="809" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="829">&gt;</text>
         <rect x="71" y="855" width="26" height="32" rx="10"/>
         <rect x="69" y="853" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="873">?</text>
         <rect x="71" y="899" width="30" height="32" rx="10"/>
         <rect x="69" y="897" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="917">^</text>
         <rect x="71" y="943" width="28" height="32" rx="10"/>
         <rect x="69" y="941" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="961">_</text>
         <rect x="71" y="987" width="28" height="32" rx="10"/>
         <rect x="69" y="985" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1005">`</text>
         <rect x="71" y="1031" width="26" height="32" rx="10"/>
         <rect x="69" y="1029" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1049">|</text>
         <rect x="71" y="1075" width="30" height="32" rx="10"/>
         <rect x="69" y="1073" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1093">~</text>
         <rect x="71" y="1119" width="26" height="32" rx="10"/>
         <rect x="69" y="1117" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1137">-</text>
         <path class="line" d="m17 33 h2 m40 0 h10 m62 0 h10 m-102 0 h20 m82 0 h20 m-122 0 q10 0 10 10 m102 0 q0 -10 10 -10 m-112 10 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m62 0 h10 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m60 0 h10 m0 0 h2 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m34 0 h10 m0 0 h28 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-122 -1100 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m122 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-122 0 h10 m0 0 h112 m23 32 h-3"/>
         <polygon points="191 33 199 29 199 37"/>
         <polygon points="191 33 183 29 183 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#FilesUnquoted" title="FilesUnquoted">FilesUnquoted</a></div>
               <div>         ::= [0-9A-Za-z \t!$%&amp;()*+./;&lt;&gt;?^_`|~#x2D]+</div>
               <div>          /* ws: explicit */</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#FilesPair" title="FilesPair">FilesPair</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="QuotedString">QuotedString:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="321" height="115">
         <polygon points="9 51 1 47 1 55"/>
         <polygon points="17 51 9 47 9 55"/>
         <rect x="51" y="37" width="24" height="32" rx="10"/>
         <rect x="49" y="35" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="55">'</text>
         <rect x="51" y="81" width="26" height="32" rx="10"/>
         <rect x="49" y="79" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="99">"</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Char" xlink:title="Char">
            <rect x="137" y="3" width="50" height="32"/>
            <rect x="135" y="1" width="50" height="32" class="nonterminal"/>
            <text class="nonterminal" x="145" y="21">Char</text></a><rect x="247" y="37" width="24" height="32" rx="10"/>
         <rect x="245" y="35" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="255" y="55">'</text>
         <rect x="247" y="81" width="26" height="32" rx="10"/>
         <rect x="245" y="79" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="255" y="99">"</text>
         <path class="line" d="m17 51 h2 m20 0 h10 m24 0 h10 m0 0 h2 m-66 0 h20 m46 0 h20 m-86 0 q10 0 10 10 m66 0 q0 -10 10 -10 m-76 10 v24 m66 0 v-24 m-66 24 q0 10 10 10 m46 0 q10 0 10 -10 m-56 10 h10 m26 0 h10 m40 -44 h10 m0 0 h60 m-90 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -14 q0 -10 10 -10 m70 34 l20 0 m-20 0 q10 0 10 -10 l0 -14 q0 -10 -10 -10 m-70 0 h10 m50 0 h10 m40 34 h10 m24 0 h10 m0 0 h2 m-66 0 h20 m46 0 h20 m-86 0 q10 0 10 10 m66 0 q0 -10 10 -10 m-76 10 v24 m66 0 v-24 m-66 24 q0 10 10 10 m46 0 q10 0 10 -10 m-56 10 h10 m26 0 h10 m23 -44 h-3"/>
         <polygon points="311 51 319 47 319 55"/>
         <polygon points="311 51 303 47 303 55"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#QuotedString" title="QuotedString">QuotedString</a></div>
               <div>         ::= ['"] <a href="#Char" title="Char">Char</a>* ['"]</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#FilesPair" title="FilesPair">FilesPair</a></li>
            <li><a href="#HeaderPair" title="HeaderPair">HeaderPair</a></li>
            <li><a href="#Pair" title="Pair">Pair</a></li>
            <li><a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></li>
            <li><a href="#VarJSONPair" title="VarJSONPair">VarJSONPair</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Unquoted">Unquoted:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="201" height="1197">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <polygon points="71 35 78 19 126 19 133 35 126 51 78 51"/>
         <polygon points="69 33 76 17 124 17 131 33 124 49 76 49" class="regexp"/>
         <text class="regexp" x="84" y="37">[0-9]</text>
         <polygon points="71 79 78 63 126 63 133 79 126 95 78 95"/>
         <polygon points="69 77 76 61 124 61 131 77 124 93 76 93" class="regexp"/>
         <text class="regexp" x="84" y="81">[A-Z]</text>
         <polygon points="71 123 78 107 124 107 131 123 124 139 78 139"/>
         <polygon points="69 121 76 105 122 105 129 121 122 137 76 137" class="regexp"/>
         <text class="regexp" x="84" y="125">[a-z]</text>
         <rect x="71" y="151" width="24" height="32" rx="10"/>
         <rect x="69" y="149" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="169"/>
         <rect x="71" y="195" width="28" height="32" rx="10"/>
         <rect x="69" y="193" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="213">\</text>
         <rect x="71" y="239" width="24" height="32" rx="10"/>
         <rect x="69" y="237" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="257">t</text>
         <rect x="71" y="283" width="24" height="32" rx="10"/>
         <rect x="69" y="281" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="301">!</text>
         <rect x="71" y="327" width="28" height="32" rx="10"/>
         <rect x="69" y="325" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="345">$</text>
         <rect x="71" y="371" width="34" height="32" rx="10"/>
         <rect x="69" y="369" width="34" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="389">%</text>
         <rect x="71" y="415" width="30" height="32" rx="10"/>
         <rect x="69" y="413" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="433">&amp;</text>
         <rect x="71" y="459" width="26" height="32" rx="10"/>
         <rect x="69" y="457" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="477">(</text>
         <rect x="71" y="503" width="26" height="32" rx="10"/>
         <rect x="69" y="501" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="521">)</text>
         <rect x="71" y="547" width="28" height="32" rx="10"/>
         <rect x="69" y="545" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="565">*</text>
         <rect x="71" y="591" width="30" height="32" rx="10"/>
         <rect x="69" y="589" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="609">+</text>
         <rect x="71" y="635" width="24" height="32" rx="10"/>
         <rect x="69" y="633" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="653">.</text>
         <rect x="71" y="679" width="28" height="32" rx="10"/>
         <rect x="69" y="677" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="697">/</text>
         <rect x="71" y="723" width="24" height="32" rx="10"/>
         <rect x="69" y="721" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="741">;</text>
         <rect x="71" y="767" width="30" height="32" rx="10"/>
         <rect x="69" y="765" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="785">&lt;</text>
         <rect x="71" y="811" width="30" height="32" rx="10"/>
         <rect x="69" y="809" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="829">=</text>
         <rect x="71" y="855" width="30" height="32" rx="10"/>
         <rect x="69" y="853" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="873">&gt;</text>
         <rect x="71" y="899" width="26" height="32" rx="10"/>
         <rect x="69" y="897" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="917">?</text>
         <rect x="71" y="943" width="30" height="32" rx="10"/>
         <rect x="69" y="941" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="961">^</text>
         <rect x="71" y="987" width="28" height="32" rx="10"/>
         <rect x="69" y="985" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1005">_</text>
         <rect x="71" y="1031" width="28" height="32" rx="10"/>
         <rect x="69" y="1029" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1049">`</text>
         <rect x="71" y="1075" width="26" height="32" rx="10"/>
         <rect x="69" y="1073" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1093">|</text>
         <rect x="71" y="1119" width="30" height="32" rx="10"/>
         <rect x="69" y="1117" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1137">~</text>
         <rect x="71" y="1163" width="26" height="32" rx="10"/>
         <rect x="69" y="1161" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="79" y="1181">-</text>
         <path class="line" d="m17 33 h2 m40 0 h10 m62 0 h10 m-102 0 h20 m82 0 h20 m-122 0 q10 0 10 10 m102 0 q0 -10 10 -10 m-112 10 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m62 0 h10 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m60 0 h10 m0 0 h2 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m34 0 h10 m0 0 h28 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m24 0 h10 m0 0 h38 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m28 0 h10 m0 0 h34 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m30 0 h10 m0 0 h32 m-92 -10 v20 m102 0 v-20 m-102 20 v24 m102 0 v-24 m-102 24 q0 10 10 10 m82 0 q10 0 10 -10 m-92 10 h10 m26 0 h10 m0 0 h36 m-122 -1144 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m122 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-122 0 h10 m0 0 h112 m23 32 h-3"/>
         <polygon points="191 33 199 29 199 37"/>
         <polygon points="191 33 183 29 183 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Unquoted" title="Unquoted">Unquoted</a> ::= [0-9A-Za-z \t!$%&amp;()*+./;&lt;=&gt;?^_`|~#x2D]+</div>
               <div>          /* ws: explicit */</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#HeaderPair" title="HeaderPair">HeaderPair</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="JSONType">JSONType:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="203" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#ComplexType" xlink:title="ComplexType">
            <rect x="51" y="3" width="104" height="32"/>
            <rect x="49" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">ComplexType</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#PrimitiveType" xlink:title="PrimitiveType">
            <rect x="51" y="47" width="104" height="32"/>
            <rect x="49" y="45" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">PrimitiveType</text></a><path class="line" d="m17 17 h2 m20 0 h10 m104 0 h10 m-144 0 h20 m124 0 h20 m-164 0 q10 0 10 10 m144 0 q0 -10 10 -10 m-154 10 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m104 0 h10 m23 -44 h-3"/>
         <polygon points="193 17 201 13 201 21"/>
         <polygon points="193 17 185 13 185 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#JSONType" title="JSONType">JSONType</a> ::= <a href="#ComplexType" title="ComplexType">ComplexType</a></div>
               <div>           | <a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#DataInput" title="DataInput">DataInput</a></li>
            <li><a href="#List" title="List">List</a></li>
            <li><a href="#Pair" title="Pair">Pair</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="ComplexType">ComplexType:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="145" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#List" xlink:title="List">
            <rect x="51" y="3" width="42" height="32"/>
            <rect x="49" y="1" width="42" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">List</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Map" xlink:title="Map">
            <rect x="51" y="47" width="46" height="32"/>
            <rect x="49" y="45" width="46" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">Map</text></a><path class="line" d="m17 17 h2 m20 0 h10 m42 0 h10 m0 0 h4 m-86 0 h20 m66 0 h20 m-106 0 q10 0 10 10 m86 0 q0 -10 10 -10 m-96 10 v24 m86 0 v-24 m-86 24 q0 10 10 10 m66 0 q10 0 10 -10 m-76 10 h10 m46 0 h10 m23 -44 h-3"/>
         <polygon points="135 17 143 13 143 21"/>
         <polygon points="135 17 127 13 127 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#ComplexType" title="ComplexType">ComplexType</a></div>
               <div>         ::= <a href="#List" title="List">List</a></div>
               <div>           | <a href="#Map" title="Map">Map</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#JSONType" title="JSONType">JSONType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="PrimitiveType">PrimitiveType:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="203" height="169">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Null" xlink:title="Null">
            <rect x="51" y="3" width="44" height="32"/>
            <rect x="49" y="1" width="44" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">Null</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Boolean" xlink:title="Boolean">
            <rect x="51" y="47" width="72" height="32"/>
            <rect x="49" y="45" width="72" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">Boolean</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="51" y="91" width="104" height="32"/>
            <rect x="49" y="89" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="109">QuotedString</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Number" xlink:title="Number">
            <rect x="51" y="135" width="70" height="32"/>
            <rect x="49" y="133" width="70" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="153">Number</text></a><path class="line" d="m17 17 h2 m20 0 h10 m44 0 h10 m0 0 h60 m-144 0 h20 m124 0 h20 m-164 0 q10 0 10 10 m144 0 q0 -10 10 -10 m-154 10 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m72 0 h10 m0 0 h32 m-134 -10 v20 m144 0 v-20 m-144 20 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m104 0 h10 m-134 -10 v20 m144 0 v-20 m-144 20 v24 m144 0 v-24 m-144 24 q0 10 10 10 m124 0 q10 0 10 -10 m-134 10 h10 m70 0 h10 m0 0 h34 m23 -132 h-3"/>
         <polygon points="193 17 201 13 201 21"/>
         <polygon points="193 17 185 13 185 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></div>
               <div>         ::= <a href="#Null" title="Null">Null</a></div>
               <div>           | <a href="#Boolean" title="Boolean">Boolean</a></div>
               <div>           | <a href="#QuotedString" title="QuotedString">QuotedString</a></div>
               <div>           | <a href="#Number" title="Number">Number</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#JSONType" title="JSONType">JSONType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Map">Map:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="427" height="85">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <rect x="31" y="19" width="28" height="32" rx="10"/>
         <rect x="29" y="17" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="37">{</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Pair" xlink:title="Pair">
            <rect x="99" y="51" width="44" height="32"/>
            <rect x="97" y="49" width="44" height="32" class="nonterminal"/>
            <text class="nonterminal" x="107" y="69">Pair</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Pair" xlink:title="Pair">
            <rect x="223" y="19" width="44" height="32"/>
            <rect x="221" y="17" width="44" height="32" class="nonterminal"/>
            <text class="nonterminal" x="231" y="37">Pair</text></a><rect x="287" y="19" width="24" height="32" rx="10"/>
         <rect x="285" y="17" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="295" y="37">,</text>
         <rect x="371" y="19" width="28" height="32" rx="10"/>
         <rect x="369" y="17" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="379" y="37">}</text>
         <path class="line" d="m17 33 h2 m0 0 h10 m28 0 h10 m20 0 h10 m0 0 h54 m-84 0 h20 m64 0 h20 m-104 0 q10 0 10 10 m84 0 q0 -10 10 -10 m-94 10 v12 m84 0 v-12 m-84 12 q0 10 10 10 m64 0 q10 0 10 -10 m-74 10 h10 m44 0 h10 m60 -32 h10 m44 0 h10 m0 0 h10 m24 0 h10 m-128 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m108 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-108 0 h10 m0 0 h98 m-148 32 h20 m148 0 h20 m-188 0 q10 0 10 10 m168 0 q0 -10 10 -10 m-178 10 v14 m168 0 v-14 m-168 14 q0 10 10 10 m148 0 q10 0 10 -10 m-158 10 h10 m0 0 h138 m20 -34 h10 m28 0 h10 m3 0 h-3"/>
         <polygon points="417 33 425 29 425 37"/>
         <polygon points="417 33 409 29 409 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Map" title="Map">Map</a>      ::= '{' <a href="#Pair" title="Pair">Pair</a>? ( <a href="#Pair" title="Pair">Pair</a> ',' )* '}'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#ComplexType" title="ComplexType">ComplexType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="List">List:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="499" height="85">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/>
         <rect x="31" y="19" width="26" height="32" rx="10"/>
         <rect x="29" y="17" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="37">[</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#JSONType" xlink:title="JSONType">
            <rect x="97" y="51" width="82" height="32"/>
            <rect x="95" y="49" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="105" y="69">JSONType</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#JSONType" xlink:title="JSONType">
            <rect x="259" y="19" width="82" height="32"/>
            <rect x="257" y="17" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="267" y="37">JSONType</text></a><rect x="361" y="19" width="24" height="32" rx="10"/>
         <rect x="359" y="17" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="369" y="37">,</text>
         <rect x="445" y="19" width="26" height="32" rx="10"/>
         <rect x="443" y="17" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="453" y="37">]</text>
         <path class="line" d="m17 33 h2 m0 0 h10 m26 0 h10 m20 0 h10 m0 0 h92 m-122 0 h20 m102 0 h20 m-142 0 q10 0 10 10 m122 0 q0 -10 10 -10 m-132 10 v12 m122 0 v-12 m-122 12 q0 10 10 10 m102 0 q10 0 10 -10 m-112 10 h10 m82 0 h10 m60 -32 h10 m82 0 h10 m0 0 h10 m24 0 h10 m-166 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m146 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-146 0 h10 m0 0 h136 m-186 32 h20 m186 0 h20 m-226 0 q10 0 10 10 m206 0 q0 -10 10 -10 m-216 10 v14 m206 0 v-14 m-206 14 q0 10 10 10 m186 0 q10 0 10 -10 m-196 10 h10 m0 0 h176 m20 -34 h10 m26 0 h10 m3 0 h-3"/>
         <polygon points="489 33 497 29 497 37"/>
         <polygon points="489 33 481 29 481 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#List" title="List">List</a>     ::= '[' <a href="#JSONType" title="JSONType">JSONType</a>? ( <a href="#JSONType" title="JSONType">JSONType</a> ',' )* ']'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#ComplexType" title="ComplexType">ComplexType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Pair">Pair:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="309" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#QuotedString" xlink:title="QuotedString">
            <rect x="31" y="3" width="104" height="32"/>
            <rect x="29" y="1" width="104" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">QuotedString</text></a><rect x="155" y="3" width="24" height="32" rx="10"/>
         <rect x="153" y="1" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="163" y="21">:</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#JSONType" xlink:title="JSONType">
            <rect x="199" y="3" width="82" height="32"/>
            <rect x="197" y="1" width="82" height="32" class="nonterminal"/>
            <text class="nonterminal" x="207" y="21">JSONType</text></a><path class="line" d="m17 17 h2 m0 0 h10 m104 0 h10 m0 0 h10 m24 0 h10 m0 0 h10 m82 0 h10 m3 0 h-3"/>
         <polygon points="299 17 307 13 307 21"/>
         <polygon points="299 17 291 13 291 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Pair" title="Pair">Pair</a>     ::= <a href="#QuotedString" title="QuotedString">QuotedString</a> ':' <a href="#JSONType" title="JSONType">JSONType</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Map" title="Map">Map</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Boolean">Boolean:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="151" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="51" y="3" width="48" height="32" rx="10"/>
         <rect x="49" y="1" width="48" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="21">true</text>
         <rect x="51" y="47" width="52" height="32" rx="10"/>
         <rect x="49" y="45" width="52" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="65">false</text>
         <path class="line" d="m17 17 h2 m20 0 h10 m48 0 h10 m0 0 h4 m-92 0 h20 m72 0 h20 m-112 0 q10 0 10 10 m92 0 q0 -10 10 -10 m-102 10 v24 m92 0 v-24 m-92 24 q0 10 10 10 m72 0 q10 0 10 -10 m-82 10 h10 m52 0 h10 m23 -44 h-3"/>
         <polygon points="141 17 149 13 149 21"/>
         <polygon points="141 17 133 13 133 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Boolean" title="Boolean">Boolean</a>  ::= 'true'</div>
               <div>           | 'false'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Null">Null:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="103" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="31" y="3" width="44" height="32" rx="10"/>
         <rect x="29" y="1" width="44" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="21">null</text>
         <path class="line" d="m17 17 h2 m0 0 h10 m44 0 h10 m3 0 h-3"/>
         <polygon points="93 17 101 13 101 21"/>
         <polygon points="93 17 85 13 85 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Null" title="Null">Null</a>     ::= 'null'</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Number">Number:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="395" height="69">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Integer" xlink:title="Integer">
            <rect x="31" y="3" width="66" height="32"/>
            <rect x="29" y="1" width="66" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">Integer</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Fraction" xlink:title="Fraction">
            <rect x="137" y="35" width="70" height="32"/>
            <rect x="135" y="33" width="70" height="32" class="nonterminal"/>
            <text class="nonterminal" x="145" y="53">Fraction</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Exponent" xlink:title="Exponent">
            <rect x="267" y="35" width="80" height="32"/>
            <rect x="265" y="33" width="80" height="32" class="nonterminal"/>
            <text class="nonterminal" x="275" y="53">Exponent</text></a><path class="line" d="m17 17 h2 m0 0 h10 m66 0 h10 m20 0 h10 m0 0 h80 m-110 0 h20 m90 0 h20 m-130 0 q10 0 10 10 m110 0 q0 -10 10 -10 m-120 10 v12 m110 0 v-12 m-110 12 q0 10 10 10 m90 0 q10 0 10 -10 m-100 10 h10 m70 0 h10 m40 -32 h10 m0 0 h90 m-120 0 h20 m100 0 h20 m-140 0 q10 0 10 10 m120 0 q0 -10 10 -10 m-130 10 v12 m120 0 v-12 m-120 12 q0 10 10 10 m100 0 q10 0 10 -10 m-110 10 h10 m80 0 h10 m23 -32 h-3"/>
         <polygon points="385 17 393 13 393 21"/>
         <polygon points="385 17 377 13 377 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Number" title="Number">Number</a>   ::= <a href="#Integer" title="Integer">Integer</a> <a href="#Fraction" title="Fraction">Fraction</a>? <a href="#Exponent" title="Exponent">Exponent</a>?</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#PrimitiveType" title="PrimitiveType">PrimitiveType</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Exponent">Exponent:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="309" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="51" y="3" width="28" height="32" rx="10"/>
         <rect x="49" y="1" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="21">e</text>
         <rect x="51" y="47" width="28" height="32" rx="10"/>
         <rect x="49" y="45" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="65">E</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Sign" xlink:title="Sign">
            <rect x="139" y="35" width="46" height="32"/>
            <rect x="137" y="33" width="46" height="32" class="nonterminal"/>
            <text class="nonterminal" x="147" y="53">Sign</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Digits" xlink:title="Digits">
            <rect x="225" y="3" width="56" height="32"/>
            <rect x="223" y="1" width="56" height="32" class="nonterminal"/>
            <text class="nonterminal" x="233" y="21">Digits</text></a><path class="line" d="m17 17 h2 m20 0 h10 m28 0 h10 m-68 0 h20 m48 0 h20 m-88 0 q10 0 10 10 m68 0 q0 -10 10 -10 m-78 10 v24 m68 0 v-24 m-68 24 q0 10 10 10 m48 0 q10 0 10 -10 m-58 10 h10 m28 0 h10 m40 -44 h10 m0 0 h56 m-86 0 h20 m66 0 h20 m-106 0 q10 0 10 10 m86 0 q0 -10 10 -10 m-96 10 v12 m86 0 v-12 m-86 12 q0 10 10 10 m66 0 q10 0 10 -10 m-76 10 h10 m46 0 h10 m20 -32 h10 m56 0 h10 m3 0 h-3"/>
         <polygon points="299 17 307 13 307 21"/>
         <polygon points="299 17 291 13 291 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Exponent" title="Exponent">Exponent</a> ::= [eE] <a href="#Sign" title="Sign">Sign</a>? <a href="#Digits" title="Digits">Digits</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Number" title="Number">Number</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Fraction">Fraction:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="159" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="31" y="3" width="24" height="32" rx="10"/>
         <rect x="29" y="1" width="24" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="39" y="21">.</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Digits" xlink:title="Digits">
            <rect x="75" y="3" width="56" height="32"/>
            <rect x="73" y="1" width="56" height="32" class="nonterminal"/>
            <text class="nonterminal" x="83" y="21">Digits</text></a><path class="line" d="m17 17 h2 m0 0 h10 m24 0 h10 m0 0 h10 m56 0 h10 m3 0 h-3"/>
         <polygon points="149 17 157 13 157 21"/>
         <polygon points="149 17 141 13 141 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Fraction" title="Fraction">Fraction</a> ::= '.' <a href="#Digits" title="Digits">Digits</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Number" title="Number">Number</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Integer">Integer:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="201" height="169">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule4" xlink:title="IntegerRule4">
            <rect x="51" y="3" width="102" height="32"/>
            <rect x="49" y="1" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="21">IntegerRule4</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule3" xlink:title="IntegerRule3">
            <rect x="51" y="47" width="102" height="32"/>
            <rect x="49" y="45" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="65">IntegerRule3</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule2" xlink:title="IntegerRule2">
            <rect x="51" y="91" width="102" height="32"/>
            <rect x="49" y="89" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="109">IntegerRule2</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule1" xlink:title="IntegerRule1">
            <rect x="51" y="135" width="102" height="32"/>
            <rect x="49" y="133" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="153">IntegerRule1</text></a><path class="line" d="m17 17 h2 m20 0 h10 m102 0 h10 m-142 0 h20 m122 0 h20 m-162 0 q10 0 10 10 m142 0 q0 -10 10 -10 m-152 10 v24 m142 0 v-24 m-142 24 q0 10 10 10 m122 0 q10 0 10 -10 m-132 10 h10 m102 0 h10 m-132 -10 v20 m142 0 v-20 m-142 20 v24 m142 0 v-24 m-142 24 q0 10 10 10 m122 0 q10 0 10 -10 m-132 10 h10 m102 0 h10 m-132 -10 v20 m142 0 v-20 m-142 20 v24 m142 0 v-24 m-142 24 q0 10 10 10 m122 0 q10 0 10 -10 m-132 10 h10 m102 0 h10 m23 -132 h-3"/>
         <polygon points="191 17 199 13 199 21"/>
         <polygon points="191 17 183 13 183 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Integer" title="Integer">Integer</a>  ::= <a href="#IntegerRule4" title="IntegerRule4">IntegerRule4</a></div>
               <div>           | <a href="#IntegerRule3" title="IntegerRule3">IntegerRule3</a></div>
               <div>           | <a href="#IntegerRule2" title="IntegerRule2">IntegerRule2</a></div>
               <div>           | <a href="#IntegerRule1" title="IntegerRule1">IntegerRule1</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Number" title="Number">Number</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="IntegerRule1">IntegerRule1:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="107" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Digit" xlink:title="Digit">
            <rect x="31" y="3" width="48" height="32"/>
            <rect x="29" y="1" width="48" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">Digit</text></a><path class="line" d="m17 17 h2 m0 0 h10 m48 0 h10 m3 0 h-3"/>
         <polygon points="97 17 105 13 105 21"/>
         <polygon points="97 17 89 13 89 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#IntegerRule1" title="IntegerRule1">IntegerRule1</a></div>
               <div>         ::= <a href="#Digit" title="Digit">Digit</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Integer" title="Integer">Integer</a></li>
            <li><a href="#IntegerRule3" title="IntegerRule3">IntegerRule3</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="IntegerRule2">IntegerRule2:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="209" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#OneNine" xlink:title="OneNine">
            <rect x="31" y="3" width="74" height="32"/>
            <rect x="29" y="1" width="74" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">OneNine</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Digits" xlink:title="Digits">
            <rect x="125" y="3" width="56" height="32"/>
            <rect x="123" y="1" width="56" height="32" class="nonterminal"/>
            <text class="nonterminal" x="133" y="21">Digits</text></a><path class="line" d="m17 17 h2 m0 0 h10 m74 0 h10 m0 0 h10 m56 0 h10 m3 0 h-3"/>
         <polygon points="199 17 207 13 207 21"/>
         <polygon points="199 17 191 13 191 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#IntegerRule2" title="IntegerRule2">IntegerRule2</a></div>
               <div>         ::= <a href="#OneNine" title="OneNine">OneNine</a> <a href="#Digits" title="Digits">Digits</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Integer" title="Integer">Integer</a></li>
            <li><a href="#IntegerRule4" title="IntegerRule4">IntegerRule4</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="IntegerRule3">IntegerRule3:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="227" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Sign" xlink:title="Sign">
            <rect x="31" y="3" width="46" height="32"/>
            <rect x="29" y="1" width="46" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">Sign</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule1" xlink:title="IntegerRule1">
            <rect x="97" y="3" width="102" height="32"/>
            <rect x="95" y="1" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="105" y="21">IntegerRule1</text></a><path class="line" d="m17 17 h2 m0 0 h10 m46 0 h10 m0 0 h10 m102 0 h10 m3 0 h-3"/>
         <polygon points="217 17 225 13 225 21"/>
         <polygon points="217 17 209 13 209 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#IntegerRule3" title="IntegerRule3">IntegerRule3</a></div>
               <div>         ::= <a href="#Sign" title="Sign">Sign</a> <a href="#IntegerRule1" title="IntegerRule1">IntegerRule1</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Integer" title="Integer">Integer</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="IntegerRule4">IntegerRule4:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="227" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Sign" xlink:title="Sign">
            <rect x="31" y="3" width="46" height="32"/>
            <rect x="29" y="1" width="46" height="32" class="nonterminal"/>
            <text class="nonterminal" x="39" y="21">Sign</text></a><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#IntegerRule2" xlink:title="IntegerRule2">
            <rect x="97" y="3" width="102" height="32"/>
            <rect x="95" y="1" width="102" height="32" class="nonterminal"/>
            <text class="nonterminal" x="105" y="21">IntegerRule2</text></a><path class="line" d="m17 17 h2 m0 0 h10 m46 0 h10 m0 0 h10 m102 0 h10 m3 0 h-3"/>
         <polygon points="217 17 225 13 225 21"/>
         <polygon points="217 17 209 13 209 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#IntegerRule4" title="IntegerRule4">IntegerRule4</a></div>
               <div>         ::= <a href="#Sign" title="Sign">Sign</a> <a href="#IntegerRule2" title="IntegerRule2">IntegerRule2</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Integer" title="Integer">Integer</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Digits">Digits:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="147" height="53">
         <polygon points="9 33 1 29 1 37"/>
         <polygon points="17 33 9 29 9 37"/><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#Digit" xlink:title="Digit">
            <rect x="51" y="19" width="48" height="32"/>
            <rect x="49" y="17" width="48" height="32" class="nonterminal"/>
            <text class="nonterminal" x="59" y="37">Digit</text></a><path class="line" d="m17 33 h2 m20 0 h10 m48 0 h10 m-88 0 l20 0 m-1 0 q-9 0 -9 -10 l0 -12 q0 -10 10 -10 m68 32 l20 0 m-20 0 q10 0 10 -10 l0 -12 q0 -10 -10 -10 m-68 0 h10 m0 0 h58 m23 32 h-3"/>
         <polygon points="137 33 145 29 145 37"/>
         <polygon points="137 33 129 29 129 37"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Digits" title="Digits">Digits</a>   ::= <a href="#Digit" title="Digit">Digit</a> <a href="#Digit" title="Digit">Digit</a>*</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Exponent" title="Exponent">Exponent</a></li>
            <li><a href="#Fraction" title="Fraction">Fraction</a></li>
            <li><a href="#IntegerRule2" title="IntegerRule2">IntegerRule2</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Digit">Digit:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="221" height="69">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="51" y="35" width="28" height="32" rx="10"/>
         <rect x="49" y="33" width="28" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="53">0</text><a xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#OneNine" xlink:title="OneNine">
            <rect x="119" y="3" width="74" height="32"/>
            <rect x="117" y="1" width="74" height="32" class="nonterminal"/>
            <text class="nonterminal" x="127" y="21">OneNine</text></a><path class="line" d="m17 17 h2 m20 0 h10 m0 0 h38 m-68 0 h20 m48 0 h20 m-88 0 q10 0 10 10 m68 0 q0 -10 10 -10 m-78 10 v12 m68 0 v-12 m-68 12 q0 10 10 10 m48 0 q10 0 10 -10 m-58 10 h10 m28 0 h10 m20 -32 h10 m74 0 h10 m3 0 h-3"/>
         <polygon points="211 17 219 13 219 21"/>
         <polygon points="211 17 203 13 203 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Digit" title="Digit">Digit</a>    ::= '0'? <a href="#OneNine" title="OneNine">OneNine</a></div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Digits" title="Digits">Digits</a></li>
            <li><a href="#IntegerRule1" title="IntegerRule1">IntegerRule1</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="OneNine">OneNine:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="121" height="37">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <polygon points="31 19 38 3 86 3 93 19 86 35 38 35"/>
         <polygon points="29 17 36 1 84 1 91 17 84 33 36 33" class="regexp"/>
         <text class="regexp" x="44" y="21">[1-9]</text>
         <path class="line" d="m17 17 h2 m0 0 h10 m62 0 h10 m3 0 h-3"/>
         <polygon points="111 17 119 13 119 21"/>
         <polygon points="111 17 103 13 103 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#OneNine" title="OneNine">OneNine</a>  ::= [1-9]</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Digit" title="Digit">Digit</a></li>
            <li><a href="#IntegerRule2" title="IntegerRule2">IntegerRule2</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><p xmlns:xhtml="http://www.w3.org/1999/xhtml" style="font-size: 14px; font-weight:bold"><a name="Sign">Sign:</a></p><svg xmlns="http://www.w3.org/2000/svg" width="129" height="81">
         <polygon points="9 17 1 13 1 21"/>
         <polygon points="17 17 9 13 9 21"/>
         <rect x="51" y="3" width="26" height="32" rx="10"/>
         <rect x="49" y="1" width="26" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="21">-</text>
         <rect x="51" y="47" width="30" height="32" rx="10"/>
         <rect x="49" y="45" width="30" height="32" class="terminal" rx="10"/>
         <text class="terminal" x="59" y="65">+</text>
         <path class="line" d="m17 17 h2 m20 0 h10 m26 0 h10 m0 0 h4 m-70 0 h20 m50 0 h20 m-90 0 q10 0 10 10 m70 0 q0 -10 10 -10 m-80 10 v24 m70 0 v-24 m-70 24 q0 10 10 10 m50 0 q10 0 10 -10 m-60 10 h10 m30 0 h10 m23 -44 h-3"/>
         <polygon points="119 17 127 13 127 21"/>
         <polygon points="119 17 111 13 111 21"/></svg><p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <div class="ebnf"><code>
               <div><a href="#Sign" title="Sign">Sign</a>     ::= [-+]</div></code></div>
      </p>
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">referenced by:
         <ul>
            <li><a href="#Exponent" title="Exponent">Exponent</a></li>
            <li><a href="#IntegerRule3" title="IntegerRule3">IntegerRule3</a></li>
            <li><a href="#IntegerRule4" title="IntegerRule4">IntegerRule4</a></li>
         </ul>
      </p><br xmlns:xhtml="http://www.w3.org/1999/xhtml" /><hr xmlns:xhtml="http://www.w3.org/1999/xhtml" />
      <p xmlns:xhtml="http://www.w3.org/1999/xhtml">
         <table border="0" class="signature">
            <tr>
               <td style="width: 100%"> </td>
               <td valign="top">
                  <nobr class="signature">... generated by <a name="Railroad-Diagram-Generator" class="signature" title="https://www.bottlecaps.de/rr/ui" href="https://www.bottlecaps.de/rr/ui" target="_blank">RR - Railroad Diagram Generator</a></nobr>
               </td>
               <td><a name="Railroad-Diagram-Generator" title="https://www.bottlecaps.de/rr/ui" href="https://www.bottlecaps.de/rr/ui" target="_blank"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16">
                        <g transform="scale(0.178)">
                           <circle cx="45" cy="45" r="45" style="stroke:none; fill:#FFCC00"/>
                           <circle cx="45" cy="45" r="42" style="stroke:#332900; stroke-width:2px; fill:#FFCC00"/>
                           <line x1="15" y1="15" x2="75" y2="75" stroke="#332900" style="stroke-width:9px;"/>
                           <line x1="15" y1="75" x2="75" y2="15" stroke="#332900" style="stroke-width:9px;"/>
                           <text x="7" y="54" style="font-size:26px; font-family:Arial, Sans-serif; font-weight:bold; fill: #332900">R</text>
                           <text x="64" y="54" style="font-size:26px; font-family:Arial, Sans-serif; font-weight:bold; fill: #332900">R</text>
                        </g></svg></a></td>
            </tr>
         </table>
      </p>
   </body>
</html>