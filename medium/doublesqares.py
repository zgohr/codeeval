import sys
import math


def int_sqr(n):
    return int(math.sqrt(n))


def squares(n):
    count = 0
    square_n = int_sqr(n)
    # For each integer, 0 to square
    for i in range(square_n + 1):
        # Starting with n^2 + n^2 to n^2 .. 0^2
        # If when summed, equal N, count it
        y = int_sqr(n - i*i)
        if y*y + i*i == n and i <= y:
            count += 1
    return count


if __name__ == "__main__":
    f = open(sys.argv[1])
    for i, l in enumerate(f):
        if i is 0: continue
        if l.strip() == '': continue
        print squares(int(l))