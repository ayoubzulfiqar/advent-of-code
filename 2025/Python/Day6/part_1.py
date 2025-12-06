from typing import List


class Data:
    def __init__(self):
        self.operators: List[str] = []
        self.positions: List[int] = []
        self.lines: List[str] = []


def individual_grand_totals() -> int:
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            content = file.read()
    except FileNotFoundError:
        print("Error: input.txt file not found")
        exit(1)
    except IOError as e:
        print(f"Error reading input.txt: {e}")
        exit(1)

    lines = content.split("\n")

    if lines and lines[-1] == "":
        lines = lines[:-1]

    if not lines:
        return 0

    operation_line = lines[-1]
    operators = []
    operator_columns = []

    for column, char in enumerate(operation_line):
        if char != " ":
            operators.append(char)
            operator_columns.append(column)

    operator_columns.append(len(operation_line))

    data_row_count = len(lines) - 1
    if data_row_count <= 0:
        return 0

    grand_total = 0

    for operator_index in range(len(operators)):
        current_operator = operators[operator_index]

        if current_operator == "*":
            column_result = 1
        else:
            column_result = 0

        column_start = operator_columns[operator_index]
        column_end = operator_columns[operator_index + 1]

        if operator_index < len(operators) - 1:
            column_end -= 1

        for row in range(data_row_count):
            parsed_number = 0

            for col in range(column_start, column_end):
                digit_char = lines[row][col]
                if digit_char != " ":
                    parsed_number = parsed_number * 10 + int(digit_char)

            if current_operator == "*":
                column_result *= parsed_number
            else:
                column_result += parsed_number

        grand_total += column_result

    return grand_total


if __name__ == "__main__":
    result = individual_grand_totals()
    print(f"Part 1: {result}")
