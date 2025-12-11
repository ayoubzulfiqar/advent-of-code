import 'dart:io';

class PathState {
  final String node;
  final int mask; 

  PathState(this.node, this.mask);

  @override
  bool operator ==(Object other) {
    return other is PathState && other.node == node && other.mask == mask;
  }

  @override
  int get hashCode => Object.hash(node, mask);
}

int dfsBitMask(Map<String, List<String>> adj, String start) {
  final memo = <PathState, int>{};

  int dfs(String curr, int mask) {
    if (curr == "out") {
      return mask == 0x03 ? 1 : 0; // Both bits set (fft AND dac)
    }

    final state = PathState(curr, mask);
    if (memo.containsKey(state)) {
      return memo[state]!;
    }

    int ans = 0;
    final children = adj[curr] ?? [];

    for (final child in children) {
      int newMask = mask;
      switch (child) {
        case "fft":
          newMask |= 0x01; // Set fft bit
          break;
        case "dac":
          newMask |= 0x02; // Set dac bit
          break;
      }
      ans += dfs(child, newMask);
    }

    memo[state] = ans;
    return ans;
  }

  return dfs(start, 0);
}

int dacAndFftPath() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    final adj = <String, List<String>>{};

    for (final line in lines) {
      final parts = line.trim().split(RegExp(r'\s+'));
      if (parts.length < 2) {
        continue;
      }

      final key = parts[0].endsWith(':')
          ? parts[0].substring(0, parts[0].length - 1)
          : parts[0];

      final values = parts.sublist(1);
      adj[key] = values;
    }

    final result = dfsBitMask(adj, "svr");
    print(result);
    return result;
  } catch (e) {
    print("Error opening file: $e");
  }
  return -1;
}

void main() {
  dacAndFftPath();
}


/* 

import 'dart:io';

class IterativeBitmaskSolver {
  final Map<String, List<String>> _adjacency;
  
  IterativeBitmaskSolver(this._adjacency);
  
  int countPaths(String startNode) {
    if (startNode == "out") return 0;
    
    final memo = <String, Map<int, int>>{};
    final stack = <_StackFrame>[];
    
    // Initialize memo for "out" node
    memo["out"] = {0: 0, 1: 0, 2: 0, 3: 1}; // Only mask=3 (0b11) gives 1
    
    // Start DFS
    stack.add(_StackFrame(startNode, 0, false));
    
    while (stack.isNotEmpty) {
      final frame = stack.removeLast();
      
      if (frame.isProcessed) {
        // Compute result for this node
        int total = 0;
        final children = _adjacency[frame.node] ?? [];
        
        for (final child in children) {
          int newMask = frame.mask;
          if (child == "fft") newMask |= 0x01;
          if (child == "dac") newMask |= 0x02;
          
          total += memo[child]![newMask]!;
        }
        
        if (!memo.containsKey(frame.node)) {
          memo[frame.node] = {};
        }
        memo[frame.node]![frame.mask] = total;
        
      } else {
        // Check if already computed
        if (memo.containsKey(frame.node) && 
            memo[frame.node]!.containsKey(frame.mask)) {
          continue;
        }
        
        // Mark as processing and push children
        stack.add(_StackFrame(frame.node, frame.mask, true));
        
        final children = _adjacency[frame.node] ?? [];
        for (final child in children) {
          int newMask = frame.mask;
          if (child == "fft") newMask |= 0x01;
          if (child == "dac") newMask |= 0x02;
          
          // Ensure memo entry exists for child
          if (!memo.containsKey(child)) {
            memo[child] = {};
          }
          if (!memo[child]!.containsKey(newMask)) {
            stack.add(_StackFrame(child, newMask, false));
          }
        }
      }
    }
    
    return memo[startNode]?[0] ?? 0;
  }
}

class _StackFrame {
  final String node;
  final int mask;
  final bool isProcessed;
  
  _StackFrame(this.node, this.mask, this.isProcessed);
}

void dacAndFftPathIterative() {
  final stopwatch = Stopwatch()..start();
  
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();
    
    final adjacencyList = <String, List<String>>{};
    
    for (final line in lines) {
      final parts = line.trim().split(RegExp(r'\s+'));
      if (parts.length < 2) continue;
      
      String key = parts[0];
      if (key.endsWith(':')) {
        key = key.substring(0, key.length - 1);
      }
      
      adjacencyList[key] = parts.sublist(1);
    }
    
    final solver = IterativeBitmaskSolver(adjacencyList);
    final result = solver.countPaths("svr");
    
    print("Result: $result");
    
    stopwatch.stop();
    print("Time: ${stopwatch.elapsedMilliseconds} ms");
    
  } catch (e) {
    print("Error: $e");
  }
}

void main() {
  dacAndFftPathIterative();
}


*/