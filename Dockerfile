FROM alpine:latest

LABEL VENDOR="David Mendoza"
LABEL MAINTAINER="David Mendoza <jdmendozar@gmail.com>"
LABEL DESCRIPTION="Todos App built with Go/VueJS"
LABEL VERSION="1.0.0"

RUN apk update && apk upgrade \
  && apk add ca-certificates \
  && rm -rf /var/cache/apk/*

RUN apk add --update tzdata
ENV TZ=America/Chicago
RUN rm -rf /var/cache/apk/*

COPY todose /app
COPY settings.yml /
COPY web/dist /web/dist

ENTRYPOINT [ "/app" ]