FROM golang:1.13-alpine
ENV CGO_ENABLED 0
ADD  . /go/src/scratch-container/
RUN go test ./... && go install ./...

FROM scratch
COPY --from=0 /go/bin/scratch-container /usr/local/bin/scratch-container
ENTRYPOINT ["scratch-container"]
