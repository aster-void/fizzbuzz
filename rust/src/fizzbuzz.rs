mod buzz;
mod fizz;

pub fn apply(n: usize) -> String {
    match (fizz::fizz(n), buzz::buzz(n)) {
        (None, None) => n.to_string(),
        (None, Some(buzz)) => buzz,
        (Some(fizz), None) => fizz,
        (Some(fizz), Some(buzz)) => fizz + &buzz,
    }
}

pub fn range(start: usize, last: usize) -> Vec<String> {
    (start..=last).map(apply).collect::<Vec<_>>()
}
