FROM gcr.io/repcore-prod/protoc-go

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin:/usr/local/protoc/bin

RUN apt-get update
RUN apt-get install -y git
RUN go get github.com/google/go-genproto/...

COPY . /go/src/github.com/vendasta/mscli
RUN go build -o /usr/bin/protoc-gen-sdkgen github.com/vendasta/mscli/tools/sdkgen

CMD ["bash"]

VOLUME ["/src"]
