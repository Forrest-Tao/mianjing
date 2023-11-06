BINARY="BIANRYFileName"

all: goTool build

goTool:
	go fmt ./
	go vet ./

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/${BIANRYFileName}
	#go build -o bin/main main.go

run:
	go run main.go

clean:
	@if [ -f ${BINARY} ]; then
	    rm ${BINARY}
	fi

compile:
	echo "Compiling for every OS and platform"
	#GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	#GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

