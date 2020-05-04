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
			{`(format)(\s+)(\w+)(\s*)(=)(\s*\n)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.Name, chroma.Text, chroma.Punctuation, chroma.Text), chroma.Push("format")},
			{`(eq|lt|gt|le|ge|ne|not|and|or|cmp)\b`, chroma.OperatorWord, nil},
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
			{chroma.Words(``, `\b`, `abs`, `accept`, `alarm`, `atan2`, `bind`, `binmode`, `bless`, `caller`, `chdir`, `chmod`, `chomp`, `chop`, `chown`, `chr`, `chroot`, `close`, `closedir`, `connect`, `continue`, `cos`, `crypt`, `dbmclose`, `dbmopen`, `defined`, `delete`, `die`, `dump`, `each`, `endgrent`, `endhostent`, `endnetent`, `endprotoent`, `endpwent`, `endservent`, `eof`, `eval`, `exec`, `exists`, `exit`, `exp`, `fcntl`, `fileno`, `flock`, `fork`, `format`, `formline`, `getc`, `getgrent`, `getgrgid`, `getgrnam`, `gethostbyaddr`, `gethostbyname`, `gethostent`, `getlogin`, `getnetbyaddr`, `getnetbyname`, `getnetent`, `getpeername`, `getpgrp`, `getppid`, `getpriority`, `getprotobyname`, `getprotobynumber`, `getprotoent`, `getpwent`, `getpwnam`, `getpwuid`, `getservbyname`, `getservbyport`, `getservent`, `getsockname`, `getsockopt`, `glob`, `gmtime`, `goto`, `grep`, `hex`, `import`, `index`, `int`, `ioctl`, `join`, `keys`, `kill`, `last`, `lc`, `lcfirst`, `length`, `link`, `listen`, `local`, `localtime`, `log`, `lstat`, `map`, `mkdir`, `msgctl`, `msgget`, `msgrcv`, `msgsnd`, `my`, `next`, `oct`, `open`, `opendir`, `ord`, `our`, `pack`, `pipe`, `pop`, `pos`, `printf`, `prototype`, `push`, `quotemeta`, `rand`, `read`, `readdir`, `readline`, `readlink`, `readpipe`, `recv`, `redo`, `ref`, `rename`, `reverse`, `rewinddir`, `rindex`, `rmdir`, `scalar`, `seek`, `seekdir`, `select`, `semctl`, `semget`, `semop`, `send`, `setgrent`, `sethostent`, `setnetent`, `setpgrp`, `setpriority`, `setprotoent`, `setpwent`, `setservent`, `setsockopt`, `shift`, `shmctl`, `shmget`, `shmread`, `shmwrite`, `shutdown`, `sin`, `sleep`, `socket`, `socketpair`, `sort`, `splice`, `split`, `sprintf`, `sqrt`, `srand`, `stat`, `study`, `substr`, `symlink`, `syscall`, `sysopen`, `sysread`, `sysseek`, `system`, `syswrite`, `tell`, `telldir`, `tie`, `tied`, `time`, `times`, `tr`, `truncate`, `uc`, `ucfirst`, `umask`, `undef`, `unlink`, `unpack`, `unshift`, `untie`, `utime`, `values`, `vec`, `wait`, `waitpid`, `wantarray`, `warn`, `write`), chroma.NameBuiltin, nil},
			{`((__(DATA|DIE|WARN)__)|(STD(IN|OUT|ERR)))\b`, chroma.NameBuiltinPseudo, nil},
			{`(<<)([\'"]?)([a-zA-Z_]\w*)(\2;?\n.*?\n)(\3)(\n)`, chroma.ByGroups(chroma.LiteralString, chroma.LiteralString, chroma.LiteralStringDelimiter, chroma.LiteralString, chroma.LiteralStringDelimiter, chroma.Text), nil},
			{`__END__`, chroma.CommentPreproc, chroma.Push("end-part")},
			{`\$\^[ADEFHILMOPSTWX]`, chroma.NameVariableGlobal, nil},
			{"\\$[\\\\\\\"\\[\\]'&`+*.,;=%~?@$!<>(^|/-](?!\\w)", chroma.NameVariableGlobal, nil},
			{`[$@%#]+`, chroma.NameVariable, chroma.Push("varname")},
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
			{`(q|qq|qw|qr|qx)\{`, chroma.LiteralStringOther, chroma.Push("cb-string")},
			{`(q|qq|qw|qr|qx)\(`, chroma.LiteralStringOther, chroma.Push("rb-string")},
			{`(q|qq|qw|qr|qx)\[`, chroma.LiteralStringOther, chroma.Push("sb-string")},
			{`(q|qq|qw|qr|qx)\<`, chroma.LiteralStringOther, chroma.Push("lt-string")},
			{`(q|qq|qw|qr|qx)([\W_])(.|\n)*?\2`, chroma.LiteralStringOther, nil},
			{`(package)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.NameNamespace), nil},
			{`(use|require|no)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, chroma.ByGroups(chroma.Keyword, chroma.Text, chroma.NameNamespace), nil},
			{`(sub)(\s+)`, chroma.ByGroups(chroma.Keyword, chroma.Text), chroma.Push("funcname")},
			{chroma.Words(``, `\b`, `no`, `package`, `require`, `use`), chroma.Keyword, nil},
			{`(\[\]|\*\*|::|<<|>>|>=|<=>|<=|={3}|!=|=~|!~|&&?|\|\||\.{1,3})`, chroma.Operator, nil},
			{`[-+/*%=<>&^|!\\~]=?`, chroma.Operator, nil},
			{`[()\[\]:;,<>/?{}]`, chroma.Punctuation, nil},
			{`(?=\w)`, chroma.Name, chroma.Push("name")},
		},
		"format": {
			{`\.\n`, chroma.LiteralStringInterpol, chroma.Pop(1)},
			{`[^\n]*\n`, chroma.LiteralStringInterpol, nil},
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
			{`[a-zA-Z_]\w*[!?]?`, chroma.NameFunction, nil},
			{`\s+`, chroma.Text, nil},
			{`(\([$@%]*\))(\s*)`, chroma.ByGroups(chroma.Punctuation, chroma.Text), nil},
			{`;`, chroma.Punctuation, chroma.Pop(1)},
			{`.*?\{`, chroma.Punctuation, chroma.Pop(1)},
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
