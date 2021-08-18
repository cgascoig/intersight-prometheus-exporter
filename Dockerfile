FROM alpine:3.14

COPY build/intersight-prometheus-exporter-linux_amd64 /intersight-prometheus-exporter

ENTRYPOINT [ "/intersight-prometheus-exporter" ]