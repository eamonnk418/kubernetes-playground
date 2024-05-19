FROM golang:1.22.3-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o kubernetes-playground ./cmd/kubernetes-playground/

FROM ubuntu:24.04

COPY --from=builder /app/kubernetes-playground ./

CMD [ "/kubernetes-playground" ]
