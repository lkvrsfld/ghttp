FROM golang:1.24-alpine AS builder

WORKDIR /src

ARG CGO_ENABLED=0
ARG GOOS=linux

ENV CGO_ENABLED=${CGO_ENABLED} \
    GOOS=${GOOS}

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build \
    -trimpath \
    -ldflags="-s -w" \
    -gcflags="all=-l -B" \
    -o /app/ghttp .

FROM scratch

COPY --from=builder /app/ghttp /bin/ghttp

ENTRYPOINT ["/bin/ghttp"]