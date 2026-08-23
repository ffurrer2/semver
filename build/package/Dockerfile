# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian13:nonroot@sha256:1c2c046bc09ed40fad370b599a0b1ae7987f55b01e247cf27a7c27cd97e5bbc7

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
