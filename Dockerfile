FROM ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib:0.58.0

COPY /generated/otelcol-lokiexp-debug /otelcol-contrib
