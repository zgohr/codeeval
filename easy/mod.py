import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        n, m = list(int(x) for x in l.rstrip().split(','))
        print n - ((n / m) * m)
