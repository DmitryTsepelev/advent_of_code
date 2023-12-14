{:ok, input} = File.read("input.txt")

lines = input |> String.split("\n")

defmodule Passport do
  @required_fields ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
  @valid_ecls ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]

  def weak_valid?(passport) do
    @required_fields |> Enum.all?(fn key -> Map.has_key?(passport, key) end)
  end

  def strict_valid?(passport) do
    weak_valid?(passport) &&
      valid_byr?(passport) &&
      valid_iyr?(passport) &&
      valid_eyr?(passport) &&
      valid_hgt?(passport) &&
      valid_hcl?(passport) &&
      valid_ecl?(passport) &&
      valid_pid?(passport)
  end

  def parse_pasports(lines) do
    {passports, _} = lines |> Enum.reduce({[], :start_new}, fn line, acc ->
      {passports, mode} = acc

      case String.length(line) do
        0 ->
            {passports, :start_new}
        _ ->
          passport = line
            |> String.split(" ")
            |> Enum.map(fn pair -> pair |> String.split(":") end)
            |> Enum.into(%{}, fn [a, b] -> {a, b} end)

          case mode do
            :start_new -> {passports ++ [passport], :append}
            :append ->
              {previous_passport, passports} = List.pop_at(passports, -1)

              {passports ++ [Map.merge(passport, previous_passport)], :append}
          end
      end
    end)

    passports
  end

  def valid_byr?(passport) do
    byr = passport["byr"] |> String.to_integer

    byr >= 1920 && byr <= 2002
  end

  def valid_iyr?(passport) do
    iyr = passport["iyr"] |> String.to_integer

    iyr >= 2010 && iyr <= 2020
  end

  def valid_eyr?(passport) do
    eyr = passport["eyr"] |> String.to_integer

    eyr >= 2020 && eyr <= 2030
  end

  def valid_hgt?(passport) do
    captures = Regex.named_captures(~r/(?<value>\d*)(?<metric>(cm|in))/, passport["hgt"])

    case captures do
      nil -> false
      _ ->
        value = captures["value"] |> String.to_integer

        case captures["metric"] do
          "in" -> value >= 59 && value <= 76
          "cm" -> value >= 150 && value <= 193
        end
    end
  end

  def valid_hcl?(passport) do
    Regex.match?(~r/#[a-f0-9]{6}/, passport["hcl"])
  end

  def valid_ecl?(passport) do
    Enum.member?(@valid_ecls, passport["ecl"])
  end

  def valid_pid?(passport) do
    passport["pid"] |> String.length == 9
  end
end

parsed_pasports = Passport.parse_pasports(lines)

answer1 = parsed_pasports |> Enum.filter(&Passport.weak_valid?/1) |> Enum.count
IO.puts "Part 1 answer is #{answer1}"

answer2 = parsed_pasports |> Enum.filter(&Passport.strict_valid?/1) |> Enum.count
IO.puts "Part 2 answer is #{answer2}"
