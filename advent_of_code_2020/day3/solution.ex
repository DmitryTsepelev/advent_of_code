defmodule Point do
  defstruct x: 0, y: 0
end

defmodule Field do
  @empty "."
  @tree "#"

  def reached_bottom?(field, current_position) do
    current_position.y >= Enum.count(field) - 1
  end

  def width(field) do
    field |> Enum.at(0) |> Enum.count
  end

  def move(field, current_position, right, down) do
    x = rem(current_position.x + right, width(field))
    y = current_position.y + down
    new_position = %Point{x: x, y: y}

    tile = field |> Enum.at(new_position.y) |> Enum.at(new_position.x)

    case tile do
      @empty -> {:empty, new_position}
      @tree -> {:tree, new_position}
    end
  end
end

defmodule Solver do
  def count_slope(right, down) do
    field = build_field("input.txt")
    count_trees(field, %Point{}, right, down, 0)
  end

  def count_trees(field, current_position, right, down, tree_count) do
    if Field.reached_bottom?(field, current_position) do
      tree_count
    else
      {result, new_position} = Field.move(field, current_position, right, down)

      new_tree_count = case result do
        :empty -> tree_count
        :tree -> tree_count + 1
      end

      count_trees(field, new_position, right, down, new_tree_count)
    end
  end

  def build_field(filename) do
    {:ok, input} = File.read(filename)
    input |> String.split("\n", trim: true) |> Enum.map(&String.graphemes/1)
  end
end

IO.puts "Part 1 answer is #{Solver.count_slope(3, 1)}"


answer2 = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]] |> Enum.reduce(1, fn slope, acc ->
  [right, down] = slope
  tree_count = Solver.count_slope(right, down)
  acc * tree_count
end)

IO.puts "Part 2 answer is #{answer2}"
