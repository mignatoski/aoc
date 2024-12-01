import sys


def parseFile():
    a1 = []
    a2 = []
    rmap = {}
    with open("input.txt", "r") as f:
        for line in f:
            row = line.strip("\n").split("   ")
            l, r = row[0], row[1]
            a1.append(l)
            a2.append(r)
            if r in rmap:
                rmap[r] = rmap[r] + 1
            else:
                rmap[r] = 1

        a1 = sorted(a1)
        a2 = sorted(a2)

        sum = 0
        for idx, l in enumerate(a1):
            diff = abs(int(l) - int(a2[idx]))
            sum += diff

        print(sum)

        sum = 0
        for idx, l in enumerate(a1):
            if l in rmap:
                ss = int(l) * int(rmap[l])
                sum += ss

        print(sum)


def main() -> int:
    parseFile()
    return 0


if __name__ == "__main__":
    sys.exit(main())
