# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.20-alpine3.18 as builder

RUN apk add --no-cache gcc musl-dev binutils-gold linux-headers git

# Configure creds for gitea
ARG GIT_VDBTO_TOKEN

# Get dependencies - will also be cached if we won't change go.mod/go.sum
WORKDIR /go-ethereum/
COPY go.mod .
COPY go.sum .
COPY wrapmain wrapmain
RUN if [ -n "$GIT_VDBTO_TOKEN" ]; then git config --global url."https://$GIT_VDBTO_TOKEN:@git.vdb.to/".insteadOf "https://git.vdb.to/"; fi && \
    cd wrapmain && \
    go mod download && \
    rm -f ~/.gitconfig

COPY . .
RUN cd wrapmain && go build --trimpath -o /go-ethereum/build/bin/geth .

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.18

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
