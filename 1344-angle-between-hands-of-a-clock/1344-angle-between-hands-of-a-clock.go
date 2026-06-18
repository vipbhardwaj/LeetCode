func angleClock(hour int, minutes int) float64 {
    x := float64(hour) + (float64(minutes) / 60.0)
    diff := math.Mod(11.0 * x, 12.0)
    return min(diff, 12.0 - diff) * 30
}