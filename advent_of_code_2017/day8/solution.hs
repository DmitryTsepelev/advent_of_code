{-# OPTIONS_GHC -Wno-incomplete-patterns #-}
import Prelude hiding (GT, LT, EQ)
import Data.Map (Map, insert, (!), fromList, toList)
import Distribution.Simple.Register (register)

type Register = String
data Condition = GT | LT | EQ | NEQ | GTEQ | LTEQ deriving Show
data Command = INC | DEC deriving Show

data Instruction = Instruction {
  targetRegister :: Register,
  command :: Command,
  argument :: Int,
  conditionRegister :: Register,
  condition :: Condition,
  conditionValue :: Int
} deriving Show

type Program = [Instruction]

instructionsFromInput :: [String] -> Program
instructionsFromInput = map parseInstruction

parseInstruction :: String -> Instruction
parseInstruction string =
  Instruction targetRegister command argument conditionRegister condition conditionValue
  where
    [targetRegister, rawCommand, rawArgument, _, conditionRegister, rawCondition, rawConditionValue] = words string

    command =
      case rawCommand of
        "inc" -> INC
        "dec" -> DEC
    argument = read rawArgument :: Int

    condition =
      case rawCondition of
        ">" -> GT
        "<" -> LT
        ">=" -> GTEQ
        "<=" -> LTEQ
        "==" -> EQ
        "!=" -> NEQ

    conditionValue = read rawConditionValue :: Int

type Registers = Map Register Int

execute :: Program -> [Registers]
execute instructions = execute' [registers] instructions
  where
    registers = fromList $ map (\instruction -> (targetRegister instruction, 0)) instructions

execute' :: [Registers] -> Program -> [Registers]
execute' =
  foldl (\history instruction ->
    performInstruction (head history) instruction:history
  )

performInstruction :: Registers -> Instruction -> Registers
performInstruction registers instruction
  | satisfyCondition instruction registers = applyInstruction registers instruction
  | otherwise = registers

applyInstruction :: Registers -> Instruction -> Registers
applyInstruction registers instruction =
  insert registerName newValue registers
  where
    registerName = targetRegister instruction
    commandFunction = commandToFunction (command instruction)
    currentValue = registers ! targetRegister instruction
    newValue = commandFunction currentValue (argument instruction)

commandToFunction :: Command -> Int -> Int -> Int
commandToFunction condition =
  case condition of
    INC -> (+)
    DEC -> (-)

satisfyCondition :: Instruction -> Registers -> Bool
satisfyCondition instruction registers =
  conditionFunction registerValue (conditionValue instruction)
  where
    conditionFunction = conditionToFunction (condition instruction)
    registerValue = registers ! conditionRegister instruction

conditionToFunction :: Condition -> Int -> Int -> Bool
conditionToFunction condition =
  case condition of
    GT -> (>)
    LT -> (<)
    EQ -> (==)
    NEQ -> (/=)
    GTEQ -> (>=)
    LTEQ -> (<=)

solve1 :: Program -> Int
solve1 = maximum . map snd . toList . head . execute

solve2 :: Program -> Int
solve2 = maximum . map snd . concatMap toList . execute

main = do
  content <- readFile "input.txt"
  let instructions = instructionsFromInput . lines $ content

  print $ solve1 instructions
  print $ solve2 instructions
