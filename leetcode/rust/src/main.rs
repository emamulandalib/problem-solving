mod two_sum;
mod longest_substring;

use crate::longest_substring::length_of_longest_substring;
use crate::two_sum::TwoSum;

fn main() {
    let c = length_of_longest_substring(String::from("pwwkew"));
    println!("{}", c)
}
