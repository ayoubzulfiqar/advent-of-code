import math


def xCoordinatesJunctionMultiplicationBoxes():
    vertices = []
    try:
        with open("input.txt", "r", encoding="utf-8") as data_file:
            for text_row in data_file:
                text_row = text_row.strip()
                if not text_row:
                    continue

                elements = text_row.split(",")
                if len(elements) != 3:
                    continue

                node = {
                    "coordX": int(elements[0].strip()),
                    "coordY": int(elements[1].strip()),
                    "coordZ": int(elements[2].strip()),
                }
                vertices.append(node)
    except FileNotFoundError as file_err:
        print(f"Cannot access data file: input.txt {file_err}")
        return
    except ValueError as parse_err:
        print(f"Failed to convert value: {parse_err}")
        return
    except Exception as read_err:
        print(f"Error during file reading: {read_err}")
        return

    def separation(v1, v2):
        delta_x = float(v1["coordX"] - v2["coordX"])
        delta_y = float(v1["coordY"] - v2["coordY"])
        delta_z = float(v1["coordZ"] - v2["coordZ"])
        return math.sqrt(delta_x * delta_x + delta_y * delta_y + delta_z * delta_z)

    connections = []
    vertex_count = len(vertices)

    for first_idx in range(vertex_count):
        for second_idx in range(first_idx + 1, vertex_count):
            span = separation(vertices[first_idx], vertices[second_idx])
            connections.append(
                {"idxA": first_idx, "idxB": second_idx, "distance": span}
            )

    connections.sort(key=lambda conn: conn["distance"])

    class UnionStructure:
        def __init__(self, total):
            self.ancestor = list(range(total))
            self.level = [0] * total
            self.groups = total

        def locate(self, pos):
            if self.ancestor[pos] != pos:
                self.ancestor[pos] = self.locate(self.ancestor[pos])
            return self.ancestor[pos]

        def combine(self, x, y):
            root_x = self.locate(x)
            root_y = self.locate(y)

            if root_x == root_y:
                return False

            if self.level[root_x] < self.level[root_y]:
                root_x, root_y = root_y, root_x

            self.ancestor[root_y] = root_x
            if self.level[root_x] == self.level[root_y]:
                self.level[root_x] += 1

            self.groups -= 1
            return True

    union_set = UnionStructure(vertex_count)

    for link in connections:
        if union_set.combine(link["idxA"], link["idxB"]):
            if union_set.groups == 1:
                final_product = (
                    vertices[link["idxA"]]["coordX"] * vertices[link["idxB"]]["coordX"]
                )
                print(final_product)
                return


if __name__ == "__main__":
    xCoordinatesJunctionMultiplicationBoxes()
