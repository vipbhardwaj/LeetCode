func internalAngles(sides []int) []float64 {
    slices.Sort(sides)
    a, b, c := sides[0], sides[1], sides[2]
    if a + b <= c { return []float64{} }

    angleC := math.Acos(float64(a*a + b*b - c*c) / float64(2*a*b)) * 180 / math.Pi
    angleB := math.Acos(float64(a*a + c*c - b*b) / float64(2*a*c)) * 180 / math.Pi
    angleA := 180 - angleB - angleC
    return []float64{angleA, angleB, angleC}
}