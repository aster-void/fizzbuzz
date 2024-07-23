import fizzbuzz
import gleam/int
import gleam/list
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

pub fn short_test() {
  fizzbuzz.fizzbuzz_range(1, 5)
  |> should.equal(["1", "2", "Fizz", "4", "Buzz"])
}

pub fn fizzbuzz_test() {
  [15, 30, 45]
  |> list.map(fizzbuzz.apply_fizzbuzz_on_single_number)
  |> list.each(should.equal(_, "FizzBuzz"))
}

pub fn only_fizz_test() {
  [3, 6, 9, 12]
  |> list.map(fizzbuzz.apply_fizzbuzz_on_single_number)
  |> list.each(should.equal(_, "Fizz"))
}

pub fn only_buzz_test() {
  [5, 10, 20, 25]
  |> list.map(fizzbuzz.apply_fizzbuzz_on_single_number)
  |> list.each(should.equal(_, "Buzz"))
}

pub fn number_test() {
  let ls = [1, 2, 4, 7, 8, 11, 13, 14]
  let fizzbuzzed = ls |> list.map(fizzbuzz.apply_fizzbuzz_on_single_number)
  ls
  |> list.map(int.to_string)
  |> list.strict_zip(fizzbuzzed)
  |> should.be_ok
  |> list.each(fn(elem: #(String, String)) { should.equal(elem.0, elem.1) })
}
