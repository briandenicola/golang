FROM golang:1.10 AS builder

WORKDIR $GOPATH/src/restapi 
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch
COPY --from=builder /main ./
EXPOSE 8000
ENTRYPOINT ["./main" ]