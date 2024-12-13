def fewestTokenToWin() -> int:
    # Parse the constants out of each test
    with open("./input.txt") as file:
        con = file.read()
    machines = []
    for block in con.strip().split("\n\n"):
        button_a, button_b, prize = block.split("\n")
        c1, c4 = [int(num.split("+")[1]) for num in button_a.split(": ")[1].split(", ")]
        c2, c5 = [int(num.split("+")[1]) for num in button_b.split(": ")[1].split(", ")]
        c3, c6 = [int(num.split("=")[1]) for num in prize.split(": ")[1].split(", ")]
        machines.append({"c1": c1, "c2": c2, "c3": c3, "c4": c4, "c5": c5, "c6": c6})

    def calculateSum() -> int:
        total= 0
        for _, machine in enumerate(machines):
            c1, c2, c3, c4, c5, c6 = (
                machine["c1"],
                machine["c2"],
                machine["c3"],
                machine["c4"],
                machine["c5"],
                machine["c6"],
            )

            # Solve for a and b
            b = (c1 * c6 - c4 * c3) / (c1 * c5 - c4 * c2)
            a = (c3 - c2 * b) / c1

            # Check if a and b are integers
            if a.is_integer() and b.is_integer():
                total += a * 3 + b
        print(int(total))
        return total

    return calculateSum()


