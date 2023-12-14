defmodule BoardingPass do
  def seat_id(pass) do
    find_row(pass) * 8 + find_column(pass)
  end

  def find_row(pass) do
    pass |> String.slice(0..6) |> String.graphemes |> find_in_range([0, 127]) |> Enum.at(0)
  end

  def find_column(pass) do
    pass |> String.slice(-3..-1) |> String.graphemes |> find_in_range([0, 7]) |> Enum.at(0)
  end

  def find_in_range(letters, range) do
    letters |> Enum.reduce(range, fn letter, acc ->
      [start_index, end_index] = acc

      case letter do
        value when value in ["B", "R"] -> [start_index + ceil((end_index - start_index) / 2), end_index]
        value when value in ["F", "L"] -> [start_index, start_index + floor((end_index - start_index) / 2)]
      end
    end)
  end
end


{:ok, input} = File.read("input.txt")

seat_ids = input |> String.split("\n", trim: true) |> Enum.map(&BoardingPass.seat_id/1)

answer1 = seat_ids |> Enum.max
IO.puts "Part 1 answer is #{answer1}"

answer2 = seat_ids |> Enum.sort |> Enum.reduce_while(nil, fn id, acc ->
  case acc do
    nil -> {:cont, id}
    _ ->
      if id - acc == 1 do
        {:cont, id}
      else
        {:halt, acc + 1}
      end
  end
end)

IO.puts "Part 2 answer is #{answer2}"
