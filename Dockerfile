FROM alpine:3.14

COPY build/prom-intersight-metrics-linux_amd64 /prom-intersight-metrics

ENTRYPOINT [ "/prom-intersight-metrics" ]