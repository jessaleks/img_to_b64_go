FROM golang:bullseye

ENV GOROOT=/usr/local/go \
    GOPATH=/go
ENV PATH=${GOPATH}/bin:${GOROOT}/bin:${PATH}
COPY library-scripts/go-debian.sh /tmp/library-scripts/
RUN apt-get update && bash /tmp/library-scripts/go-debian.sh "latest" "${GOROOT}" "${GOPATH}"
RUN apt-get update && apt-get install libwebp6
RUN apt-get update && apt-get install zsh
