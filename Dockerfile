FROM golang:1.14.4-alpine3.12 as builder
WORKDIR /app
COPY . .
COPY /go/pkg /go/pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN ls -lrt

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app
CMD ["./main"]