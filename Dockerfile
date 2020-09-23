FROM golang:1.14.3
ENV GOPROXY "https://goproxy.cn"
WORKDIR $GOPATH/src/github.com/asynccnu/grade_service_v2
COPY . $GOPATH/src/github.com/asynccnu/grade_service_v2
RUN go build -o main -v .
EXPOSE 8081
CMD ["./main"]
