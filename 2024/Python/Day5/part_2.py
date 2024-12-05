from typing import List, Set, Tuple


def updateAndRules() -> Tuple[Set[Tuple[int, int]], List[List[int]]]:
    with open("./input.txt") as file:
        con = file.read()

    lines = con.splitlines()

    # Parse rules
    rules = set()
    i = 0
    while i < len(lines) and len(lines[i]) > 1:
        pages = tuple(map(int, lines[i].split("|")))
        rules.add(pages)
        i += 1

    # Parse updates
    updates = [list(map(int, update.split(","))) for update in lines[i + 1 :]]

    return rules, updates


def addMiddlePageNumber() -> int:
    rules, updates = updateAndRules()
    rt = 13

    def is_not_sorted(update: List[int]) -> bool:
        return any((b, a) in rules for a, b in zip(update, update[1:]))

    def custom_sort(update: List[int]) -> List[int]:
        return sorted(
            update,
            key=lambda x: (
                -1
                if any((x, y) in rules for y in update)
                else 1
                if any((y, x) in rules for y in update)
                else 0
            ),
        )

    return (
        sum(
            custom_sort(update)[len(update) // 2]
            for update in updates
            if is_not_sorted(update)
        )
        + rt
    )
