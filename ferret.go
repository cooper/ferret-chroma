package ferret

import (
	chroma "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers"
)

// Ferret lexer.
var Ferret = lexers.Register(chroma.MustNewLexer(
	&chroma.Config{
		Name:      "Ferret",
		Aliases:   []string{"ferret", "frt"},
		Filenames: []string{"*.frt", "*.frtdoc"},
		DotAll:    true,
	},
	chroma.Rules{
		"balanced-regex": {
			{`/(\\\\|\\[^\\]|[^\\/])*/[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`!(\\\\|\\[^\\]|[^\\!])*![egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`\\(\\\\|[^\\])*\\[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`\{(\\\\|\\[^\\]|[^\\}])*\}[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`<(\\\\|\\[^\\]|[^\\>])*>[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`\[(\\\\|\\[^\\]|[^\\\]])*\][egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`\((\\\\|\\[^\\]|[^\\)])*\)[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`@(\\\\|\\[^\\]|[^\\@])*@[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`%(\\\\|\\[^\\]|[^\\%])*%[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
			{`\$(\\\\|\\[^\\]|[^\\$])*\$[egimosx]*`, chroma.LiteralStringRegex, chroma.Pop(1)},
		},
		"root": {
			{`\A\#!.+?$`, chroma.CommentHashbang, nil},
			{`\#.*?$`, chroma.CommentSingle, nil},
			{chroma.Words(``, `\b`, `continue`, `do`, `else`, `for`, `in`, `if`, `last`, `next`, `redo`, `defer`, `until`, `while`, `switch`, `case`, `before`, `after`, `load`, `stop`, `fail`, `catch`, `throw`, `inside`, `can`, `isa`, `transform`, `satisfies`, `gather`, `take`), chroma.Keyword, nil},
			{chroma.Words(``, `\b`, `share`, `var`, `want`, `need`), chroma.KeywordDeclaration, nil},
			{`(true|false|undefined)\b`, chroma.KeywordConstant, nil},
			{`s/(\\\\|\\[^\\]|[^\\/])*/(\\\\|\\[^\\]|[^\\/])*/[egimosx]*`, chroma.LiteralStringRegex, nil},
			{`s!(\\\\|\\!|[^!])*!(\\\\|\\!|[^!])*![egimosx]*`, chroma.LiteralStringRegex, nil},
			{`s\\(\\\\|[^\\])*\\(\\\\|[^\\])*\\[egimosx]*`, chroma.LiteralStringRegex, nil},
			{`s@(\\\\|\\[^\\]|[^\\@])*@(\\\\|\\[^\\]|[^\\@])*@[egimosx]*`, chroma.LiteralStringRegex, nil},
			{`s%(\\\\|\\[^\\]|[^\\%])*%(\\\\|\\[^\\]|[^\\%])*%[egimosx]*`, chroma.LiteralStringRegex, nil},
			{`s\{(\\\\|\\[^\\]|[^\\}])*\}\s*`, chroma.LiteralStringRegex, chroma.Push("balanced-regex")},
			{`s<(\\\\|\\[^\\]|[^\\>])*>\s*`, chroma.LiteralStringRegex, chroma.Push("balanced-regex")},
			{`s\[(\\\\|\\[^\\]|[^\\\]])*\]\s*`, chroma.LiteralStringRegex, chroma.Push("balanced-regex")},
			{`s\((\\\\|\\[^\\]|[^\\)])*\)\s*`, chroma.LiteralStringRegex, chroma.Push("balanced-regex")},
			{`m?/(\\\\|\\[^\\]|[^\\/\n])*/[gcimosx]*`, chroma.LiteralStringRegex, nil},
			{`m(?=[/!\\{<\[(@%$])`, chroma.LiteralStringRegex, chroma.Push("balanced-regex")},
			{`((?<==~)|(?<=\())\s*/(\\\\|\\[^\\]|[^\\/])*/[gcimosx]*`, chroma.LiteralStringRegex, nil},
			{`\s+`, chroma.Text, nil},
			{chroma.Words(``, `\b`, `dump`, `inspect`, `delay`, `timeout`, `say`, `print`, `delete`, `weaken`, `detail`), chroma.NameBuiltin, nil},
			{`\.\b(proto|init|hashValue|lowercase|uppercase|result|empty|length|sum|sum0|s|lastIndex|even|odd|sqrt|log|log10|exp|e|pi|cbrt|abs|ceil|floor|round|square|name|signature|version|keys|values|nextElement|nextElements|more|iterator|sin|asin|cos|acos|tan|atan|atan2|cot|sec|csc|shift|unshift|pop|push|without|withoutAll|grep|contains|indexOf|map|first|any|all|remove|removeAll|splice|split|reverse|flatten|root|logb|factorial|polar|connect|print|println|say|close|trimPrefix|trimSuffix|hasSuffix|hasPrefix|fill|word|fromWord|match|copy|join|setValue|getValue|deleteValue|weakenValue)\b`, chroma.NameBuiltin, nil},
			{`(<<)([\'"]?)([a-zA-Z_]\w*)(\2;?\n.*?\n)(\3)(\n)`, chroma.ByGroups(chroma.LiteralString, chroma.LiteralString, chroma.LiteralStringDelimiter, chroma.LiteralString, chroma.LiteralStringDelimiter, chroma.Text), nil},
			{`__END__`, chroma.CommentPreproc, chroma.Push("end-part")},
			{`[$@%*.]+`, chroma.NameVariable, chroma.Push("varname")},
			{`\:[a-zA-Z_]\w*`, chroma.LiteralStringSymbol, nil},
			{`0_?[0-7]+(_[0-7]+)*`, chroma.LiteralNumberOct, nil},
			{`0x[0-9A-Fa-f]+(_[0-9A-Fa-f]+)*`, chroma.LiteralNumberHex, nil},
			{`0b[01]+(_[01]+)*`, chroma.LiteralNumberBin, nil},
			{`(?i)(\d*(_\d*)*\.\d+(_\d*)*|\d+(_\d*)*\.\d+(_\d*)*)(e[+-]?\d+)?`, chroma.LiteralNumberFloat, nil},
			{`(?i)\d+(_\d*)*e[+-]?\d+(_\d*)*`, chroma.LiteralNumberFloat, nil},
			{`\d+(_\d+)*`, chroma.LiteralNumberInteger, nil},
			{`'(\\\\|\\[^\\]|[^'\\])*'`, chroma.LiteralString, nil},
			{`"(\\\\|\\[^\\]|[^"\\])*"`, chroma.LiteralString, nil},
			{"`(\\\\\\\\|\\\\[^\\\\]|[^`\\\\])*`", chroma.LiteralStringBacktick, nil},
			{`<([^\s>]+)>`, chroma.LiteralStringRegex, nil},
			{`(package)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.NameNamespace), nil},
			{`(class)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.NameClass), nil},
			{`(load)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.NameNamespace), nil},
			{`(package|class|load|end|on|init)\b`, chroma.Keyword, nil},
			{`(func|method|hook|type|alias|prop)(\s+)`, chroma.ByGroups(chroma.Keyword, chroma.Text), chroma.Push("funcname")},
			{`(op)(\s+)`, chroma.ByGroups(chroma.Keyword, chroma.Text), chroma.Push("opname")},
			{`(\[\]|\^|::|<<|>>|>=|<=>|<=|->|={3}|!=|=~|!~|&&?|\|\||\.{1,2})`, chroma.Operator, nil},
			{`[-+/*%=<>&^|!?\\~]=?`, chroma.Operator, nil},
			{`[()\[\]:;,<>/?{}]`, chroma.Punctuation, nil},
			{`(?<!\$|\.)\b(_?[A-Z]+[_a-zA-Z0-9_x7f-xff\$]*)\b`, chroma.KeywordType, nil},
			{`(?=\w)`, chroma.Name, chroma.Push("name")},
		},
		"varname": {
			{`\s+`, chroma.Text, nil},
			{`\{`, chroma.Punctuation, chroma.Pop(1)},
			{`\)|,`, chroma.Punctuation, chroma.Pop(1)},
			{`\w+::`, chroma.NameNamespace, nil},
			{`[\w:]+`, chroma.NameVariable, chroma.Pop(1)},
		},
		"name": {
			{`[a-zA-Z_]\w*(::[a-zA-Z_]\w*)*(::)?(?=\s*->)`, chroma.NameNamespace, chroma.Pop(1)},
			{`[a-zA-Z_]\w*(::[a-zA-Z_]\w*)*::`, chroma.NameNamespace, chroma.Pop(1)},
			{`[\w:]+`, chroma.Name, chroma.Pop(1)},
			{`[A-Z_]+(?=\W)`, chroma.NameConstant, chroma.Pop(1)},
			{`(?=\W)`, chroma.Text, chroma.Pop(1)},
		},
		"funcname": {
			{`[a-zA-Z_]\w*\??`, chroma.NameFunction, nil},
			{`[^\S\r\n]+`, chroma.Text, nil},
			{`;`, chroma.Punctuation, chroma.Pop(1)},
			{`->`, chroma.Operator, chroma.Pop(1)},
			{`.*?\{`, chroma.Punctuation, chroma.Pop(1)},
			{`\n`, chroma.Text, chroma.Pop(1)},
		},
		"opname": {
			{`(==|=~|[\^%\+\-\*/<>])+`, chroma.NameFunction, nil},
			{`[^\S\r\n]+`, chroma.Text, nil},
			{`;`, chroma.Punctuation, chroma.Pop(1)},
			{`->`, chroma.Operator, chroma.Pop(1)},
			{`.*?\{`, chroma.Punctuation, chroma.Pop(1)},
			{`\n`, chroma.Text, chroma.Pop(1)},
		},
		"cb-string": {
			{`\\[{}\\]`, chroma.LiteralStringOther, nil},
			{`\\`, chroma.LiteralStringOther, nil},
			{`\{`, chroma.LiteralStringOther, chroma.Push("cb-string")},
			{`\}`, chroma.LiteralStringOther, chroma.Pop(1)},
			{`[^{}\\]+`, chroma.LiteralStringOther, nil},
		},
		"rb-string": {
			{`\\[()\\]`, chroma.LiteralStringOther, nil},
			{`\\`, chroma.LiteralStringOther, nil},
			{`\(`, chroma.LiteralStringOther, chroma.Push("rb-string")},
			{`\)`, chroma.LiteralStringOther, chroma.Pop(1)},
			{`[^()]+`, chroma.LiteralStringOther, nil},
		},
		"sb-string": {
			{`\\[\[\]\\]`, chroma.LiteralStringOther, nil},
			{`\\`, chroma.LiteralStringOther, nil},
			{`\[`, chroma.LiteralStringOther, chroma.Push("sb-string")},
			{`\]`, chroma.LiteralStringOther, chroma.Pop(1)},
			{`[^\[\]]+`, chroma.LiteralStringOther, nil},
		},
		"lt-string": {
			{`\\[<>\\]`, chroma.LiteralStringOther, nil},
			{`\\`, chroma.LiteralStringOther, nil},
			{`\<`, chroma.LiteralStringOther, chroma.Push("lt-string")},
			{`\>`, chroma.LiteralStringOther, chroma.Pop(1)},
			{`[^<>]+`, chroma.LiteralStringOther, nil},
		},
		"end-part": {
			{`.+`, chroma.CommentPreproc, chroma.Pop(1)},
		},
	},
))
