install:
	go install
test:
	go test ./git ./tree -coverprofile cover.out
