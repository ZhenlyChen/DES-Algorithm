package des

import "fmt"

var ipTable []byte
var ipInvTable []byte

func ipTransform(raw []byte) []byte {
	if len(raw)%8 != 0 {
		panic("Error length")
	} else if len(ipTable) == 0 {
		makeIpTable()
	}
	fmt.Println(raw)
	times := len(raw) / 8
	res := make([]byte, len(raw))
	table := []uint8{2, 4, 6, 8, 1, 3, 5, 7}
	for i := 0; i < times; i++ {
		current := raw[i*8 : (i+1)*8]
		for j := 0; j < 8; j++ {
			var b byte = 0
			for k := 0; k < 8; k++ {
				b |= current[7-k] & (1 << (8 - table[j])) & (1 << uint8(7 - k))
				fmt.Println(b, current[7-k], 1 << (8 - table[j]))
			}
			res[i * 8 + j] = b
		}
	}
	fmt.Println(res)
	return res
}

func ipInverseTransform(raw []byte) []byte {
	if len(ipInvTable) == 0 {
		makeIpInvTable()
	}
	return raw
}

func makeIpTable() {
	for i := 0; i <= 64; i++ {
		ipTable = append(ipTable, byte(getIpPos(i)))
	}
}

func getIpPos(pos int) int {
	add := pos / 33
	if pos > 32 {
		pos -= 32
	}
	pos--
	return 58 - 8*(pos%8) + (pos/8)*2 - add
}

func makeIpInvTable() {
	for i := 0; i <= 64; i++ {
		ipInvTable = append(ipInvTable, byte(getIpInvPos(i)))
	}
}

func getIpInvPos(x int) int {
	x--
	return 40 + 8*((x%8)/2) - x/8 - 32*(x%2)
}
