package main
 
import (
    "fmt"
)
 
func main() {
    var t int
    counter := 0
    fmt.Scanln(&t)
    var guesses [5]int
    fmt.Scanln(&guesses[0], &guesses[1], &guesses[2], &guesses[3], &guesses[4])
    for _, guess := range guesses {
        if guess == t {
            counter++
        }
    }
    fmt.Println(counter)
}
