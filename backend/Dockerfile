FROM golang:1.21.3-alpine3.18 as builder

WORKDIR /app

COPY . .

RUN go build -o backend.exe

FROM alpine:3.18

WORKDIR /app

COPY --from=builder app/backend.exe .

ENV GIN_MODE=release

CMD [ "./backend.exe" ]