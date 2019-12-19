FROM liudongkenny/rpc:v4
ADD . /go/src/dev-framework-go/
WORKDIR /go/src/dev-framework-go/
#EXPOSE 8890
#CMD ["go run", "main.go"]
