{-# OPTIONS_GHC -Wno-incomplete-patterns #-}
import Data.List (find)
import Data.Maybe (fromJust, isJust, isNothing)
import System.Posix (SystemID(nodeName))

data Node = Node { name :: String, childrenNames :: [String], weight :: Int } deriving Show

nodesFromInput :: [[Char]] -> [Node]
nodesFromInput = map parseWord

parseWord :: [Char] -> Node
parseWord string = Node (head cmp) (drop 3 $ map removeCommas cmp) weight
  where
    cmp = words string
    removeCommas = filter (/= ',')
    weight = read (cmp!!1) :: Int

getRootName :: [Node] -> String
getRootName nodes = fromJust $ find (`notElem` allChildrenNames) allNames
  where
    allNames = map name nodes
    allChildrenNames = concatMap childrenNames nodes

getNode :: String -> [Node] -> Node
getNode nodeName = fromJust . find (\node -> nodeName == name node)

setWeight :: String -> Int -> [Node] -> [Node]
setWeight nodeName weight =
  map (\node ->
    if name node == nodeName then
      Node nodeName (childrenNames node) weight
    else
      node
  )

calculateTotalWeigth :: String -> [Node] -> [Node]
calculateTotalWeigth nodeName nodes = do
  let node = getNode nodeName nodes
  let childNodes = map (`getNode` nodes) (childrenNames node)

  let nodesWithUpdatedChildren = foldl (\acc child -> calculateTotalWeigth (name child) acc) nodes childNodes

  let newWeigth = weight node + sum (map (\name -> weight $ getNode name nodes) (childrenNames node))

  setWeight nodeName newWeigth nodesWithUpdatedChildren

findWrongWeightDiff :: [Node] -> [Node] -> Maybe Int
findWrongWeightDiff currentLevel allNodes = do

  childNodes

-- findWrongWeightDiff current nodes = do
  -- let childNodes = map (`getNode` nodes) (childrenNames current)

  -- let maybeDiff = tryFindDiff childNodes

  -- if isJust maybeDiff then do
  --   let (maxWeight, minWeight) = fromJust maybeDiff
  --   Just (maxWeight - minWeight)
  -- else do
  --   let candidates = map (`findWrongWeightDiff` nodes) childNodes
  --   fromJust $ find isJust candidates


-- tryFindDiff :: [Node] -> Maybe (Node, Node)
-- tryFindDiff [x] = Nothing
-- tryFindDiff (x:y:rest)
--   | weight x > weight y = Just (x, y)
--   | weight y > weight x = Just (y, x)
--   | otherwise = tryFindDiff (y:rest)

main = do
  content <- readFile "input.txt"
  let nodes = nodesFromInput . lines $ content
  print nodes

  let allNames = map name nodes
  let allChildrenNames = concatMap childrenNames nodes

  let rootName = getRootName nodes
  print rootName

  let nodesWithTotalWeight = calculateTotalWeigth rootName nodes
  let rootNode = getNode rootName nodes
  print $ findWrongWeightDiff [rootNode] nodesWithTotalWeight
