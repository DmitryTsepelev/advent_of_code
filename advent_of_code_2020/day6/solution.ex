defmodule Answers do
  def parse_groups(lines) do
    {groups, _} = lines |> Enum.reduce({[], :start_new}, fn line, acc ->
      {groups, mode} = acc

      case String.length(line) do
        0 ->
            {groups, :start_new}
        _ ->
          person_answers = line |> String.graphemes

          case mode do
            :start_new -> {groups ++ [[person_answers]], :append}
            :append ->
              {previous_group, groups} = List.pop_at(groups, -1)

              {groups ++ [previous_group ++ [person_answers]], :append}
          end
      end
    end)

    groups
  end

  def anyone_with_yes(group) do
    group |> List.flatten |> Enum.uniq
  end

  def everyone_with_yes(group) do
    [first_group | rest_groups] = group

    rest_groups |> Enum.reduce(first_group, fn group, acc -> acc -- (acc -- group) end)
  end
end

{:ok, input} = File.read("input.txt")
lines = input |> String.split("\n")
groups = Answers.parse_groups(lines)

answer1 = groups |> Enum.map(&Answers.anyone_with_yes/1) |> Enum.map(&Enum.count/1) |> Enum.sum
IO.puts "Part 1 answer is #{answer1}"

answer2 =  groups |> Enum.map(&Answers.everyone_with_yes/1) |> Enum.map(&Enum.count/1) |> Enum.sum
IO.puts "Part 2 answer is #{answer2}"
