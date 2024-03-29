# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/Nirvana-Chain/nirvanad/

WORKDIR /go/src/github.com/Nirvana-Chain/nirvanad/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache nirvanad dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o nirvanad .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/Nirvana-Chain/nirvanad/nirvanad /app/
COPY --from=build /go/src/github.com/Nirvana-Chain/nirvanad/infrastructure/config/sample-nirvanad.conf /app/

USER root
ENTRYPOINT [ "/sbin/tini", "--" ]
