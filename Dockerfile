FROM golang:1.24.4 AS builder
WORKDIR /app
RUN go install github.com/air-verse/air@latest
# COPY go.mod go.sum ./
# RUN go mod download

FROM golang:1.24.4
WORKDIR /app
COPY --from=builder /go/bin/air /go/bin/air
ENV PATH=/go/bin:$PATH
COPY . . 
CMD [ "air" ]