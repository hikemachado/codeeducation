FROM golang as builder

RUN sh -c "go get github.com/hikemachado/codeeducation/"

FROM scratch
COPY --from=builder /go/bin/codeeducation .
#COPY ./codeeducation .
ENTRYPOINT ["/codeeducation"]






