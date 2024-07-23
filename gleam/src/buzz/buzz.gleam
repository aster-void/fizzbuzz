import gleam/option.{type Option}

pub fn buzz(n: Int) -> Option(String) {
  case is_buzz(n) {
    True -> option.Some("Buzz")
    False -> option.None
  }
}

fn is_buzz(n: Int) {
  n % 5 == 0
}
