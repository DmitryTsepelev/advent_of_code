{:ok, input} = File.read("input.txt")

report = input |> String.split("\n", trim: true)
               |> Enum.map(&String.to_integer/1)
               |> Enum.sort

expected_sum = 2020

defmodule Report do
  def findSolution(report, number_count, sum) do
    Enum.reduce_while(report, nil, fn x, acc ->
      mult_result = case number_count do
        2 -> Enum.find(report, fn y -> x + y == sum end)
        _ -> findSolution(report, number_count - 1, sum - x)
      end

      case mult_result do
        nil -> {:cont, acc}
        _ -> {:halt, x * mult_result}
      end
    end)
  end
end

IO.puts "Part 1 answer is #{Report.findSolution(report, 2, 2020)}"
IO.puts "Part 2 answer is #{Report.findSolution(report, 3, 2020)}"
