FROM golang:alpine as build-env
ENV GO111MODULE=on
WORKDIR /app
ADD ./ /app
RUN cd /app && go build -o golang-test  .

FROM alpine
ENV GO111MODULE=on
WORKDIR /app
COPY --from=build-env /app/golang-test /app
RUN apk update --no-cache \
    apk add git
EXPOSE 8000
ENTRYPOINT ["/app/golang-test"]
