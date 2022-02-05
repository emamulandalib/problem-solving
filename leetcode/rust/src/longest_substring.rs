use std::cmp::max;
use std::collections::HashMap;

pub fn length_of_longest_substring(s: String) -> i32 {
    let mut start = 0;
    let mut count = 0;
    let mut db = HashMap::new();
    s.chars().enumerate().for_each(|(i, c)| {
        match db.get(&c) {
            Some(v) => {
                if v + 1 > start {
                    start = v + 1
                }
            }
            None => {}
        }
        db.insert(c, i);
        count = max(count, i - start + 1)
    });
    return count as i32
}