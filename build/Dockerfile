
# Description: Dockerfile for building the application

# Build arguments
ARG GO_VERSION=1.24


###################################################################
# Stage 1: Build the application                                  #
###################################################################
FROM golang:${GO_VERSION} AS build
WORKDIR /go/src

# Copy source code
COPY . .

# Install dependencies
RUN go mod download

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/server


###################################################################
# Stage 2: Create the runtime image                               #
###################################################################
FROM alpine:3.21 AS runtime
WORKDIR /

ENV USER=nonroot
ENV HOME=/home/$USER

# Create non-root user
RUN addgroup -S $USER \
  && adduser -S $USER -G $USER

# Copy application binary
WORKDIR $HOME
COPY --from=build /go/bin/app ./
RUN chown -R $USER:$USER $HOME

# Runtime configuration
USER $USER
EXPOSE 8080
CMD [ "./app" ]
