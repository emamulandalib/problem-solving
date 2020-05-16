mod two_sum;

use crate::two_sum::TwoSome;

fn main() {
    let ts = TwoSome::new(vec![1, 2, 3], 5);
    println!("{:?}", ts)
}
