package recursive

var maxV = 0
var Plist []int
var Elist []int

func AfPackage(Wlimit int, cW int, cV int, cI int, Wlist []int, Vlist []int, Rlist []int) {
	if cI == len(Wlist) || cW == Wlimit {
		if cV > maxV {
			maxV = cV
			Plist = nil
			Plist = Rlist[:]
		}
		return
	}
	AfPackage(Wlimit, cW, cV, cI+1, Wlist, Vlist, Rlist)
	if cW+Wlist[cI] <= Wlimit {
		Tlist := append(Rlist, Vlist[cI])
		AfPackage(Wlimit, cW+Wlist[cI], cV+Vlist[cI], cI+1, Wlist, Vlist, Tlist)
	}
}
