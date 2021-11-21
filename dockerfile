############################
# STEP 1 build executable binary
############################
FROM golang AS builder
WORKDIR $GOPATH/src/github.com/Jimbo4794/galasa-kubernetes-operator
COPY . /go/src/github.com/Jimbo4794/galasa-kubernetes-operator

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=auto
ENV GOPATH=/go

RUN go build -o /go/bin/operator /go/src/github.com/Jimbo4794/galasa-kubernetes-operator/cmd/operator/main.go

############################
# STEP 2 build a small image
############################
FROM alpine
COPY --from=builder /go/bin/operator /go/bin/operator
ENTRYPOINT ["/go/bin/operator"]