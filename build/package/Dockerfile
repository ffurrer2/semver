# SPDX-License-Identifier: MIT
FROM scratch

COPY semver /

ARG BUILD_DATE
ARG GIT_SHA
ARG BUILD_VERSION

LABEL org.opencontainers.image.authors="Felix Furrer" \
      org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.description="A semantic versioning command line utility written in Go." \
      org.opencontainers.image.documentation="https://github.com/ffurrer2/semver/blob/main/README.md" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.revision="${GIT_SHA}" \
      org.opencontainers.image.source="https://github.com/ffurrer2/semver.git" \
      org.opencontainers.image.title="SemVer" \
      org.opencontainers.image.url="https://github.com/ffurrer2/semver" \
      org.opencontainers.image.vendor="Felix Furrer" \
      org.opencontainers.image.version="${BUILD_VERSION}"

ENTRYPOINT ["/semver"]
CMD ["help"]
