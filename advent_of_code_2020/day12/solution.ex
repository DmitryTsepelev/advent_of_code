defmodule Interpreter do
  def route_ship(instructions) do
    ship = %{:east => 0, :south => 0, :dir => "E"}

    Enum.reduce(instructions, ship, fn instruction, ship ->
      {dir, string_value} = String.split_at(instruction, 1)
      value = String.to_integer(string_value)

      case dir do
        "L" -> rotate_ship(ship, -value)
        "R" -> rotate_ship(ship, value)
        "F" -> move(ship, ship.dir, value)
        _ -> move(ship, dir, value)
      end
    end)
  end

  def route_waypoint(instructions) do
    ship = %{:east => 0, :south => 0}
    waypoint = %{:east => 10, :south => -1}
    acc = %{:ship => ship, :waypoint => waypoint}

    Enum.reduce(instructions, acc, fn instruction, acc ->
      {dir, string_value} = String.split_at(instruction, 1)
      value = String.to_integer(string_value)

      case dir do
        "F" ->
          move_ship_to_waypoint(acc, value)
        "L" ->
          Map.merge(acc, %{:waypoint => rotate_waypoint(acc.waypoint, -value)})
        "R" ->
          Map.merge(acc, %{:waypoint => rotate_waypoint(acc.waypoint, value)})
        _ ->
          Map.merge(acc, %{:waypoint => move(acc.waypoint, dir, value)})
      end
    end)
  end

  defp move(point, dir, value) do
    case dir do
      "N" -> Map.merge(point, %{:south => point.south - value})
      "S" -> Map.merge(point, %{:south => point.south + value})
      "W" -> Map.merge(point, %{:east => point.east - value})
      "E" -> Map.merge(point, %{:east => point.east + value})
    end
  end

  @directions ["E", "S", "W", "N"]

  defp rotate_ship(point, value) do
    shift = div value, 90
    direction_idx = Enum.find_index(@directions, fn dir -> dir == point.dir; end) + shift
    new_dir = rem(direction_idx, Enum.count(@directions))
    Map.merge(point, %{:dir => Enum.at(@directions, new_dir)})
  end

  defp rotate_waypoint(point, value) do
    case value do
      90 -> %{:south => point.east, :east => - point.south}
      180 -> %{:south => - point.south, :east => - point.east}
      270 -> %{:south => - point.east, :east => point.south}
      -90 -> %{:south => - point.east, :east => point.south}
      -180 -> %{:south => - point.south, :east => - point.east}
      -270 -> %{:south => point.east, :east => - point.south}
    end
  end

  defp move_ship_to_waypoint(objects, times) do
    ship = %{
      east: objects.ship.east + times * objects.waypoint.east,
      south: objects.ship.south + times * objects.waypoint.south
    }

    Map.merge(objects, %{:ship => ship})
  end
end

{:ok, input} = File.read("input.txt")
instructions = input |> String.split("\n", trim: true)

ship = Interpreter.route_ship(instructions)
answer1 = abs(ship.east) + abs(ship.south)
IO.puts "Part 1 answer is #{answer1}"

objects = Interpreter.route_waypoint(instructions)
answer2 = abs(objects.ship.east) + abs(objects.ship.south)
IO.puts "Part 2 answer is #{answer2}"
