# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian12:nonroot@sha256:e8a4044e0b4ae4257efa45fc026c0bc30ad320d43bd4c1a7d5271bd241e386d0

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
