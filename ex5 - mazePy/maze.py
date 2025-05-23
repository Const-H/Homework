from collections import deque

def load_maze(file_path):
    with open(file_path) as f:
        return [list(line.strip()) for line in f]

def locate_points(maze):
    start = end = None
    for row in range(len(maze)):
        for col in range(len(maze[row])):
            if maze[row][col] == 'S':
                start = (row, col)
            elif maze[row][col] == 'E':
                end = (row, col)
    return start, end

def find_path(maze, start, target):
    q = deque([(start, [start])])
    seen = {start}
    
    while q:
        (i, j), path = q.popleft()
        
        if (i, j) == target:
            return path
            
        for di, dj in [(0,1), (1,0), (0,-1), (-1,0)]:
            ni, nj = i + di, j + dj
            if (0 <= ni < len(maze) and 0 <= nj < len(maze[0])) and maze[ni][nj] != '#' and (ni, nj) not in seen:
                seen.add((ni, nj))
                q.append(((ni, nj), path + [(ni, nj)]))
    return None

def draw_path(maze, path):
    steps = len(path) - 1
    for i, j in path[1:-1]:
        maze[i][j] = '*'
    return steps

def save_result(file_path, maze, steps):
    with open(file_path, 'w') as f:
        f.write(f"Steps: {steps}\n")
        f.write('\n'.join(''.join(row) for row in maze) + '\n')

def process_maze(input_file, output_file="output.txt"):
    maze = load_maze(input_file)
    start, end = locate_points(maze)
    
    if not start or not end:
        with open(output_file, 'w') as f:
            f.write("Start or end point missing.\n")
        return
    
    path = find_path(maze, start, end)
    
    if path:
        steps = draw_path(maze, path)
        save_result(output_file, maze, steps)
    else:
        with open(output_file, 'w') as f:
            f.write("No path found.\n")

if __name__ == "__main__":
    process_maze("input.txt")