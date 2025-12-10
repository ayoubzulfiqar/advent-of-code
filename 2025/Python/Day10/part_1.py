from collections import deque


def main():
    try:
        with open("input.txt", "r", encoding="utf-8") as file:
            total_steps = 0

            for line in file:
                line = line.strip()
                if not line:
                    continue

                parts = line.split()
                if len(parts) < 2:
                    continue

                lights_part = parts[0]
                if (
                    len(lights_part) < 2
                    or lights_part[0] != "["
                    or lights_part[-1] != "]"
                ):
                    continue

                lights_str = lights_part[1:-1]
                n = len(lights_str)

                wiring_parts = parts[1:-1]

                start_state = 0
                for i, char in enumerate(lights_str):
                    if char == "#":
                        start_state |= 1 << i

                if start_state == 0:
                    continue

                buttons = []
                for wp in wiring_parts:
                    if len(wp) < 2 or wp[0] != "(" or wp[-1] != ")":
                        continue

                    inner = wp[1:-1]
                    if inner == "":
                        buttons.append(0)
                        continue

                    indices = inner.split(",")
                    mask = 0
                    for idx_str in indices:
                        idx_str = idx_str.strip()
                        if idx_str:
                            try:
                                idx = int(idx_str)
                                if 0 <= idx < n:
                                    mask |= 1 << idx
                            except ValueError:
                                pass
                    buttons.append(mask)

                visited = {start_state}
                queue = deque([start_state])
                steps = 0
                found = False

                while not found and queue:
                    level_size = len(queue)

                    for _ in range(level_size):
                        state = queue.popleft()

                        for btn in buttons:
                            new_state = state ^ btn

                            if new_state == 0:
                                steps += 1
                                found = True
                                break

                            if new_state not in visited:
                                visited.add(new_state)
                                queue.append(new_state)

                        if found:
                            break

                    if found:
                        break

                    steps += 1

                total_steps += steps

        print(total_steps)

    except FileNotFoundError:
        print("Error opening file: file not found")
    except Exception as e:
        print(f"Error: {e}")


if __name__ == "__main__":
    main()
