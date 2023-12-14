import Data.List (group, sort)

solve :: ([String] -> [String]) -> [String] -> Int
solve preprocess = length . filter (== 1) . map (maximum . map length . group . sort . preprocess . words)

main = do
  content <- readFile "input.txt"
  let lists = lines content

  print $ solve id lists
  print $ solve (map sort) lists
