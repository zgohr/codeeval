import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        print l.lower().rstrip()
