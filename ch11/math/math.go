package math

func Average(xs []float64) float64{
	total := float64(0)
	for _, x := range xs{
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64{
	mn := float64(100000)
	for _, x := range xs{
		if x<=mn{
			mn = x
		}
	}
	return mn
}

func Max(xs []float64) float64{
	mx := float64(0)
	for _, x := range xs{
		if x >= mx {
			mx = x
		}
	}
	return mx
}

