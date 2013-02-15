--import Debug.Trace
import Data.List
import Data.Maybe

numOneChild n = length [x | x <- [1..(n-1)], oneChild x]

oneChild :: Int -> Bool
oneChild x = judge (separateString strx) (length strx)
    where strx = show x

separateString :: String -> [Int]
separateString [x] = [read [x] :: Int]
separateString xs = separateString' xs (length xs) ++ separateString (tail xs)

separateString' :: String -> Int -> [Int]
separateString' xs 0 = []
separateString' xs n
    | isNothing (elemIndex num afterwards) = num:afterwards
    | otherwise = []
    where afterwards = separateString' xs (n - 1)
          num = read (take n xs) :: Int

judge xs y = True
--judge :: [Int] -> Int -> Bool
--judge xs y = divisible xs y == 1

--divisible :: [Int] -> Int -> Int
--divisible [] n = 0
--divisible (x:xs) n = divisible xs n + modit x n

--modit :: Int -> Int -> Int
--modit x y
--    | x `mod` y == 0 = 1
--    | otherwise = 0