make: fmt lint test build-alphabet

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go clean --testcache
	go test ./...

build-alphabet:
	go build -o ./main ./cmd/main.go
	docker build -t bentangled/alphabet-api/alphabet .

deploy-alphabet:
	kubectl apply -f ./alphabet.yaml

delete-alphabet:
	kubectl delete -f ./alphabet.yaml