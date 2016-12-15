package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	return int(pc[(x>>(0*8))] +
		pc[(x>>(1*8))] +
		pc[(x>>(2*8))] +
		pc[(x>>(3*8))] +
		pc[(x>>(4*8))] +
		pc[(x>>(5*8))] +
		pc[(x>>(6*8))] +
		pc[(x>>(7*8))])
}
