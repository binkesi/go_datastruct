package recursive

var Matched = false

func MatchReg(reg string, str string, rp int, sp int) {
	if rp == len(reg) && sp == len(str) {
		Matched = true
	}
	if rp == len(reg) || sp == len(str) {
		return
	}
	if reg[rp] != '*' {
		if reg[rp] == str[sp] {
			MatchReg(reg, str, rp+1, sp+1)
		} else if reg[rp] == '?' {
			MatchReg(reg, str, rp+1, sp)
			MatchReg(reg, str, rp+1, sp+1)
		} else {
			return
		}
	} else {
		for i := 0; i <= len(str)-sp; i++ {
			MatchReg(reg, str, rp+1, sp+i)
		}
	}
}
