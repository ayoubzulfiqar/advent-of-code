import 'dart:io';
import 'dart:math';

class Vertex {
  final int coordX;
  final int coordY;
  final int coordZ;

  Vertex(this.coordX, this.coordY, this.coordZ);

  double separation(Vertex w) {
    final deltaX = (coordX - w.coordX).toDouble();
    final deltaY = (coordY - w.coordY).toDouble();
    final deltaZ = (coordZ - w.coordZ).toDouble();
    return sqrt(deltaX * deltaX + deltaY * deltaY + deltaZ * deltaZ);
  }
}

class EdgeInfo {
  final int idxA;
  final int idxB;
  final double distance;

  EdgeInfo(this.idxA, this.idxB, this.distance);
}

class UnionStructure {
  final List<int> ancestor;
  final List<int> level;
  int groups;

  UnionStructure(int total)
    : ancestor = List<int>.generate(total, (i) => i),
      level = List<int>.filled(total, 0),
      groups = total;

  int locate(int pos) {
    if (ancestor[pos] != pos) {
      ancestor[pos] = locate(ancestor[pos]);
    }
    return ancestor[pos];
  }

  bool combine(int x, int y) {
    int rootX = locate(x);
    int rootY = locate(y);

    if (rootX == rootY) {
      return false;
    }

    if (level[rootX] < level[rootY]) {
      int temp = rootX;
      rootX = rootY;
      rootY = temp;
    }

    ancestor[rootY] = rootX;
    if (level[rootX] == level[rootY]) {
      level[rootX]++;
    }

    groups--;
    return true;
  }
}

void xCoordinatesJuntionMultiplicationBoxs() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    final locations = <Vertex>[];
    for (final textRow in lines) {
      if (textRow.isEmpty) continue;

      final elements = textRow.split(',');
      if (elements.length != 3) continue;

      final node = Vertex(
        int.parse(elements[0].trim()),
        int.parse(elements[1].trim()),
        int.parse(elements[2].trim()),
      );
      locations.add(node);
    }

    final connections = <EdgeInfo>[];
    final vertexCount = locations.length;

    for (int firstIdx = 0; firstIdx < vertexCount; firstIdx++) {
      for (int secondIdx = firstIdx + 1; secondIdx < vertexCount; secondIdx++) {
        final span = locations[firstIdx].separation(locations[secondIdx]);
        connections.add(EdgeInfo(firstIdx, secondIdx, span));
      }
    }

    connections.sort((a, b) => a.distance.compareTo(b.distance));

    final unionSet = UnionStructure(vertexCount);

    for (final link in connections) {
      if (unionSet.combine(link.idxA, link.idxB)) {
        if (unionSet.groups == 1) {
          final finalProduct =
              locations[link.idxA].coordX * locations[link.idxB].coordX;
          print(finalProduct);
          return;
        }
      }
    }
  } on FileSystemException catch (openErr) {
    print('Cannot access data file: $openErr');
  } on FormatException catch (parseErr) {
    print('Failed to convert value: $parseErr');
  } catch (readErr) {
    print('Error during file reading: $readErr');
  }
}

void main() {
  xCoordinatesJuntionMultiplicationBoxs();
}
