import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        a, b, n = list(int(x) for x in l.split(" "))
        out = []
        for x in range(1, n + 1):
            out.append('FB' if x % a is 0 and x % b is 0
                       else 'F' if x % a is 0
                       else 'B' if x % b is 0
                       else str(x))
        print ' '.join(out)
