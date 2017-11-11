FROM scratch

LABEL maintainer="JorritSalverda" \
      description="The nike-plus-to-runkeeper-sync application exports runs from Nike+ and imports them into RunKeeper"

COPY ca-certificates.crt /etc/ssl/certs/
COPY nike-plus-to-runkeeper-sync /

ENTRYPOINT ["/nike-plus-to-runkeeper-sync"]