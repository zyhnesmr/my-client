FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN echo "machine github.com" >> ~/.netrc \
    && echo "    login zyhnesmr" >> ~/.netrc \
    && echo "    password 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" >> ~/.netrc \
    && chmod 600 ~/.netrc
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache git

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
#COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/mydemo my.go


FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/mydemo /app/mydemo
#COPY --from=builder /app/etc /app/etc

EXPOSE 8888

CMD ["./mydemo"]
