# use the yacc .y file to generate parser
go tool yacc -o parser/parser.go parser/parser.y

# change first letter to upper case for export the function
sed -i '' "s/yyParse/YyParse/g" `grep yyParse parser/ -rl`
