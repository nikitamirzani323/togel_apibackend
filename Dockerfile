FROM golang:alpine AS agenapibuilds

WORKDIR /appbuilds

COPY . .

RUN go mod tidy
RUN go build -o binary


FROM alpine:latest as agenapirelease
WORKDIR /app
RUN apk add tzdata
COPY --from=agenapibuilds /appbuilds/binary .
COPY --from=agenapibuilds /appbuilds/.env /app/.env
ENV PORT=7072
ENV DB_USER="sperma"
ENV DB_PASS="asdQWE123!@#"
ENV DB_HOST="165.22.242.64"
ENV DB_PORT="3306"
ENV DB_NAME="db_tot"
ENV DB_DRIVER="mysql"
ENV DB_REDIS_HOST="128.199.124.131:6379"
ENV DB_REDIS_PASSWORD="asdQWE123!@#"
ENV DB_REDIS_NAME="0"
ENV JWT_SECRET_KEY="secrettotobackend"
ENV JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=1440
ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]