FROM golang:alpine AS build

# Set current working dir
WORKDIR /tmp/hello-world

COPY . .

RUN go mod download

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build Go app
RUN go build -o hello .

FROM alpine

#RUN mkdir /app

COPY --from=build /tmp/hello-world/hello /app/hello

CMD ["app/hello"]