# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian13:nonroot@sha256:e2e927ec666bae08560abb3c55d0659eceabb657f56b6782ab500a9fc7f555e3

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
