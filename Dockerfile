FROM golang:1.21-alpine3.20 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=build /app/main .
CMD ["/app/main"]