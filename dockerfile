FROM alpine:latest  as builder

RUN apk add  --no-cache go

COPY . ./source

WORKDIR /source

RUN go build -o /app

FROM alpine:latest

COPY --from=builder /app /home/daemon/app

USER daemon

EXPOSE 2255

ENTRYPOINT ["/home/daemon/app"]
