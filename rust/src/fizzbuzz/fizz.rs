pub fn fizz(n: usize) -> Option<String> {
    match is_fizz(n) {
        true => Some("Fizz".to_owned()),
        false => None,
    }
}

fn is_fizz(n: usize) -> bool {
    return n % 3 == 0;
}
