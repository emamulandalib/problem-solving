mod two_sum;
mod longest_substring;
mod add_two_numbers;

use crate::add_two_numbers::*;
use crate::longest_substring::length_of_longest_substring;
use crate::two_sum::TwoSum;

fn main() {
    let mut list3: Option<Box<ListNode>> = None;
    let mut prev_node = None;

    for i in 1..10 {
        let new_node = Option::from(Box::from(ListNode::new(i)));
        if prev_node.is_none() {
            list3 = new_node;
            prev_node = list3;
            continue;
        }
        prev_node.unwrap().next = new_node;
        prev_node = prev_node.unwrap().next.clone();
    }

    // let _ = add_two_numbers(list3, list3);
}
