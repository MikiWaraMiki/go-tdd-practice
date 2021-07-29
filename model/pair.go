package model

type IPair interface {
	GetFrom() string
	GetTo() string
}
type Pair struct {
	from string
	to   string
}

func NewPair(from string, to string) *Pair {
	return &Pair{
		from: from,
		to:   to,
	}
}

func (this Pair) GetFrom() string {
	return this.from
}

func (this Pair) GetTo() string {
	return this.to
}

func (this Pair) Equals(obj IPair) bool {
	if obj == nil {
		return false
	}

	if this.from != obj.GetFrom() {
		return false
	}

	if this.to != obj.GetTo() {
		return false
	}

	return true
}
