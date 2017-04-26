# go get golang.org/x/tools/cmd/goyacc

# use the yacc .y file to generate parser
goyacc -o parser/parser.go parser/parser.y

# change first letter to upper case for export the function
sed -i -e "s/yyParse/YyParse/g" `grep yyParse parser/ -rl`
