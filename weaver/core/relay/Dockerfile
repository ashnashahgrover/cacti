FROM rust:1.69.0-slim-bullseye AS builder

RUN apt-get update

RUN apt-get install curl musl-tools perl build-essential -y

# adding the musl library target so that we can build
# the components for busybox/alpine environments
#
RUN rustup target add x86_64-unknown-linux-musl

# We add the rustfmt component. Ideally this should be having the same
# toolchain than the target we are building, but given that we are building
# on linux (debian), we can use the toolchain mentioned below also because
# custom toolchains are not supported.
# 
RUN rustup component add rustfmt --toolchain 1.69.0-x86_64-unknown-linux-gnu
RUN mkdir -p /opt/relay
COPY . /opt/relay

# Build OpenSSL

ENV SSL_VERSION=3.0.14
    # https://www.openssl.org/source/openssl-1.1.1w.tar.gz
RUN ln -s /usr/include/x86_64-linux-gnu/asm /usr/include/x86_64-linux-musl/asm && \
    ln -s /usr/include/asm-generic /usr/include/x86_64-linux-musl/asm-generic && \
    ln -s /usr/include/linux /usr/include/x86_64-linux-musl/linux
RUN mkdir /musl && \
    curl https://www.openssl.org/source/openssl-$SSL_VERSION.tar.gz -LO && \
    tar zxvf openssl-$SSL_VERSION.tar.gz && \
    cd openssl-$SSL_VERSION/ && \
    CC="musl-gcc -fPIE -pie" ./Configure no-shared no-async --prefix=/musl --openssldir=/musl/ssl linux-x86_64 && \
    make depend && make -j$(nproc) && make install
    
ENV PKG_CONFIG_ALLOW_CROSS=1 \
    OPENSSL_STATIC=true \
    OPENSSL_DIR=/musl

# We install missing dependencies within the
# the same layer where we do the compilation
# to then remove it from the image in the
# same docker command.
RUN cd /opt/relay && RUSTFLAGS=-Clinker="-Clinker=musl-gcc -Crelocation-model=static" cargo build --release --target x86_64-unknown-linux-musl



FROM busybox AS all

LABEL COMPONENT="relay all-in-one"
LABEL SOLUTION=dlt-interop

RUN addgroup -g 1000 relay
RUN adduser -D -s /bin/sh -u 1000 -G relay relay

RUN mkdir -p /opt/relay/config
RUN mkdir -p /opt/relay/driver

COPY --from=builder /opt/relay/target/x86_64-unknown-linux-musl/release/client /opt/relay/
COPY --from=builder /opt/relay/target/x86_64-unknown-linux-musl/release/client-tls /opt/relay/
COPY --from=builder /opt/relay/target/x86_64-unknown-linux-musl/release/dummy-driver /opt/relay/
COPY --from=builder /opt/relay/target/x86_64-unknown-linux-musl/release/server /opt/relay/
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

COPY docker/server.template.toml /opt/relay/config/
COPY docker/driver.template.toml /opt/relay/config/
COPY docker/init.sh /opt/relay/
COPY docker/entrypoint-server.sh /opt/relay/
COPY docker/entrypoint-client.sh /opt/relay/
COPY docker/entrypoint-driver.sh /opt/relay/
COPY docker/entrypoint.sh /opt/relay/
COPY driver/driver-error-constants.json /opt/relay/driver

COPY fingerprint.json /opt/relay/ 

RUN chmod +x /opt/relay/server /opt/relay/client /opt/relay/client-tls /opt/relay/dummy-driver /opt/relay/entrypoint.sh /opt/relay/entrypoint-client.sh /opt/relay/entrypoint-driver.sh /opt/relay/entrypoint-server.sh
RUN chown -R relay:relay /opt/relay

USER relay

WORKDIR /opt/relay
ENTRYPOINT [ "./entrypoint.sh" ]

# These labels will be changing at every build
# therefore we leave them at the end in order
# to minimise the amount of layers that are
# built every time.
ARG COMMIT
ARG BRANCH
ARG VERSION
ARG PROTOS_VERSION
LABEL COMMIT=${COMMIT} BRANCH=${BRANCH} VERSION=${VERSION} ROTOS_VERSION=${PROTOS_VERSION}
