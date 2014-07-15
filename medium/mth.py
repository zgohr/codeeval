import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        x = l.rstrip().split(' ')
        m = int(x[-1])
        if m < len(x):
            print x[-1 * (m + 1)]
