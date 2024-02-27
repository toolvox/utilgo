.SILENT:
.PHONY: godoc diff dump cover count install install.codump

godoc:
	start "http://localhost:6060" ; \
	godoc -http=:6060 -play -index 

COVERFILE = testing/cover.out

diff:
	(git diff --cached | tee diff) | clip ; \
	echo "diff saved to 'diff' file and clipboard"

dump:
	codump -path . -out dump -exts .go

cover:
	go test ./... -coverprofile=testing/cover.out ; \
	(head -n 1 $(COVERFILE) && tail -n +2 $(COVERFILE) | sort -V) > $(COVERFILE).tmp && mv $(COVERFILE).tmp $(COVERFILE)

count: 
	countula -gitignore -excludes "zz_" -extensions "go" > lines

install.codump:
	cd cmd/codump ; \
	go install

install: install.codump