START_FLAG = /Users/oscar/Documents/gym-partner-env/

run:
	go run main.go -start=$(START_FLAG)

build:
	GOOS=linux GOARCH=amd64 go build -o gym-partner-api main.go

test-all:
	go test ./usecases/interactor/test -v