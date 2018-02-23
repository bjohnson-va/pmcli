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

ENV GOLANG_VERSION 1.9.1
ENV GOLANG_DOWNLOAD_URL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN curl "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

ENV CLOUDSDK_CORE_DISABLE_PROMPTS 1
ENV PATH /opt/google-cloud-sdk/bin:$PATH

RUN apt-get update -y
RUN apt-get install -y jq
RUN curl https://sdk.cloud.google.com | bash
RUN mv /root/google-cloud-sdk /opt
RUN gcloud components install kubectl
# Temporarily pin gcloud to 183
RUN gcloud components update --version=183.0.0

RUN go get -u github.com/golang/lint/golint
RUN go get -u github.com/pierrre/gotestcover
RUN go get -u github.com/jstemmer/go-junit-report

COPY . /go/src/github.com/vendasta/mscli
RUN go install github.com/vendasta/mscli
RUN rm -rf /go/src/github.com/vendasta/mscli

CMD ["/bin/bash"]
