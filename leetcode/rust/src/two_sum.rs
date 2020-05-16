use std::collections::HashMap;

pub struct TwoSome {}

impl TwoSome {
    pub fn new(nums: Vec<i32>, target: i32) -> Vec<i32> {
        return TwoSome::run(nums, target);
    }

    fn run(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut comp = HashMap::new();

        for i in 0..nums.len() {
            match comp.get(&nums[i]) {
                Some(v) => {
                    if *v >= 0 {
                        return vec![*v, i as i32];
                    }
                }
                _ => {}
            }

            comp.insert(target - nums[i], i as i32);
        }

        return vec![];
    }
}