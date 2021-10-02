# BUILD
FROM golang:latest as build

WORKDIR /app
COPY . /app

RUN cd /app && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service .

# TEST
FROM build as test

RUN curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash
ENV PATH $PATH:/app/pact/bin
RUN go test

# PRODUCTION
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates
COPY --from=build /http-service ./
RUN chmod +x ./http-service


EXPOSE 4000
CMD ["./http-service"]