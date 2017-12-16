package bitmap

type Bitmap map[uint64]uint64

func (b Bitmap) IsSet(i uint64) bool {
	idx := i / 64
	v, ok := b[idx]
	return ok && v&(1<<(i%64)) != 0
}

func (b Bitmap) Set(i uint64) {
	idx := i / 64
	_, ok := b[idx]
	if !ok {
		b[idx] = 0
	}
	bitSet := i % 64
	b[idx] |= 1 << bitSet
}

func (b Bitmap) Clear(i uint64) {
	idx := i / 64
	_, ok := b[idx]
	if !ok {
		return
	}
	bitSet := i % 64
	b[idx] ^= 1 << bitSet
	if b[idx] == 0 {
		delete(b, idx)
	}
}

func (b Bitmap) Sets(xs ...uint64) {
	for _, x := range xs {
		b.Set(x)
	}
}

func (b Bitmap) LeastSignificantZeroBit() uint64 {
	var v uint64
	h := ^uint64(0)
	for i, k := range b {
		if k&0xffffffff != 0xffffffff {
			if i < h {
				h = i
				v = k
			}
		}
	}
	v = ^v
	v &= (^v + 1)
	return uint64(v+h*64) / 2
}
