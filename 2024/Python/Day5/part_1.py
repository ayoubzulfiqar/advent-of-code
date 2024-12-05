from typing import List, Set, Tuple


def rulesAndUpdates() -> Tuple[Set[Tuple[int, int]], List[List[int]]]:
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


def middlePageNumber() -> int:
    rules, updates = rulesAndUpdates()

    def is_sorted(update: List[int]) -> bool:
        return all((b, a) not in rules for a, b in zip(update, update[1:]))

    return sum(update[len(update) // 2] for update in updates if is_sorted(update))
