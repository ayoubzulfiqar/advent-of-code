# Read the input file and process it
with open("input.txt", "r") as file:
    input_data = [line.strip() for line in file.readlines()]

rows = len(input_data)
cols = len(input_data[0])

# Direction vectors (up, right, down, left)
directions = [
    (-1, 0),  # Up
    (0, 1),  # Right
    (1, 0),  # Down
    (0, -1),  # Left
]

# Initialize variables for the guard's starting position and direction
guard_row, guard_col, guard_dir = None, None, None
for r in range(rows):
    for c in range(cols):
        if input_data[r][c] in "^>v<":
            guard_row, guard_col = r, c
            guard_dir = "^>v<".index(input_data[r][c])  # Map direction symbol to index
            break

# Set to track distinct visited positions
visited = set()
visited.add(f"{guard_row},{guard_col}")

# Simulate guard movement
while True:
    dr, dc = directions[guard_dir]
    next_row, next_col = guard_row + dr, guard_col + dc

    # Check if the next position is outside the grid
    if next_row < 0 or next_row >= rows or next_col < 0 or next_col >= cols:
        break  # Guard leaves the grid

    if input_data[next_row][next_col] == "#":
        # Obstacle ahead: turn right (clockwise)
        guard_dir = (guard_dir + 1) % 4
    else:
        # Move forward
        guard_row, guard_col = next_row, next_col
        visited.add(f"{guard_row},{guard_col}")

# Output the number of distinct positions visited
print(f"Distinct positions visited: {len(visited)}")
