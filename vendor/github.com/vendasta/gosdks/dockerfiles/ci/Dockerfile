FROM openjdk:8-jdk

# Install golang
# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.9
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 d70eadefce8e160638a9a6db97f7192d8463069ab33138893ad3bf31b0650a79

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN go get -u github.com/golang/lint/golint

RUN \
  apt-get update && \
  apt-get install -y python python-dev python-pip python-virtualenv git unzip && \
  rm -rf /var/lib/apt/lists/*
RUN pip install invoke==0.15.0
RUN mkdir -p /tmp/bin && \
    cd /tmp/bin && \
    wget https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip && \
    unzip protoc-3.2.0-linux-x86_64.zip && \
    cp -r . /proto && \
    cp -v ./bin/protoc /usr/bin/ && \
    git clone https://github.com/googleapis/googleapis.git && \
    cp -r googleapis/google/api /proto/include/google && \
    cd / && \
    rm -rfv /tmp/bin
RUN mkdir -p /tm

CMD ["/bin/bash"]