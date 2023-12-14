{-# OPTIONS_GHC -Wno-incomplete-patterns #-}

import qualified Data.Map as M
import Data.List (isPrefixOf, find)

import Debug.Trace (trace)
import Data.Maybe (fromJust)

data Instruction = Const Integer | Function (Monkeys -> Integer)
type Monkeys = M.Map String Instruction

parseLine :: String -> (String, Instruction)
parseLine s = (name, fn)
  where
    name = take 4 s
    fn = buildFn (drop 5 s)

buildFn :: String -> Instruction
buildFn s
  | length cmp == 1 = Function $ const (read (head cmp) :: Integer)
  | otherwise = Function $ \monkeys -> fn (resolve monkeys (head cmp)) (resolve monkeys (last cmp))
  where
    cmp = words s
    fn = case cmp!!1 of
      "+" -> (+)
      "-" -> (-)
      "*" -> (*)
      "/" -> div

resolve :: Monkeys -> String -> Integer
resolve monkeys name = do
  let fn = monkeys M.! name
  case fn of
    (Const value) -> value
    (Function fn) -> fn monkeys

-- findHumn monkeys varMonkey target (l, r) i = do
--   -- let humn = l + (r - l) `div` 2
--   let humn = ((l + r) `div` 2)
--   let result = resolve (M.insert "humn" (Const humn) monkeys) varMonkey

--   if result == target || i == 120 then
--     trace ("= " ++ (show humn) ++ " " ++ (show result) ++ " " ++ (show target) ++ " " ++ (show (l, r)))$
--     humn
--   else if result > target then
--     trace ("> " ++ (show humn) ++ " " ++ (show result) ++ " " ++ (show target) ++ " " ++ (show (l, r))) $
--     findHumn monkeys varMonkey target (l, humn - 1) (i + 1)
--   else
--     trace ("< " ++ (show humn) ++ " " ++ (show result) ++ " " ++ (show target) ++ " " ++ (show (l, r))) $

--     findHumn monkeys varMonkey target  (humn + 1, r) (i + 1)

isVariable :: Monkeys -> String -> Bool
isVariable monkeys name =
  trace ("isVariable " ++ (name) ++ " " ++ (show (resolve (withHumn 1) name)) ++ " " ++ (show (resolve (withHumn 100) name))) $ resolve (withHumn 1) name /= resolve (withHumn 100) name
  where withHumn value = M.insert "humn" (Const value) monkeys

findRootVariables lines =
  (take 4 . drop 6 $ line, take 4 . drop 13 $ line)
  where
    line = fromJust . find (\row -> "root" `isPrefixOf` row) $ lines

rotateKeys :: M.Map String Instruction -> [String] -> [String] -> M.Map String Instruction
rotateKeys acc lines = foldl (\acc currentName -> M.union acc (rotateKey lines acc currentName)) acc

rotateKey :: [String] -> M.Map String Instruction -> String -> M.Map String Instruction
rotateKey lines acc currentName = do
  let line = fromJust . find (\row -> currentName `isPrefixOf` row) $ lines

  if length (words line) == 2 then
    M.union (M.singleton currentName (Const (read (last . words $ line )))) acc
  else do
    let var1 = take 4. drop 6 $ line
    let var2 = take 4. drop 13 $ line
    let (nextVarName, nextConstName) = if var1 == currentName then (var1, var2) else (var2, var1)
    let operation = head . drop 10 $ line
    let fn = buildRotatedFn operation var1 var2 currentName nextVarName nextConstName

    if nextVarName == "root" then
      M.union acc (M.singleton nextVarName fn)
    else
      rotateKeys (M.union acc (M.singleton nextVarName fn)) lines [nextVarName, currentName]

buildRotatedFn :: Eq a => Char -> a -> p -> [Char] -> a -> [Char] -> Instruction
buildRotatedFn operation var1 var2 currentName nextVarName nextConstName =
  Function fn
  where
  fn
    | operation == '+' = \monkeys -> (resolve monkeys nextConstName) - (resolve monkeys currentName)
    | operation == '-' = if nextVarName == var1 then
            \monkeys -> (resolve monkeys nextConstName) + (resolve monkeys currentName)
          else
            \monkeys -> (resolve monkeys nextConstName) - (resolve monkeys currentName)
    | operation == '*' = \monkeys -> (resolve monkeys currentName) `div` (resolve monkeys nextConstName)
    | operation == '/' = if nextVarName == var1 then
            \monkeys -> (resolve monkeys nextConstName) * (resolve monkeys currentName)
          else
            \monkeys -> (resolve monkeys nextConstName) `div` (resolve monkeys currentName)
main = do
  content <- readFile "input.txt"
  let monkeys = M.fromList . map parseLine . lines $ content
  print $ resolve monkeys "root"

  -- let monkeysWithoutRootAndHum = foldl (flip M.delete) monkeys ["root", "humn"]
  -- print $ resolve monkeys "pppw"
  -- print $ resolve monkeys "sjmn"
  -- print $ isVariable monkeys "jhpn"
  -- print $ isVariable monkeys "jmsg"
  -- let varMonkey = if isVariable monkeys "jhpn" then "jhpn" else "jmsg"
  -- let target = if isVariable monkeys "jhpn" then resolve monkeys "jmsg" else resolve monkeys "jhpn"
  -- print $ "varMonkey " ++ (show varMonkey)
  -- print $ resolve monkeys "jmsg"
  -- print $ "target " ++ (show target)

  -- print $ findHumn monkeys varMonkey target (1, maxBound) 0

  -- print "---"
  -- print $ resolve monkeys "jmsg"
  -- print $ resolve (M.insert "humn" (Const (3305669217842)) monkeys) "jhpn"
  -- print target
  -- print $ resolve (M.insert "jhpn" (Const 9223372036854775807) monkeys) "jmsg"

  -- print $ isVariable monkeys "pppw"
  -- print $ isVariable monkeys "sjmn"

  let (name1, name2) = findRootVariables . lines $ content
  let (varMonkey, targetName) = if isVariable monkeys name1 then (name1, name2) else (name2, name1)
  let rotated = rotateKeys M.empty (lines content) ["humn"]
  print $ resolve rotated "humn"

  -- let target = resolve monkeys targetName
  -- print $ "varMonkey " ++ (show varMonkey)
  -- print $ "targetName " ++ (show targetName)
  -- print $ "target " ++ (show target)

  -- print $ findHumn monkeys varMonkey target (1, 1000000000000) 0

  -- print $ resolve (M.insert "humn" (Const 303) monkeys) varMonkey
  -- print $ resolve (M.insert "humn" (Const 302) monkeys) varMonkey
  -- print $ resolve (M.insert "humn" (Const 301) monkeys) varMonkey
  -- print $ resolve (M.insert "humn" (Const 300) monkeys) varMonkey
