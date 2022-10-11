# first stage - builds the binary from sources
FROM golang:alpine as build

# using build as current directory
WORKDIR /app

# adding the source code to current dir:
COPY . .

# downloading dependencies and verifying
RUN go mod download && go mod verify


# building the project
RUN go build  -o client .

# second stage - using minimal image to run the server
FROM alpine:latest

# using /app as current directory
WORKDIR /app

# copy server binary from `build` layer
COPY --from=build /app/client .
COPY --from=build /app/jsonConfig ./jsonConfig

# binary to run
CMD "/app/client"
