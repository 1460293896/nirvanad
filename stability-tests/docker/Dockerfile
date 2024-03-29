ARG NIRVANAD_IMAGE
ARG NIRVANAMINER_IMAGE

FROM ${NIRVANAD_IMAGE} as nirvanad
FROM ${NIRVANAMINER_IMAGE} as nirvanaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/Nirvana-Chain/nirvanad/

WORKDIR /go/src/github.com/Nirvana-Chain/nirvanad/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=nirvanad /app/ /app/
COPY --from=nirvanaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Nirvana-Chain/nirvanad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
