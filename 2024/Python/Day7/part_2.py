from typing import Callable, List, Optional, Tuple


def parseInputFile() -> List[Tuple[int, List[int]]]:
    """
    Parse the input file to extract calibration equations.
    Each line is formatted as `result: term1 term2 ...`.
    """
    with open("./input.txt", "r") as file:
        lines = file.read().strip().splitlines()

    calibration_equations = []
    for line in lines:
        parts = line.split(":")
        result = int(parts[0].strip())
        terms = list(map(int, parts[1].strip().split()))
        calibration_equations.append((result, terms))

    return calibration_equations


def add(acc: Optional[int], term: int) -> int:
    """
    Add operator: sum the accumulator and the current term.
    """
    return (acc or 0) + term


def multiply(acc: Optional[int], term: int) -> int:
    """
    Multiply operator: multiply the accumulator and the current term.
    """
    return (acc or 1) * term


def concatenate(acc: Optional[int], term: int) -> int:
    """
    Concatenate operator: append the current term to the accumulator.
    """
    if acc is None:
        return term
    return int(f"{acc}{term}")


def validateEquation(
    result: int,
    terms: List[int],
    acc: Optional[int],
    operators: List[Callable[[Optional[int], int], int]],
) -> bool:
    """
    Recursively validate whether the result can be achieved using the terms
    and allowed operators.
    """
    # Base case
    if not terms:
        return result == acc

    # Recursive case: try all operators
    return any(
        validateEquation(result, terms[1:], op(acc, terms[0]), operators)
        for op in operators
    )


def elephantHidingTotalCalibration() -> int:
    """
    Solve part 2: Use add, multiply, and concatenate operators.
    """
    calibrationEquations = parseInputFile()
    return sum(
        result
        for result, terms in calibrationEquations
        if validateEquation(result, terms, None, [add, multiply, concatenate])
    )
