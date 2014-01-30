SRCS = \
	main.icn \
	emit.icn \
	trans.icn \
	parse.icn

parser.go: main oberon0.ebnf
	./main < oberon0.ebnf | gofmt > parser.go

main: $(SRCS)
	icont -u -o main $(SRCS)
