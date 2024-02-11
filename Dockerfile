FROM golang:1.18-alpine

ENV APP_DIR=/app/bitwyre/gateway/rest/p2p_api
ENV APP_BIN=app
ENV SRC_DIR=./gateway/rest/p2p_api
ENV SHARED_DST_DIR=/app/bitwyre/shared
ENV SHARED_SRC_DIR=./shared

WORKDIR ${APP_DIR}

COPY ${SHARED_SRC_DIR}/go-entity ${SHARED_DST_DIR}/go-entity
COPY ${SRC_DIR}/go.mod ${SRC_DIR}/go.sum ./
RUN go mod download

COPY ${SRC_DIR}/ ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go -o ${APP_DIR}/docs
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ${APP_DIR}/${APP_BIN} main.go
RUN chmod +x ${APP_DIR}/entrypoint.sh

RUN adduser -D -H bitwyre \
    && chown -R bitwyre:bitwyre ${APP_DIR}

USER bitwyre
EXPOSE 3051

ENTRYPOINT ["/app/bitwyre/gateway/rest/p2p_api/entrypoint.sh"]