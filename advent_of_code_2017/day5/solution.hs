import Data.Char (digitToInt)
import Data.Map (Map, fromList, insert, (!))

type Instructions = Map Int Int

incAt :: Instructions -> Int -> (Int -> Int) -> Instructions
incAt instructions idx modOffset =
  insert idx newValue instructions
  where newValue = modOffset (instructions!idx)

interpret :: (Int -> Int) -> Instructions -> Int
interpret = interpret' 0 0

interpret' :: Int -> Int -> (Int -> Int) -> Instructions -> Int
interpret' current steps modOffset instructions
  | current < 0 || (current >= length instructions) = steps
  | otherwise = interpret' newCurrent (succ steps) modOffset newInstructions
  where
    newCurrent = current + instructions!current
    newInstructions = incAt instructions current modOffset

main = do
  content <- readFile "input.txt"
  let array = map (\word -> read word :: Int) $ lines content
  let input = fromList (zip [0..] array)

  print $ interpret succ input
  print $ interpret (\offset -> if offset >= 3 then pred offset else succ offset) input
