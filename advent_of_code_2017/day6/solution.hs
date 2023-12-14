import Data.List (elemIndex)
import Data.Maybe (fromJust)

solve2 :: [Int] -> Int
solve2 banks = solve2' [banks]

solve2' :: [[Int]] -> Int
solve2' banks
  | length (filter (==newBank) banks) == 2 = length banks
  | otherwise = solve2' (newBank:banks)
  where newBank = redistributeSlow $ head banks

solve :: [Int] -> Int
solve banks = solve' [banks]

solve' :: [[Int]] -> Int
solve' banks
  | newBank `elem` banks = length banks
  | otherwise = solve' (newBank:banks)
  where newBank = redistributeSlow $ head banks

-- redistribute :: [Int] -> [Int]
-- redistribute memory = do
--   zipWith (curry distribute) [0..] memory
--   where
--     maxIdx = fromJust $ elemIndex max memory
--     max = maximum memory
--     addedToAllBanks = max `div` length memory
--     reminingBlocks = max `rem` length memory
--     distribute (idx, value)
--       | idx == maxIdx = addedToAllBanks
--       | idx > maxIdx && idx <= maxIdx + reminingBlocks ||
--         idx < maxIdx && idx <= (reminingBlocks - maxIdx) = value + addedToAllBanks + 1
--       | otherwise = value + addedToAllBanks

redistributeSlow :: [Int] -> [Int]
redistributeSlow memory = do
  redistributeSlow' max (succ maxIdx) (replace memory maxIdx 0)
  where
    maxIdx = fromJust $ elemIndex max memory
    max = maximum memory

redistributeSlow' blocks currentIdx memory
  | blocks == 0 = memory
  | currentIdx == length memory = redistributeSlow' blocks 0 memory
  | otherwise = redistributeSlow' (pred blocks) (succ currentIdx) (replace memory currentIdx (succ (memory!!currentIdx)))

replace memory idx value = (take idx memory) ++ [value] ++ (drop (succ idx) memory)

main = do
  -- print $ redistribute [0, 2, 7 ,0]
  -- print $ redistribute [2, 4, 1 ,2]
  -- print $ redistribute [3, 1, 2 ,3]
  -- print $ redistribute [0,2,3,4]
  -- print $ redistribute  [1,3,4,1]
  content <- readFile "input.txt"
  let input = map (\word -> read word :: Int) . words . head . lines $ content
  -- print $ redistribute $ redistribute $ redistribute input

  print $ solve input
  print $ solve2 input - solve input
