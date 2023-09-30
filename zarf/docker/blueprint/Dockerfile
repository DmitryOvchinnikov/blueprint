# Build the Go Binary.
FROM golang:1.21.1 as build_blu-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Create the service directory and the copy the module files first and then
# download the dependencies. If this doesn't change, we won't need to do this
# again in future builds.
# RUN mkdir /blueprint
# COPY go.* /blueprint/
# WORKDIR /blueprint
# RUN go mod download

# Copy the source code into the container.
COPY . /blueprint

# Build the admin binary.
WORKDIR /blueprint/app/tooling/blu-admin
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Build the service binary.
WORKDIR /blueprint/app/services/blu-api
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.18
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S blu && \
    adduser -u 1000 -h /blueprint -G blu -S blu
COPY --from=build_blu-api --chown=blu:blu /blueprint/zarf/keys/. /blueprint/zarf/keys/.
COPY --from=build_blu-api --chown=blu:blu /blueprint/app/tooling/blu-admin/blu-admin /blueprint/blu-admin
COPY --from=build_blu-api --chown=blu:blu /blueprint/app/services/blu-api/blu-api /blueprint/blu-api
WORKDIR /blueprint
USER blu
CMD ["./blu-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="blu-api" \
      org.opencontainers.image.authors="Dmitry Ovchinnikov <dmitry.v.ovchinnikov@icloud.com>" \
      org.opencontainers.image.source="https://github.com/dmitryovchinnikov/blueprint/tree/master/app/services/sales-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Mold Labs"