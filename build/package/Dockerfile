# SPDX-License-Identifier: MIT
FROM gcr.io/distroless/static-debian11:nonroot

COPY --chown=nonroot:nonroot semver /usr/local/bin/semver

USER nonroot

ENTRYPOINT ["/usr/local/bin/semver"]
CMD ["help"]
