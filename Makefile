START_FLAG = "/Users/oscar/Library/Mobile Documents/com~apple~CloudDocs/Documents/Gym Partner/gym-partner-env/api/"

run:
	go run main.go -start=$(START_FLAG)

build:
	GOOS=linux GOARCH=amd64 go build -o gym-partner-api main.go

test-all:
	go test ./usecases/interactor/test -v
	go test ./interfaces/repository/test -v
	go test ./interfaces/controller/test -v