FROM golang:1.14.4-alpine3.12 as builder
ENV DB_USERNAME="REPLACE_ME"
ENV DB_PASSWORD="REPLACE_ME"
ENV DB_URL="172.17.0.20:3306"
ENV BASE_FOLDER="/mnt/videos/auto-porter"
ENV TMDB_API_KEY="REPLACE_ME"

WORKDIR /app
COPY . .
COPY /go/pkg /go/pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN ls -lrt

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app
CMD ["./main"]