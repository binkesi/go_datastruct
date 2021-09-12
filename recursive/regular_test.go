package recursive

import "testing"

func TestMatchReg(t *testing.T) {
	reg := "sung?andwm*gether"
	str := "sungnandwmxtogether"
	MatchReg(reg, str, 0, 0)
	if Matched != true {
		t.Errorf("Regex matched failed")
	}
}
