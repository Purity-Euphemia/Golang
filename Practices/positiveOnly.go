func positiveOnly(n int) error {
    if n < 0 {
        return errors.New("negative number not allowed")
    }
    return nil
}
