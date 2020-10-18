FROM golang:1.14-alpine
RUN apk --no-cache add ca-certificates
RUN apk add --update make
RUN apk add --no-cache bash
RUN apk add build-base

WORKDIR /go/src/copyrighted-phrases
COPY . .
RUN go get ./...
RUN make install
RUN /go/src/copyrighted-phrases/init.sh

COPY docker/* /opt/
RUN chmod +x /opt/*.sh

WORKDIR /opt

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

# CMD ["/usr/bin/copyrightedphrasesd version"]
CMD ["/opt/run_all.sh"]