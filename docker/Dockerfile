FROM golang:1.20

WORKDIR /usr/src/app

RUN apt-get -q update
RUN apt install -y protobuf-compiler

RUN git config --global --add safe.directory /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod  ./
RUN go mod download && go mod verify
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install mvdan.cc/gofumpt@latest
RUN go install golang.org/x/tools/cmd/godoc@latest
RUN go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN export PATH="$PATH:/usr/src/app"

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server
RUN go build -v -o /usr/local/bin/app ./cmd/agent