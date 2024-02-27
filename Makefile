t="coverage.txt"

test:
	go test ./...

coverage:
	go test -coverprofile=$t ./... && go tool cover -html=$t && unlink $t

