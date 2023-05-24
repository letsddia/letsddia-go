package t_digest

import (
	"lo"
	"math"
	"sort"
)

type Centroid struct {
	mean  float64
	count int
}
type TDigestV1 struct {
	centroids []Centroid
}

func NewTDigestV1() TDigestV1 {
	centroids := make([]Centroid, 0)
	return TDigestV1{centroids}
}
func (t *TDigestV1) Insert(x float64) {
	centroid := Centroid{
		mean:  x,
		count: 1,
	}
	t.centroids = append(t.centroids, centroid)
}
func (t *TDigestV1) Quantile(q float64) float64 {
	if len(t.centroids) == 0 {
		return math.NaN()
	}
	if len(t.centroids) == 1 {
		return t.centroids[0].mean
	}
	sort.Slice(t.centroids, func(i, j int) bool {
		return t.centroids[i].mean < t.centroids[j].mean
	})

	totalCount := lo.SumBy(t.centroids, func(c Centroid) int {
		return c.count
	})
	target := q * float64(totalCount)

	totalCount = 0
	for _, c := range t.centroids {
		if float64(totalCount+c.count) >= target {
			return c.mean
		}
		totalCount += c.count
	}
	return t.centroids[len(t.centroids)-1].mean
}
