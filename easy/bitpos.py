import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        n, p1, p2 = list(int(b) for b in l.split(','))
        b = "{0:b}".format(n)
        print str(b[-1 * p1] == b[-1 * p2]).lower()
