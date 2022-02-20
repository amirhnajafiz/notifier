package telemetry

import "io"

func NewExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace
}
