package recursive

var maxW = 0

func Zerone(Ulimit int, things []int, num int, sum int) {
	if num == len(things) || sum == Ulimit {
		if sum > maxW {
			maxW = sum
		}
		return
	}
	Zerone(Ulimit, things, num+1, sum)
	if sum+things[num] <= Ulimit {
		sum += things[num]
		if sum > maxW {
			maxW = sum
		}
		Zerone(Ulimit, things, num+1, sum)
	}
}
