package ichigo

//Point struct
type Point struct {
	X float64
	Y float64
}

func linearRegressionLSE(series []Point) []Point {

	q := len(series)

	if q == 0 {
		return make([]Point, 0, 0)
	}

	p := float64(q)

	sum_x, sum_y, sum_xx, sum_xy := 0.0, 0.0, 0.0, 0.0

	for _, p := range series {
		sum_x += p.X
		sum_y += p.Y
		sum_xx += p.X * p.X
		sum_xy += p.X * p.Y
	}

	m := (p*sum_xy - sum_x*sum_y) / (p*sum_xx - sum_x*sum_x)
	b := (sum_y / p) - (m * sum_x / p)

	r := make([]Point, q, q)

	for i, p := range series {
		r[i] = Point{p.X, (p.X*m + b)}
	}

	return r
}
