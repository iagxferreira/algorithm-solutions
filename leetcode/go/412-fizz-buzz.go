package hackerrank

func fizzBuzz(n int) []string {
    response := make([]string, 0)
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            response = append(response, "FizzBuzz")
        } else if i % 3 == 0 {
            response = append(response, "Fizz")
        } else if i % 5 == 0 {
            response = append(response, "Buzz")
        }else {
            response = append(response, strconv.Itoa(i))
        }
    }
    return response
}