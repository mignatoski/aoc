import sys
from collections import defaultdict

rules_b = defaultdict(list)
rules_a = defaultdict(list)
part1 = 0
part2 = 0


def parseFile():
    with open("input.txt", "r") as f:
        # with open("sample.txt", "r") as f:
        global rules_b
        global rules_a
        global part1
        section = "rules"
        for line in f:
            if line == "\n":
                print(rules_b)
                print(rules_a)
                section = "updates"
                continue
            if section == "rules":
                data = line.strip().split("|")
                rules_b[int(data[0])].append(int(data[1]))
                rules_a[int(data[1])].append(int(data[0]))
            if section == "updates":
                data = line.strip().split(",")
                print(data)
                list_b = []
                good = True
                for page in data:
                    intersect = list(set(rules_b[int(page)]) & set(list_b))
                    if len(intersect) > 0:
                        good = False
                        calc_part2(data)
                        break
                    else:
                        list_b.append(int(page))
                print(good)
                if good:
                    part1 += int(data[(len(data) - 1) // 2])


def calc_part2(data: list):
    global part2
    int_list = []
    for item in data:
        added = False
        for i, entry in enumerate(int_list):
            if entry in set(rules_b[int(item)]):
                int_list.insert(i, int(item))
                added = True
                break
        if not added:
            int_list.append(int(item))
    print(int_list)
    part2 += int_list[len(data) // 2]


def main() -> int:
    parseFile()
    print(part1)
    print(part2)
    return 0


if __name__ == "__main__":
    sys.exit(main())
