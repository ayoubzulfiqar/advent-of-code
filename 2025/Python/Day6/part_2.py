def individual_grand_answers_totals():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            content_string = file.read()
    except FileNotFoundError:
        print("Error: input.txt file not found")
        exit(1)
    except IOError as e:
        print(f"Error reading input.txt: {e}")
        exit(1)

    all_lines = content_string.split("\n")

    if all_lines and all_lines[-1] == "":
        all_lines = all_lines[:-1]

    if not all_lines:
        return 0

    bottom_operation_line = all_lines[-1]
    operators = []
    operator_column_starts = []

    for column_index, char in enumerate(bottom_operation_line):
        if char != " ":
            operators.append(char)
            operator_column_starts.append(column_index)

    operator_column_starts.append(len(bottom_operation_line))

    data_rows_count = len(all_lines) - 1
    if data_rows_count <= 0:
        return 0

    total_answer_sum = 0

    for operator_idx in range(len(operators)):
        current_operator = operators[operator_idx]
        vertical_column_result = 0

        if current_operator == "*":
            vertical_column_result = 1

        group_start_column = operator_column_starts[operator_idx]
        group_end_column = operator_column_starts[operator_idx + 1]

        if operator_idx < len(operators) - 1:
            group_end_column -= 1

        for current_column in range(group_start_column, group_end_column):
            vertical_number = 0

            for row_index in range(data_rows_count):
                character = all_lines[row_index][current_column]
                if character != " ":
                    vertical_number = vertical_number * 10 + int(character)

            if current_operator == "*":
                vertical_column_result *= vertical_number
            else:
                vertical_column_result += vertical_number

        total_answer_sum += vertical_column_result

    return total_answer_sum


if __name__ == "__main__":
    result = individual_grand_answers_totals()
    print(f"Part 2: {result}")
