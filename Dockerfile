FROM golang:latest 
WORKDIR $GOPATH/src/github.com/asynccnu/grade_service_v2
COPY . $GOPATH/src/github.com/asynccnu/grade_service_v2
RUN go build -o main . 
EXPOSE 8080
CMD ["./main"]