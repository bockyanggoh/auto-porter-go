FROM golang:1.14.5-alpine3.12 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN ls -lrt

FROM scratch
WORKDIR /app
ENV DB_USERNAME="REPLACE_ME"
ENV DB_PASSWORD="REPLACE_ME"
ENV DB_URL="172.17.0.20:3306"
ENV BASE_FOLDER="auto-porter"
ENV TMDB_API_KEY="REPLACE_ME"
ENV PROJECT_RELATIVE_FOLDER="REPLACE_ME"
COPY --from=builder /app/main /app
CMD ["./main"]