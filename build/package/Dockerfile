# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian12:nonroot

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
