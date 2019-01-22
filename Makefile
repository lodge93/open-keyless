.PHONY: mocks

$GOPATH/bin/mockgen:
	 GO111MODULE=off go get github.com/golang/mock/gomock
	 GO111MODULE=off go install github.com/golang/mock/mockgen

$GOPATH/src/periph.io/x/periph/conn/gpio:
	 GO111MODULE=off go get periph.io/x/periph/conn/gpio

tools: $GOPATH/bin/mockgen

mocks: tools $GOPATH/src/periph.io/x/periph/conn/gpio
	 mockgen -source $(GOPATH)/src/periph.io/x/periph/conn/gpio/gpio.go -destination internal/mocks/pin_io.go -package mocks

test:
	go test -v -cover ./...
