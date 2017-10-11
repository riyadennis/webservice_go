GOFILES= $$(go list -f '{{join .GoFiles " "}}')

clean:
	rm -rf vendor/

deps:
	glide install -v

test: deps
	go test -timeout=5s -cover -race $$(glide novendor)

ci_test: test

build: deps
	go build -o $(GOPATH)/bin/api_issa_products $$(go list -f '{{join .GoFiles " "}}')

run:
	go run $(GOFILES) server