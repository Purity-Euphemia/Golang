func safeIndex(slice []int, index int) (int, error) {
    if index < 0 || index >= len(slice) {
        return 0, errors.New("index out of range")
    }
    return slice[index], nil
}
