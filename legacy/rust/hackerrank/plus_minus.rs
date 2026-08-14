#[allow(dead_code)]
pub fn plus_minus(arr: &[i32]) {
    let mut count_positive: f32 = 0.0;
    let mut count_negative: f32 = 0.0;
    let mut count_zero: f32 = 0.0;
    for &item in arr.iter(){
        match item {
            item if item > 0 => count_positive += 1.0,
            item if item < 0 => count_negative += 1.0,
            _ => count_zero += 1.0,
        }
    }
    println!("{}", count_positive/(arr.len() as f32));
    println!("{}", count_negative/(arr.len() as f32));
    println!("{}", count_zero/(arr.len() as f32));
}