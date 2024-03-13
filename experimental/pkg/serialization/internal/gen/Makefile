.PHONY: count fmt run tag

count:
	countula -extensions .go -ignore-prefix "//" > lines

run:
	go run -tags=experimental .

fmt:
	go run -tags=experimental . -fmt 2> log

tag:
	go run -tags=experimental . -fmt -tag 2> log
