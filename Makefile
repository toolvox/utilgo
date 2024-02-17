.SILENT:
.PHONY: godoc cover count

godoc:
	start "http://localhost:6060" ; \
	godoc -http=:6060 -play -index 

cover:
	go test ./... -coverprofile=testing/cover.out

count: 
	countula -gitignore -excludes "zz_" -extensions "go" > lines