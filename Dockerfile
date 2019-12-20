FROM liudongkenny/rpc:v5
ADD . /go/src/dev-framework-go/
WORKDIR /go/src/dev-framework-go/
#EXPOSE 8890
#CMD ["go run", "main.go"]
