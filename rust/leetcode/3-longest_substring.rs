use std::collections::HashMap;


pub fn length_of_longest_substring(s: String) -> i32 {
    let mut char_map = HashMap::new();
    let mut max_length = 0;
    let mut start = 0;

    for (index, ch) in s.chars().enumerate() {
        if let Some(&last_pos) = char_map.get(&ch) {
            start = start.max(last_pos + 1);
        }
        max_length = max_length.max(index - start + 1);
        char_map.insert(ch, index);
    }

    max_length as i32
}

#[test]
fn test_length_of_longest_substring(){
    assert_eq!(length_of_longest_substring("abcabcbb".to_string()), 3);
}