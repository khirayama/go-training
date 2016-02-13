package trace

import (
	"testing"
)

// filenameは*_test.go
// 名前がTestから始まり、*testing.T型の引数を1つ受け取る関数はユニットテストとみなされる
func TestNew(t *testing.T) {
	t.Error("まだテストを作成していません。")
}
