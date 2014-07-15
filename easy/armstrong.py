import sys
import math

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        s = l.rstrip()
        f = float(s)
        l = len(s)
        print sum([math.pow(int(x), l) for x in s]) == f
