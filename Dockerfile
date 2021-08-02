##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy other files and build binary
COPY *.go ./

# Run binary
RUN go build -o /ose

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ose /ose

EXPOSE 8888

USER nonroot:nonroot

ENTRYPOINT ["/ose"]
