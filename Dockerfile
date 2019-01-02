FROM scratch
COPY ca-certificates.crt /etc/ssl/certs/
ADD blocktor /
ENTRYPOINT ["/blocktor"]
