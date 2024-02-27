.SILENT:
.PHONY: godoc cover count install install.codump

godoc:
	start "http://localhost:6060" ; \
	godoc -http=:6060 -play -index 

COVERFILE = testing/cover.out

cover:
	go test ./... -coverprofile=testing/cover.out ; \
	(head -n 1 $(COVERFILE) && tail -n +2 $(COVERFILE) | sort -V) > $(COVERFILE).tmp && mv $(COVERFILE).tmp $(COVERFILE)

count: 
	countula -gitignore -excludes "zz_" -extensions "go" > lines

install.codump:
	cd cmd/codump ; \
	go install

install: install.codump