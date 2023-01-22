install_goyacc:
	go install golang.org/x/tools/cmd/goyacc@latest

yacc: install_goyacc
	goyacc -o parser/parser.go parser/parser.y

export_function:
	sed -i -e "s/yyParse/YyParse/g" `grep yyParse parser/ -rl`

build: yacc export_function