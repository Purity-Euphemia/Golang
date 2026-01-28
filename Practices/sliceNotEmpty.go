func sliceNotEmpty(nums []int) error {
    if len(nums) == 0 {
        return errors.New("slice is empty")
    }
    return nil
}
