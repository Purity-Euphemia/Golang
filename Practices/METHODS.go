func (u User) IsAdult() bool {
    return u.Age >= 18
}
fmt.Println(u.IsAdult()) 
