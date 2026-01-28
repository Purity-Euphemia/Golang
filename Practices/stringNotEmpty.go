func stringNotEmpty(s string) error {
    if s == "" {
        return errors.New("string is empty")
    }
    return nil
}
