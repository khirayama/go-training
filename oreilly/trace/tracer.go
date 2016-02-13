package trace

// Tracerはコード内での出来事を記録できるオブジェクトを表すインターフェースです。
type Tracer interface {
	Trace(...interface{}) // Goコミュニティではこのような引数を何個でも受け取れる定義が多い
}
