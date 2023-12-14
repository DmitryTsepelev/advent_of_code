defmodule Field do
  def build_field(filename) do
    {:ok, input} = File.read(filename)
    input |> String.split("\n", trim: true) |> Enum.map(&String.graphemes/1)
  end

  @directions [-1, 0, 1]

  def adjacent_occupied_seats_count(field, x, y) do
    row_count = height(field)
    col_count = width(field)

    count_seats_by(fn dx, dy ->
      y = y + dy
      x = x + dx

      if x >= 0 && x < col_count && y >= 0 && y < row_count do
        occupied_seat?(field, x, y)
      else
        false
      end
    end)
  end

  def occupied_seats_count(field, x, y) do
    row_count = height(field)
    col_count = width(field)

    count_seats_by(fn dx, dy ->
      direction_occupied?(field, col_count, row_count, x, y, dx, dy)
    end)
  end

  defp count_seats_by(direction_checker) do
    Enum.reduce(@directions, 0, fn dy, acc ->
      Enum.reduce(@directions, acc, fn dx, acc ->
        cond do
          dx == 0 && dy == 0 -> acc
          true -> if direction_checker.(dx, dy), do: acc + 1, else: acc
        end
      end)
    end)
  end

  defp direction_occupied?(field, col_count, row_count, x, y, dx, dy) do
    y = y + dy
    x = x + dx

    if x >= 0 && x < col_count && y >= 0 && y < row_count do
      cond do
        occupied_seat?(field, x, y) -> true
        empty_seat?(field, x, y) -> false
        true -> direction_occupied?(field, col_count, row_count, x, y, dx, dy)
      end
    else
      false
    end
  end

  def empty_seat?(field, x, y) do
    at(field, x, y) == "L"
  end

  def occupied_seat?(field, x, y) do
    at(field, x, y) == "#"
  end

  def at(field, x, y) do
    field |> Enum.at(y) |> Enum.at(x)
  end

  def width(field) do
    field |> Enum.at(0) |> Enum.count
  end

  def height(field) do
    field |> Enum.count
  end
end

defmodule Simulator do
  def stabilize(field) do
    solve(field, &move_seats_by_ajacent/1)
  end

  def stabilize_all(field) do
    solve(field, &move_seats_by_all/1)
  end

  defp solve(field, mover) do
    # Enum.each(field, fn row -> IO.puts(row); end)
    new_field = mover.(field)

    # IO.puts ""
    # Enum.each(new_field, fn row -> IO.puts(row); end)

    if new_field == field do
      Enum.map(new_field, &count_occupied/1) |> Enum.sum
    else
      solve(new_field, mover)
    end
  end

  @ajacent_tolerance 4

  def move_seats_by_ajacent(field) do
    move_seats(field, &Field.adjacent_occupied_seats_count/3, @ajacent_tolerance)
  end

  @all_tolerance 5

  def move_seats_by_all(field) do
    move_seats(field, &Field.occupied_seats_count/3, @all_tolerance)
  end

  def move_seats(field, counter, tolerance) do
    row_count = Field.height(field) - 1
    col_count = Field.width(field) - 1

    Enum.map(0..row_count, fn y ->
      Enum.map(0..col_count, fn x ->
        occupied_seats_count = counter.(field, x, y)

        cond do
          Field.empty_seat?(field, x, y) && occupied_seats_count == 0 -> "#"
          Field.occupied_seat?(field, x, y) && occupied_seats_count >= tolerance -> "L"
          true -> Field.at(field, x, y)
        end
      end)
    end)
  end

  defp count_occupied(row) do
    row |> Enum.count(fn seat -> seat == "#"; end)
  end
end

field = Field.build_field("input.txt")

answer1 = Simulator.stabilize(field)
IO.puts "Part 1 answer is #{answer1}"

answer2 = Simulator.stabilize_all(field)

IO.puts "Part 2 answer is #{answer2}"
