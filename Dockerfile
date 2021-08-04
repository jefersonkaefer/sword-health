# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.15-alpine3.12

RUN apk update \
    && apk add curl git protoc

ARG GIT_USERNAME
ARG GIT_TOKEN

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air \
    && GO111MODULE=on go get golang.org/x/tools/gopls@latest

# Install other tools.
RUN go get -u -v \
    github.com/mdempsky/gocode \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    golang.org/x/tools/cmd/guru \
    golang.org/x/tools/cmd/gorename \
    github.com/go-delve/delve/cmd/dlv \
    github.com/stamblerre/gocode \
    github.com/rogpeppe/godef \
    golang.org/x/tools/cmd/goimports \
    golang.org/x/lint/golint \
    github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc \ 
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
COPY docker-entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

WORKDIR /go/src 


ENTRYPOINT ["docker-entrypoint.sh"]

CMD [ "air" ]