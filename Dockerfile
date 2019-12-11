FROM liudongkenny/rpc:v3
RUN go get gopkg.in/gomail.v2
ADD . /go/src/dev-framework-go/
WORKDIR /go/src/dev-framework-go/
#EXPOSE 8890
#CMD ["go run", "main.go"]
