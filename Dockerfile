FROM golang:1.22.3-alpine3.20 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=build /app/main .
CMD ["/app/main"]