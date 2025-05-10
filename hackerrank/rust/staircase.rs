#[allow(dead_code)]
pub fn staircase(n: i32) {
    for i in 0..n {
        let spaces = (n - i - 1) as usize;
        let hashes = (i + 1) as usize;
        let line: String = " ".repeat(spaces) + &"#".repeat(hashes);
        println!("{}", line);
    }
}