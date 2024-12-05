import sys
import re

part1 = 0
part2 = 0


def calcMul(m: re.Match) -> int:
    return int(m.group(1)) * int(m.group(2))


def parseFile2():
    with open("input.txt", "r") as f:
        data = f.read()
        enabled = True
        for m in re.finditer(r"mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)", data):
            match m.group(0):
                case "do()":
                    enabled = True
                case "don't()":
                    enabled = False
            if enabled and m.group(1):
                global part2
                part2 = calcMul(m) + part2


def parseFile():
    with open("input.txt", "r") as f:
        data = f.read()
        for m in re.finditer(r"mul\((\d{1,3}),(\d{1,3})\)", data):
            global part1
            part1 = calcMul(m) + part1


def main() -> int:
    parseFile()
    print("part 1", part1)
    parseFile2()
    print("part 2", part2)
    return 0


if __name__ == "__main__":
    sys.exit(main())
