func multiplicationTable(n int) []string {
    table := []string{}
    for i := 1; i <= 12; i++ {
        table = append(table, fmt.Sprintf("%d x %d = %d", n, i, n*i))
    }
    return table
}
