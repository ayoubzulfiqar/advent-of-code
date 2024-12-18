bytes = [
    x + y * 1j
    for r in open("input.txt").readlines()
    for x, y in [map(int, r.strip().split(","))]
]
Width, Height, Front = 70, 70, 1024


def bfs(cut):
    start, goal, steps = 0, Width + Height * 1j, 0
    front, seen, walls = {start}, {start}, set(bytes[:cut])
    while front:
        new = set()
        steps += 1
        for pos in front:
            for d in [1, -1, 1j, -1j]:
                nPosition = pos + d
                if nPosition == goal:
                    return steps
                if not (0 <= nPosition.real <= Width and 0 <= nPosition.imag <= Height):
                    continue
                if nPosition in walls or nPosition in seen:
                    continue
                seen.add(nPosition)
                new.add(nPosition)
        front = new
    return 0


def firstByteReachablePosition():
    low, high = Front, len(bytes)

    while high - low > 1:
        i = (low + high) // 2
        if bfs(i):
            low = i
        else:
            high = i
    print(int(bytes[low].real), int(bytes[low].imag), sep=",")
