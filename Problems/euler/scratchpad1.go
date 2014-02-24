




const (
	totientTableLength = 100000
)

var totientTable [totientTableLength]int64

func totient(n int64) int64 {
	if n < 2 {
		return 0
	}
	if n < totientTableLength && totientTable[n] != 0 {
		return totientTable[n]
	}

	factors := Factor(n)

	if factors[0] == factors[len(factors)-1] {
		answer := IntExp(factors[0], int64(len(factors))) - IntExp(factors[0], int64(len(factors)-1))
		if n < totientTableLength {
			totientTable[n] = answer
		}
		return answer
	}

	for i := 0; i < len(factors); i++ {
		if factors[i] != factors[0] {
			split := IntExp(factors[0], int64(i))
			answer := totient(split) * totient(n/split)
			if n < totientTableLength {
				totientTable[n] = answer
			}
			return answer
		}
	}

	//bad things have happenned if we're here
	return 0

}