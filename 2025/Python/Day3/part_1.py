def output_joltage():
    try:
        with open("input.txt", "r") as file:
            total = 0

            for line in file:
                line = line.strip()
                if not line:
                    continue

                length = 0
                for i, char in enumerate(line):
                    if "0" <= char <= "9":
                        length += 1
                    else:
                        break

                T = [0] * length
                R = [0] * length

                for i in range(length):
                    T[i] = int(line[i])

                summands = 2
                for _ in range(summands):
                    m = 0
                    for i in range(length):
                        new_val = 10 * m + T[i]
                        if R[i] > m:
                            m = R[i]
                        R[i] = new_val

                m = 0
                for i in range(length):
                    if R[i] > m:
                        m = R[i]
                    R[i] = m

                total += R[length - 1]

        return total

    except FileNotFoundError:
        print("Error: File 'input.txt' not found")
        exit(1)
    except Exception as e:
        print(f"Error reading file: {e}")
        exit(1)


if __name__ == "__main__":
    result = output_joltage()
    print(result)
