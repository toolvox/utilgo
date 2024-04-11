.SILENT:
.PHONY: default godoc diff dump cover count install

TEST=./test
COVERFILE = $(TEST)/cover.out

default: cover count
	
godoc:
	start "http://localhost:6060" ; \
	godoc -http=:6060 -play -index 

diff:
	(git diff --cached | tee diff) | clip ; \
	echo "diff saved to 'diff' file and clipboard"

dump:
	codump -root . -out dump -include "*.go"

cover:
	go test -cover ./... -coverprofile=$(COVERFILE) ; \
	(head -n 1 $(COVERFILE) && tail -n +2 $(COVERFILE) | sort -V) > $(COVERFILE).tmp && mv $(COVERFILE).tmp $(COVERFILE)

count: 
	countula -ignore ".gitignore" -exclude "**zz_**,.git,*.sum,*.out,*.mod,test/cover.out,LICENSE,lines" -ignore-prefix "//" > lines

install.codump:
	cd cmd/codump ; \
	go install

install.countula:
	cd cmd/countula ; \
	go install
