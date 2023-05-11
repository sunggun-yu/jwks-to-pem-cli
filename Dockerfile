FROM alpine
COPY jwks-to-pem /usr/local/bin/jwks-to-pem
ENTRYPOINT ["/usr/local/bin/jwks-to-pem"]
