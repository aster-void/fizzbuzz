import buzz/buzz
import fizz/fizz
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, Some}

pub fn main() {
  fizzbuzz_range(1, 15)
  |> list.each(io.println)
}

pub fn fizzbuzz_range(start: Int, last: Int) -> List(String) {
  list.range(start, last)
  |> list.map(apply_fizzbuzz_on_single_number)
}

pub fn apply_fizzbuzz_on_single_number(n: Int) -> String {
  case fizz.fizz(n), buzz.buzz(n) {
    Some(fizz), Some(bazz) -> fizz <> bazz
    Some(fizz), None -> fizz
    None, Some(bazz) -> bazz
    None, None -> int.to_string(n)
  }
}
