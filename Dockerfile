FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o cloud06 .

FROM golang:1.20-alpine 

COPY --from=builder /app/cloud06 /app/cloud06

EXPOSE 8080
CMD ["/app/cloud06"]