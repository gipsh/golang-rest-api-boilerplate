# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .

RUN apk add make
RUN apk add bash
RUN apk add git
#RUN go get -u github.com/swaggo/swag/cmd/swag

#RUN make doc
RUN make build

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/test-svc .
COPY app.env .
COPY db ./db


EXPOSE 8080
CMD [ "/app/test-svc" ]
ENTRYPOINT [ "/app/test-svc" ]