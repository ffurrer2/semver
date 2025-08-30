# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian12:nonroot@sha256:a9f88e0d99c1ceedbce565fad7d3f96744d15e6919c19c7dafe84a6dd9a80c61

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
