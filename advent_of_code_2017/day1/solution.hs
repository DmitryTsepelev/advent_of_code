import Data.List ()
import Data.Char (digitToInt)

solve :: [Int] -> Int -> Int
solve digits shift = sum . map fst . filter equalPair . take (length digits) $ endlessPairs
  where endlessPairs = zip endlessDigits shiftedDigits
        endlessDigits = concat $ repeat digits
        shiftedDigits = drop shift endlessDigits
        equalPair = uncurry (==)

main = do
  content <- readFile "input.txt"
  let input = map digitToInt . head . lines $ content

  print $ solve input 1
  print $ solve input (length input `div` 2)
