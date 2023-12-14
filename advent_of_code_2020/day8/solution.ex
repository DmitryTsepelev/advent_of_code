defmodule Instruction do
  defstruct op: nil, arg: nil

  def parse(lines) do
    lines |> Enum.map(fn line ->
      [op, arg] = line |> String.split(" ")

      %Instruction{op: op, arg: String.to_integer(arg)}
    end)
  end
end

defmodule Interpreter do
  def execute(program, acc \\ 0, instr_ptr \\ 0, log \\ []) do
    instruction = program |> Enum.at(instr_ptr)

    cond do
      instruction == nil -> {:completed, acc}
      Enum.member?(log, instr_ptr) -> {:loop, acc}
      true ->
        log = log ++ [instr_ptr]

        case instruction.op do
          "nop" -> execute(program, acc, instr_ptr + 1, log)
          "acc" -> execute(program, acc + instruction.arg, instr_ptr + 1, log)
          "jmp" -> execute(program, acc, instr_ptr + instruction.arg, log)
        end
    end
  end
end

defmodule Programmer do
  def fix(program) do
    program |> Enum.with_index |> Enum.find_value(fn {instruction, idx} ->
      if Enum.member?(["acc", "jmp"], instruction.op) do
        result = reprogram(program, idx) |> Interpreter.execute

        case result do
          {:completed, acc} -> acc
          {:loop, _} -> nil
        end
      end
    end)
  end

  def reprogram(program, idx) do
    old_instruction = Enum.at(program, idx)
    new_instruction = %Instruction{op: "nop", arg: old_instruction.arg}
    Enum.take(program, idx) ++ [new_instruction] ++ Enum.take(program, -(Enum.count(program) - idx - 1))
  end
end

{:ok, input} = File.read("input.txt")
lines = input |> String.split("\n", trim: true)
program = lines |> Instruction.parse

{_, answer1} = program |> Interpreter.execute
IO.puts "Part 1 answer is #{answer1}"

answer2 = program |> Programmer.fix
IO.puts "Part 2 answer is #{answer2}"
