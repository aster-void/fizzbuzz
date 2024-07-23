pub fn buzz(n: usize) -> Option<String> {
    match is_buzz(n) {
        true => Some("Buzz".to_string()),
        false => None,
    }
}

fn is_buzz(n: usize) -> bool {
    n % 5 == 0
}
