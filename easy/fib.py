import sys
import math


def fib(n):
    phi = ((1 + math.sqrt(5)) / 2.0)
    return (math.pow(phi, n) - (math.pow(-1, n) / math.pow(phi, n))) /\
        math.sqrt(5)

if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        n = int(l)
        print int(fib(n))
