import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        x, y = l.rstrip().split(';')
        x = set(int(t) for t in x.split(','))
        y = set(int(t) for t in y.split(','))
        print ','.join(str(t) for t in x.intersection(y))
