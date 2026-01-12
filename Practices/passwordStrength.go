func passwordStrength(password string) string {
    if len(password) >= 8 {
        return "Strong"
    }
    return "Weak"
}
