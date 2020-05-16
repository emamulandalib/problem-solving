mod two_sum;

use crate::two_sum::TwoSum;

fn main() {
    let ts = TwoSum::new(vec![1, 2, 3], 5);
    println!("{:?}", ts)
}
