func main() {

        i := 10

        fmt.Println(i)

        defer fmt.Println(i)

        i = 20

    }