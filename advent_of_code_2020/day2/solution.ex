{:ok, input} = File.read("input.txt")

lines = input |> String.split("\n", trim: true)

defmodule Passwords do
  def validate_occurences(lines) do
    each_line(lines, fn password, letter, min_count, max_count ->
      occurences = password |> String.graphemes |> Enum.count(fn c -> c == letter end)
      occurences >= min_count && occurences <= max_count
    end)
  end

  def validate_letter_in_range(lines) do
    each_line(lines, fn password, letter, start_index, end_index ->
      occurences =
        [start_index - 1, end_index - 1] |>
          Enum.count(fn idx -> password |> String.at(idx) == letter end)

      occurences == 1
    end)
  end

  def each_line(lines, solver) do
    Enum.count(lines, fn line ->
      [policy, password] = line |> String.split(": ")
      [range, letter] = policy |> String.split(" ")
      [lower_bound, upper_bound] =
        range |> String.split("-")
              |> Enum.map(&String.to_integer/1)

      solver.(password, letter, lower_bound, upper_bound)
    end)
  end
end

IO.puts "Part 1 answer is #{Passwords.validate_occurences(lines)}"
IO.puts "Part 2 answer is #{Passwords.validate_letter_in_range(lines)}"
