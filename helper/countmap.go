package helper

import "sync"

type CountMap struct {
	sync.Mutex

	m map[byte]int
}

func NewCountMap(cap int) *CountMap {
	return &CountMap{
		m: make(map[byte]int, cap),
	}
}

func (m *CountMap) AddBytes(bytes ...byte) {
	m.Lock()
	defer m.Unlock()

	for _, b := range bytes {
		v, ok := m.m[b]
		if ok {
			v++
		} else {
			v = 1
		}

		m.m[b] = v
	}
}

func (m *CountMap) ToPairList() PairList {
	m.Lock()
	defer m.Unlock()

	pl := make(PairList, 0, len(m.m))

	for k, v := range m.m {
		pl = append(pl, Pair{
			K: k,
			V: v,
		})
	}

	return pl
}

func (m *CountMap) IsEmpty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.m) == 0
}
