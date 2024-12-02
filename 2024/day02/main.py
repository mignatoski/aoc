import sys


def parseFile():
    with open("input.txt", "r") as f:
        count = 0
        for line in f:
            report = line.strip("\n").split(" ")
            int_report = list(map(int, report))
            if isSafe(int_report):
                count += 1
        print(count)

    with open("input.txt", "r") as f:
        count = 0
        for line in f:
            report = line.strip("\n").split(" ")
            int_report = list(map(int, report))
            for i in range(len(int_report)):
                new_report = int_report.copy()
                del new_report[i]
                print(int_report, new_report)
                if isSafe(new_report):
                    count += 1
                    break
        print("pt2", count)


def isSafe(arr: list) -> bool:
    diff_sign = 0
    p = 0
    for idx, n in enumerate(arr):
        if idx == 0:
            p = n
            continue
        diff = n - p
        if diff_sign == 0:
            diff_sign = diff
        if abs(diff) < 1 or abs(diff) > 3:
            print("diff", idx, arr, diff, diff_sign)
            return False
        if diff_sign * diff < 0:
            print("total", idx, arr, diff, diff_sign)
            return False
        p = n
    print("safe", arr)
    return True


def main() -> int:
    parseFile()
    return 0


if __name__ == "__main__":
    sys.exit(main())
