run:
	go run cmd/server/main.go

test:
	go test -v -cover ./...

testscover:
	go test -coverprofile testdata/tests_cover.out ./...

totalcover:
	go test -coverprofile testdata/tests_cover.out ./...
	go tool cover -func testdata/tests_cover.out

