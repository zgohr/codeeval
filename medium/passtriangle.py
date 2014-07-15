import sys


triangle = []
for l in open(sys.argv[1]):
    triangle.append([int(x) for x in l.split()])

# Starting at the bottom row, move our way up
for ridx in range(len(triangle) - 1, 0, -1):
    # Starting at the final column, move our way up
    for cidx in range(len(triangle[ridx]) - 1, 0, -1):
        # Set the next row's next column total (since we're travling backward)
        triangle[ridx - 1][cidx - 1] += max(triangle[ridx][cidx],
                                            triangle[ridx][cidx - 1])
print triangle[0][0]
