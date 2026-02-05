func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	panic("error")
}
