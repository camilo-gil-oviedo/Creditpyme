FROM golang:1.25 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /creditpyme ./Creditpyme

FROM gcr.io/distroless/static-debian11
COPY --from=builder /creditpyme /creditpyme
EXPOSE 8080
ENTRYPOINT ["/creditpyme"]
