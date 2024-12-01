import sys


def parseFile():
    with open("input.txt", "r") as f:
        for line in f:
            num = int(line)
            print(num)


def main() -> int:
    parseFile()
    return 0


if __name__ == "__main__":
    sys.exit(main())
