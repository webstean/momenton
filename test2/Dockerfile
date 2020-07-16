FROM golang:alpine as build-env

WORKDIR /app

ENV GIT_URL="https://github.com/webstean/momenton.git"
ENV GO111MODULE=off

RUN apk update && \
	apk upgrade && \
	apk --no-cache add git unzip ca-certificates

RUN git clone ${GIT_URL} && export GIT_COMMIT="$(git rev-list -1 HEAD)"

RUN cd /app/momenton/test2 && go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -o test2

FROM alpine
ENV GO111MODULE=off
WORKDIR /app
COPY --from=build-env /app/test2 /app
EXPOSE 3000
ENTRYPOINT ["/app/test2"]

