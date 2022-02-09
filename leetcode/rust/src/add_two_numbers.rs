#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}

pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut l1_current = l1;
    let mut l2_current = l2;

    while l1_current.is_some() || l2_current.is_some() {
        let l1_val = if l1_current.is_some() { l1_current.as_ref().unwrap().val } else { 0 };
        let l2_val = if l2_current.is_some() { l2_current.as_ref().unwrap().val } else { 0 };
        println!("{}", l1_val + l2_val);
        l1_current = if l1_current.is_some() { l1_current.as_ref().unwrap().next.clone() } else { None };
        l2_current = if l2_current.is_some() { l2_current.as_ref().unwrap().next.clone() } else { None };
    }
    return Option::from(Box::new(ListNode::new(2)));
}