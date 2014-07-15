import sys

if __name__ == '__main__':
    f = open(sys.argv[1])
    s = 0
    for l in f:
        s += int(l.rstrip())
    print s
