type Greeter interface {
    Greet() string
}
func (u User) Greet() string {
    return "Hello, my name is " + u.Name
}
func sayHello(g Greeter) {
    fmt.Println(g.Greet())
}


