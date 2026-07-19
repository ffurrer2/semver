# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian13:nonroot@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6

COPY semver /usr/local/bin/semver

USER nonroot
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
