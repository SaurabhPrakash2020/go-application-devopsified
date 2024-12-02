FROM golang:1.22 AS stage-1
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main .


FROM gcr.io/distroless/base
COPY --from=stage-1 /app/main .
COPY --from=stage-1 /app/static ./static
EXPOSE 8080
CMD ["./main"]