package internal

import "testing"

func TestTraceIDGenerator(t *testing.T) {
	tg := NewTraceIDGenerator(12345)
	traceID := tg.GenerateTraceID()
	if traceID != "1ae969564b34a33ecd1af05fe6923d6d" {
		t.Error(traceID)
	}
	spanID := tg.GenerateSpanID()
	if spanID != "e71870997d38ef60" {
		t.Error(spanID)
	}
	if p := tg.GeneratePriority(); p.Float32() != 0.05700199 {
		t.Error(p)
	}
}

func BenchmarkTraceIDGenerator(b *testing.B) {
	tg := NewTraceIDGenerator(12345)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if id := tg.GenerateSpanID(); id == "" {
			b.Fatal(id)
		}
	}
}
