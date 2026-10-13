package main
import "fmt"
func main() {
    status := "new"
    switch status {
    case "new":
        fmt.Println("Новая")
    case "in_progress":
        fmt.Println("В работе")
    case "done":
        fmt.Println("Завершена")
    default:
        fmt.Println("Неизвестный статус")
    }
}