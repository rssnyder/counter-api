FROM scratch

LABEL org.opencontainers.image.source https://github.com/rssnyder/counter-api

COPY counter-api /

ENTRYPOINT ["/counter-api"]
