import qualified Data.Set as S
import Data.Maybe (isJust, fromJust, catMaybes)
import Data.List (filter)
import Debug.Trace (trace)

data Cost = Cost { oreCost :: Int, clayCost :: Int, obsidianCost :: Int }

data Blueprint = Blueprint {
  oreRobotCost :: Cost,
  clayRobotCost :: Cost,
  obsidianRobotCost :: Cost,
  geodeRobotCost :: Cost
}

data Robots = Robots { oreRobots :: Int, clayRobots :: Int, obsidianRobots :: Int, geodeRobots :: Int } deriving Show

data Resources = Resources { ore :: Int, clay :: Int, obsidian :: Int, openGeodes :: Int } deriving Show

validateResources :: Resources -> Maybe Resources
validateResources resources@(Resources ore clay obsidian _) =
  if ore >= 0 && clay >= 0 && obsidian >= 0 then
    Just resources
  else
    Nothing

tryBuildRobot :: Resources -> Cost -> Maybe Resources
tryBuildRobot (Resources ore clay obsidian openGeodes) (Cost oreCost clayCost obsidianCost) =
  validateResources (Resources (ore - oreCost) (clay - clayCost) (obsidian - obsidianCost) openGeodes)

buildRobots :: Blueprint -> (Robots, Resources) -> [(Robots, Resources)]
buildRobots blueprint (robots@(Robots oreRobots clayRobots obsidianRobots geodeRobots), resources) =
  catMaybes [Just (robots, resources), maybeOreRobot, maybeClayRobot, maybeObsidianRobot, maybeGeodeRobot]
  where
    maybeOreRobot = do
      let cost = oreRobotCost blueprint
      let newResources = tryBuildRobot resources cost
      if isJust newResources then
        Just (Robots (oreRobots + 1) clayRobots obsidianRobots geodeRobots, fromJust newResources)
      else
        Nothing

    maybeClayRobot = do
      let cost = clayRobotCost blueprint
      let newResources = tryBuildRobot resources cost
      if isJust newResources then
        Just (Robots oreRobots (clayRobots + 1) obsidianRobots geodeRobots, fromJust newResources)
      else
        Nothing

    maybeObsidianRobot = do
      let cost = obsidianRobotCost blueprint
      let newResources = tryBuildRobot resources cost
      if isJust newResources then
        Just (Robots oreRobots clayRobots (obsidianRobots + 1) geodeRobots, fromJust newResources)
      else
        Nothing

    maybeGeodeRobot = do
      let cost = geodeRobotCost blueprint
      let newResources = tryBuildRobot resources cost
      if isJust newResources then
        Just (Robots oreRobots clayRobots obsidianRobots (geodeRobots + 1), fromJust newResources)
      else
        Nothing

makeStep :: Blueprint -> (Robots, Resources) -> S.Set (Robots, Resources)
makeStep blueprint (robots@(Robots oreRobots clayRobots obsidianRobots geodeRobots), resources@(Resources ore clay obsidian openGeodes)) = do
  S.fromList $ map (\(robots, resources) -> (robots, produceResources resources)) newStates
  where
    newStates = buildRobots blueprint (robots, resources)
    produceResources (Resources ore clay obsidian openGeodes) = Resources (ore + oreRobots) (clay + clayRobots) (obsidian + obsidianRobots) (openGeodes + geodeRobots)

simulate :: Int -> Blueprint -> S.Set (Robots, Resources) -> S.Set (Robots, Resources)
simulate 0 _ states = states
simulate n blueprint states =
  trace ((show n) ++ " states: " ++ (show$ length states)) $ simulate (n - 1) blueprint states'
  where
    -- S.union . makeStep blueprint
    states' = foldl (\state acc -> qoo acc blueprint state) S.empty states

    qoo acc blueprint state = acc ++ (makeStep blueprint state)

    filterBad candidates = do
      if n > 10 then
        candidates
      else do
        let bestCandidate = maximum $ map (openGeodes . snd) candidates
        filter (\state@(Robots _ _ _ geodeRobots, (Resources _ _ _ geodes)) -> geodes + geodeRobots * n >= bestCandidate) candidates

main = do
  let bp = Blueprint (Cost 4 0 0) (Cost 2 0 0) (Cost 3 14 0) (Cost 2 0 7)
  let robots = Robots 1 0 0 0
  let resources = Resources 0 0 0 0

  let step = makeStep bp
  print $ maximum $ map (openGeodes . snd) $ simulate 24 bp (S.fromList [(robots, resources)])
