FROM golang as builder
COPY ./greetings.go ./greetings.go

CMD ["sh", "-c", "go build ./greetings.go"]

FROM scratch
COPY --from=builder /go/bin/greetings /greeetings
ENTRYPOINT ["/greetings"]



docker run -v ${pwd}:/usr/src/code -w /usr/src/code  golang go build -v



