FROM alpine:3.14.2

ENV THRIFT_VERSION=0.14.2
RUN apk  add --update --no-cache python2 python3 zlib-dev bash \
 libtool automake autoconf make gcc g++ alpine-sdk

WORKDIR /tmp

RUN wget https://dlcdn.apache.org/thrift/$THRIFT_VERSION/thrift-$THRIFT_VERSION.tar.gz \
    && tar xvzf thrift-$THRIFT_VERSION.tar.gz \
    && cd thrift-$THRIFT_VERSION/ \
    && ./bootstrap.sh \
    && ./configure \
    && make && make install

# /usr/local/bin/thrift
#ENTRYPOINT sleep infinity
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"

