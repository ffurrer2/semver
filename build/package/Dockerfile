# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian13:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
