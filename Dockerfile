FROM    alpine:latest as resources

WORKDIR /usr/share/zoneinfo

RUN     apk add -U --no-cache tzdata zip ca-certificates

FROM scratch

COPY ./clash /clash
COPY    --from=resources /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/clash"]
