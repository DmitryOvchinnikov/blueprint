# Build the Go Binary.
FROM golang:1.21.3 as build_mold-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Create the blueprint directory and the copy the module files first and then
# download the dependencies. If this doesn't change, we won't need to do this
# again in future builds.
# RUN mkdir /blueprint
# COPY go.* /blueprint/
# WORKDIR /blueprint
# RUN go mod download

# Copy the source code into the container.
COPY . /service

# Build the admin binary.
WORKDIR /service/app/tooling/mold-admin
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Build the blueprint binary.
WORKDIR /service/app/services/mold-api
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.18
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S mold && \
    adduser -u 1000 -h /blueprint -G mold -S mold
COPY --from=build_mold-api --chown=mold:mold /service/zarf/keys/. /service/zarf/keys/.
COPY --from=build_mold-api --chown=mold:mold /service/app/tooling/mold-admin/mold-admin /service/mold-admin
COPY --from=build_mold-api --chown=mold:mold /service/app/services/mold-api/mold-api /service/mold-api
WORKDIR /service
USER mold
CMD ["./mold-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="mold-api" \
      org.opencontainers.image.authors="Dmitry Ovchinnikov <dmitry@gmail.com>" \
      org.opencontainers.image.source="https://github.com/dmitryovchinnikov/blueprint/tree/master/app/services/mold-api" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Blueprint"