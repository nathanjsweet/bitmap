package bitmap

import (
	"fmt"
	"strings"
	"testing"
)

func TestBitMap(t *testing.T) {
	b := make(BitMap)

	b.Sets(2, 3, 4, 7, 11, 12, 14, 22, 24, 26, 32, 35, 37, 41, 42, 47, 49, 53, 62, 64, 256)
	var s []string
	b.Clear(3)
	for i := uint64(0); i <= 257; i++ {
		if b.IsSet(i) {
			s = append(s, fmt.Sprintf("%d", i))
		}
	}
	v := strings.Join(s, ", ")
	exp := "2, 4, 7, 11, 12, 14, 22, 24, 26, 32, 35, 37, 41, 42, 47, 49, 53, 62, 64, 256"
	if v != exp {
		t.Fatalf("bit set value not returned, got:\n%s\nexpected:\n%s\n", v, exp)
	}
	i := b.LeastSignificantZeroBit()
	if i != 0 {
		t.Fatalf("Expected least significant zero bit to be %d, but it was %d\n", 0, i)
	}
}
