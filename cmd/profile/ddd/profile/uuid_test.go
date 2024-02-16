package profile

import "testing"

func TestGenRandomCode(t *testing.T) {
	s := createRandCode()
	t.Log(s)
}
