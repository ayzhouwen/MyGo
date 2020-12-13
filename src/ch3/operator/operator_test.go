package operator

import (
	"testing"
)

/**
按位清零
*/
func TestOperator(t *testing.T) {
	a := 255
	b := 1
	c := a &^ b
	t.Logf("a二进制:%b a十六进制:%X a十进制:%d", a, a, a)
	t.Logf("b二进制:%b b十六进制:%X b十进制:%d", b, b, b)
	t.Logf("c二进制:%b c十六进制:%X c十进制:%d", c, c, c)

}
