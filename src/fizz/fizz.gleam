import gleam/option.{type Option}

pub fn fizz(n: Int) -> Option(String) {
  case is_fizz(n) {
    True -> option.Some("Fizz")
    False -> option.None
  }
}

pub fn is_fizz(n: Int) -> Bool {
  n % 3 == 0
}
