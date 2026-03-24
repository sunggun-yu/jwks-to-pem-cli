FROM alpine:3.23.3
COPY jwks-to-pem /usr/local/bin/jwks-to-pem
ENTRYPOINT ["/usr/local/bin/jwks-to-pem"]
