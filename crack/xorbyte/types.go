package xorbyte

type Scorer interface {
    Score(text []byte) float64
}

type functionScorer func([]byte) float64
func (this functionScorer) Score(text []byte) float64 {
    return this(text)
}
