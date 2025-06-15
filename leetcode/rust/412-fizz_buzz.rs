pub fn fizz_buzz(n: i32) -> Vec<String> {
    let mut result = Vec::new();
    for i in 1..n + 1 {
        match (i % 3 == 0, i % 5 == 0) {
            (true, true) => result.push("FizzBuzz".to_string()),
            (true, false) => result.push("Fizz".to_string()),
            (false, true) => result.push("Buzz".to_string()),
            (false, false) => result.push(i.to_string())
        }
    }
    return result;
}

#[test]
fn test_fizz_buzz() {
    assert_eq!(fizz_buzz(15), vec!["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]);
}