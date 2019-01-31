install:
	go install
test:
	go test -race -v ./git ./db -coverprofile cover.out
