#!/usr/bin/env bats

@test "adding 3 + 3" {
  run ./build/sum 3 3
  [ "$status" -eq 0 ]
  [ "$output" = "6" ]
}

@test "error when given string" {
  run ./build/sum 3 "XXXX"
  [ "$status" -ne 0 ]
}
