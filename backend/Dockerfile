FROM golang:1.25 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /creditpymesbackend .

FROM gcr.io/distroless/static-debian11
COPY --from=builder /creditpymesbackend /creditpymesbackend
EXPOSE 8080
ENTRYPOINT ["/creditpymesbackend"]
