func compoundInterest(p, r float64, t int) float64 {
    amount := p
    for i := 0; i < t; i++ {
        amount += amount * r / 100
    }
    return amount
}
