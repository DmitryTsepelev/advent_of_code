{:ok, input} = File.read("input.txt")
lines = input |> String.split("\n", trim: true)

earliest_departure = lines |> Enum.at(0) |> String.to_integer
buses = lines |> Enum.at(1) |> String.split(",") |> Enum.filter(fn el -> el != "x"; end) |> Enum.map(&String.to_integer/1)

bus = Enum.sort_by(buses, fn bus -> bus - rem(earliest_departure, bus); end) |> Enum.at(0)
answer1 = bus * (bus * ceil(earliest_departure / bus) - earliest_departure)
IO.puts "Part 1 answer is #{answer1}"
