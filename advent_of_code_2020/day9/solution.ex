defmodule Weakness do
  def find_wrong_number(numbers, preamble) do
    cache = build_cache(numbers, preamble)

    Enum.find_value(preamble..Enum.count(numbers) - 1, fn idx ->
      number = Enum.at(numbers, idx)
      pairs = Map.get(cache, number)

      if pairs == nil do
        number
      else
        pair = find_pair(pairs, idx, preamble)
        if pair == nil, do: number, else: false
      end
    end)
  end

  def find_weakness(wrong_number, numbers) do
    range = find_weakness_range(wrong_number, numbers)
    number_range = Enum.slice(numbers, range)
    Enum.min(number_range) + Enum.max(number_range)
  end

  defp build_cache(numbers, preamble) do
    numbers |> Enum.with_index |> Enum.reduce(%{}, fn {x, x_idx}, acc ->
      start_idx = x_idx
      end_idx = x_idx + preamble - 1

      if end_idx >= Enum.count(numbers) do
        acc
      else
        Enum.reduce(start_idx..end_idx, acc, fn y_idx, acc ->
          sum = x + Enum.at(numbers, y_idx)
          pairs = Map.get(acc, sum, []) ++ [[x_idx, y_idx]]
          Map.merge(acc, %{sum => pairs})
        end)
      end
    end)
  end

  defp find_pair(pairs, idx, preamble) do
    pairs |> Enum.find(fn [x, y] ->
      start_idx = idx - preamble
      end_idx = idx - 1
      x >= start_idx && x <= end_idx && y >= start_idx && y <= end_idx
    end)
  end

  defp find_weakness_range(wrong_number, numbers) do
    0..Enum.count(numbers) |> Enum.find_value(fn start_idx ->
      [_, range] = Enum.reduce_while(start_idx..Enum.count(numbers) - 1, [0, nil], fn current_idx, acc ->
        sum = Enum.at(acc, 0) + Enum.at(numbers, current_idx)

        cond do
          sum == wrong_number -> {:halt, [sum, start_idx..current_idx]}
          sum < wrong_number -> {:cont, [sum, nil]}
          sum > wrong_number -> {:halt, [sum, nil]}
        end
      end)

      range
    end)
  end
end

{:ok, input} = File.read("input.txt")
numbers = input |> String.split("\n", trim: true) |> Enum.map(&String.to_integer/1)

answer1 = Weakness.find_wrong_number(numbers, 25)
IO.puts "Part 1 answer is #{answer1}"

answer2 = Weakness.find_weakness(answer1, numbers)
IO.puts "Part 2 answer is #{answer2}"
