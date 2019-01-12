package check

import "testing"

func TestParam(t *testing.T) {
	r := CheckTime(1500)
	t.Log(r)
}
