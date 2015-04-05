
all: agent boss proxy

prep: 
	rm -rf $GOPATH/pkg
	mkdir -p bin
	go get ./...

agent: prep
	go build -o=./bin/agent agent/main.go

boss: prep
	go build -o=./bin/boss boss/main.go

proxy: prep
	go build -o=./bin/proxy proxy/main.go


test: all
	go test -v ./...
	bin/proxy &
	bin/agent &
	bin/agent &
	bin/agent &
	sleep 1
	bin/boss
	sleep 2
	killall proxy
	killall agent
	echo Success

clean:
	rm -rf ./bin/
