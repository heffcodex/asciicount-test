package helper

type Pair struct {
	K byte
	V int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].V < p[j].V }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
