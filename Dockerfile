FROM registry.access.redhat.com/rhel7:latest
COPY go-http-metrics /usr/local/bin
ENTRYPOINT ["go-http-metrics"]
