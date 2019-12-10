FROM liudongkenny/rpc:v3
ADD . /go/src/dev-framework-go/
WORKDIR /go/src/dev-framework-go/
CMD ["go run", "main.go"]
