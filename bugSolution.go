func main() {

        i := 10

        fmt.Println(i)

        defer func() {

            fmt.Println(i)

        }()

        i = 20

    } 

//Alternatively 

func main() {
        i := 10
        fmt.Println(i)
        defer fmt.Println(i) // captures the value of i at this point in the execution
        i = 20
    } 