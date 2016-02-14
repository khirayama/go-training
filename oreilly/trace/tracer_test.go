package trace

import (
	"bytes"
	"testing"
)

// filenameは*_test.go
// 名前がTestから始まり、*testing.T型の引数を1つ受け取る関数はユニットテストとみなされる
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Newからの戻り値がnilです")
	} else {
		tracer.Trace("こんちには、traceパッケージ")
		if buf.String() != "こんちには、traceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}
