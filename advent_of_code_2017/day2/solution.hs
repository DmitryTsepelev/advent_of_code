import Data.Char (digitToInt)

solve :: [[Int]] -> Int
solve = sum . map (\list -> maximum list - minimum list)

solve2 :: [[Int]] -> Int
solve2 lists =
    sum $ map (uncurry div . head . filter dividable . pairs) lists
    where
      pairs list = [(x, y) | x <- list, y <- list, x > y]
      dividable (x, y) = x `mod` y == 0

main = do
  content <- readFile "input.txt"
  let lists = map (map (\word -> read word :: Int) . words) . lines $ content

  print $ solve lists

  print $ solve2 lists
