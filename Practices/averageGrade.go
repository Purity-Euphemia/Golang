type Student struct {
    Name  string
    Grade int
}

func averageGrade(students []Student) float64 {
    total := 0
    for _, s := range students {
        total += s.Grade
    }
    return float64(total) / float64(len(students))
}

func main() {
    students := []Student{
        {"Alice", 85},
        {"Bob", 90},
        {"Charlie", 78},
    }

    for _, s := range students {
        fmt.Println(s.Name, s.Grade)
    }

    avg := averageGrade(students)
    fmt.Println("Average grade:", avg)
}
