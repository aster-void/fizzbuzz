import buzz/buzz
import fizz/fizz
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, unwrap}

pub fn main() {
  fizzbuzz_range(1, 15)
  |> list.each(io.println)
}

pub fn fizzbuzz_range(start: Int, last: Int) -> List(String) {
  list.range(start, last)
  |> list.map(apply)
}

pub fn apply(n: Int) -> String {
  case fizz.fizz(n), buzz.buzz(n) {
    None, None -> int.to_string(n)
    fizz, buzz -> unwrap(fizz, "") <> unwrap(buzz, "")
  }
}
