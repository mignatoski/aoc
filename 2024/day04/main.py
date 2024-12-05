import sys
import numpy as np

matrix = []
rowlen = 0
collen = 0
part1 = 0
part2 = 0
XMAS = "XMAS"


def get_next_list(x, y, px, py):
    global matrix
    global part1
    ch = matrix[y][x]
    idx = XMAS.index(ch) + 1

    output = []

    # found X
    if idx == 1:
        for x1 in range(max(0, x - 1), min(x + 2, rowlen)):
            for y1 in range(max(0, y - 1), min(y + 2, collen)):
                if matrix[y1][x1] == XMAS[idx]:
                    output.append([x1, y1])
    else:
        ny = y + y - py
        nx = x + x - px
        if ny >= 0 and ny < collen and nx >= 0 and nx < rowlen:
            if matrix[ny][nx] == XMAS[idx]:
                output.append([nx, ny])

    if len(output) > 0:
        if idx == 3:
            part1 += len(output)
            for letter in output:
                print(XMAS[idx], letter[0], letter[1])
            return True
        else:
            for letter in output:
                if get_next_list(letter[0], letter[1], x, y):
                    print(XMAS[idx], letter[0], letter[1])
            return True

    return False


def parseFile():
    global matrix
    global rowlen
    global collen
    input = np.loadtxt("input.txt", dtype=str, delimiter=None)
    # input = np.loadtxt("sample.txt", dtype=str, delimiter=None)
    for r in input:
        matrix.append(list(r))

    collen = len(matrix)
    rowlen = len(matrix[0])

    for y, row in enumerate(matrix):
        for x, val in enumerate(row):
            if val == "X":
                get_next_list(x, y, 0, 0)
            if val == "A" and x > 0 and y > 0 and x < rowlen - 1 and y < collen - 1:
                part2_bfs(x, y)


def part2_bfs(x, y):
    pair1 = matrix[y - 1][x - 1] + matrix[y + 1][x + 1]
    pair2 = matrix[y + 1][x - 1] + matrix[y - 1][x + 1]
    print(pair1, pair2)
    global part2
    if pair1 == "MS" or pair1 == "SM":
        if pair2 == "MS" or pair2 == "SM":
            part2 += 1


def main() -> int:
    parseFile()
    print("part1", part1)
    print("part2", part2)
    return 0


if __name__ == "__main__":
    sys.exit(main())
