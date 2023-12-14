import Data.List (findIndex)
import Data.Maybe (fromJust)
import Debug.Trace (trace)

-- move :: [(Int, Int)] -> Int -> [(Int, Int)]
-- move acc idxToMove = trace ("moving " ++ (show idxToMove) ++ " " ++ (show (map snd acc))) $ insertAt newElement (movedIdx + movedValue - 1) withoutElement
--   where
--     newElement = (idxToMove, movedValue)
--     withoutElement = deleteAt movedIdx acc
--     movedValue = snd $ acc!!movedIdx
--     movedIdx = fromJust $ findIndex (\(initialIdx, value) -> initialIdx == idxToMove) acc


insertAt :: a -> Int -> [a] -> [a]
insertAt newElement _ [] = [newElement]
insertAt newElement i (a:as)
  | i >= (length as + 1) = insertAt newElement (i `mod` (length as + 1)) (a:as)
  | i < 0 = insertAt newElement ((length as) + i) (a:as)
  | otherwise = trace (show i) $ a : insertAt newElement (i - 1) as

deleteAt :: Int -> [a] -> [a]
deleteAt idx xs = lft ++ rgt
  where (lft, (_:rgt)) = splitAt idx xs

-- mix :: [Int] -> [Int]
-- mix value = map snd $ foldl move (zip [0..] value) ([0..length value])

-- type Node = (Int, Bool)

-- fromList :: [Int] -> [Node]
-- fromList = map (\e -> (e, False))

mix :: [Int] -> [Int]
mix = mix' . map (\e -> (e, False))

mix' pairs
  | all snd pairs = map fst pairs
  | otherwise = do
    let idx = fromJust $ findIndex (\e -> snd e == False) pairs
    let value = fst $ pairs !! idx
    let deleted = deleteAt idx pairs
    if value > 0 then
      trace (show pairs) $ mix' $ insertAt (value, True) (idx + value - 1) deleted
    else
      mix' $ insertAt (value, True) (idx + value) deleted

main = do
  print $ mix [1, 2, 3]
  -- print $ mix [1, 2, -3, 3, -2, 0, 4]
