# amolecoin build
# reference https://github.com/amolecoin/amolecoin
ARG IMAGE_FROM=busybox
FROM golang:1.14-stretch AS build
ARG ARCH=amd64
ARG GOARM
ARG AMOLECOIN_VERSION
ARG SCOMMIT
ARG SBRANCH
ARG STAG

ADD cmd $GOPATH/src/github.com/amolecoin/amolecoin/cmd
ADD src $GOPATH/src/github.com/amolecoin/amolecoin/src
ADD vendor $GOPATH/src/github.com/amolecoin/amolecoin/vendor
ADD template $GOPATH/src/github.com/amolecoin/amolecoin/template
ADD fiber.toml $GOPATH/src/github.com/amolecoin/amolecoin/fiber.toml

# This code checks if AMOLECOIN_VERSION is set and checkouts to that version if
# so. The git stash line prevents the build to fail if there are any uncommited
# changes in the working copy. It won't affect the host working copy.
RUN sh -c \
    'if test ${AMOLECOIN_VERSION};then \
        echo "Revision is set to: "${AMOLECOIN_VERSION}; \
        cd $GOPATH/src/github.com/amolecoin/amolecoin; \
        git stash; \
        git checkout ${AMOLECOIN_VERSION}; \
     fi'

ENV GOARCH="$ARCH" \
    GOARM="$GOARM" \
    CGO_ENABLED="0" \
    GOOS="linux" \
    GOLDFLAGS="-X main.Commit=${SCOMMIT} -X main.Branch=${SBRANCH}"

RUN cd $GOPATH/src/github.com/amolecoin/amolecoin && \
    echo "Building with GOLDFLAGS=$GOLDFLAGS GOARCH=$GOARCH GOARM=$GOARM CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS " && \
    go install -ldflags "${GOLDFLAGS}" ./cmd/... && \
    sh -c "if test -d $GOPATH/bin/linux_arm ; then mv $GOPATH/bin/linux_arm/* $GOPATH/bin/; fi; \
           if test -d $GOPATH/bin/linux_arm64 ; then mv $GOPATH/bin/linux_arm64/* $GOPATH/bin/; fi"

RUN apt-get update && \
    apt-get install -y ca-certificates


RUN /bin/bash -c 'mkdir -p /tmp/files/{usr/bin,/usr/local/amolecoin/src/gui/static,/usr/local/bin/,/etc/ssl}'
RUN cp -r /go/bin/* /tmp/files/usr/bin/
RUN cp -r  /go/src/github.com/amolecoin/amolecoin/src/gui/static /tmp/files/usr/local/amolecoin/src/gui/
RUN cp -r  /etc/ssl/certs /tmp/files/etc/ssl/certs
COPY docker_launcher.sh /tmp/files/usr/local/bin/docker_launcher.sh

# amolecoin image
FROM $IMAGE_FROM
ARG BDATE
ARG SCOMMIT
ARG SBRANCH
ARG STAG

# Image labels
LABEL "org.label-schema.name"="Amolecoin" \
      "org.label-schema.description"="Amolecoin core docker image" \
      "org.label-schema.vcs-url"="https://github.com/amolecoin/amolecoin/tree/develop/docker/images/mainnet" \
      "org.label-schema.vendor"="Amolecoin project" \
      "org.label-schema.url"="amolecoin.com" \
      "org.label-schema.schema-version"="1.0" \
      "org.label-schema.build-date"=$BDATE \
      "org.label-schema.vcs-ref"=$SCOMMIT \
      "org.label-schema.version"=$STAG \
      "org.label-schema.usage"="https://github.com/amolecoin/amolecoin/blob/"$SCOMMIT"/docker/images/mainnet/README.md" \
      "org.label-schema.docker.cmd"="docker volume create amolecoin-data; docker volume create amolecoin-wallet; docker run -d -v amolecoin-data:/data/.amolecoin -v amolecoin-wallet:/wallet -p 9982:9982 -p 9981:9981 --name amolecoin-node-stable amolecoin/amolecoin"

ENV COIN="amolecoin"
ENV RPC_ADDR="http://0.0.0.0:9981" \
    DATA_DIR="/data/.$COIN" \
    WALLET_DIR="/wallet" \
    WALLET_NAME="$COIN_cli.wlt"

# copy all the binaries
COPY --from=build /tmp/files /

# volumes
VOLUME $WALLET_DIR
VOLUME $DATA_DIR

EXPOSE 9982 9981

ENTRYPOINT ["docker_launcher.sh", "--web-interface-addr=0.0.0.0", "--gui-dir=/usr/local/amolecoin/src/gui/static"]
