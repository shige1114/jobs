package value

type SelfPR struct {
	Sentence string
}

func (pr *SelfPR) validate() bool {
	return len(pr.Sentence)/3 <= 200
}

func (pr *SelfPR) ToDB() string {
	return pr.Sentence
}

type GoodPoint struct {
	Sentence string
}

func (gp *GoodPoint) validate() bool {
	return len(gp.Sentence)/3 <= 200
}
func (pr GoodPoint) ToDB() string {
	return pr.Sentence
}

type ConcernPoint struct {
	Sentence string
}

func (cp *ConcernPoint) validate() bool {
	return len(cp.Sentence)/3 <= 200
}
func (pr ConcernPoint) ToDB() string {
	return pr.Sentence
}
