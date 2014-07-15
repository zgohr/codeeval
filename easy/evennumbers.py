import sys

if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        print 1 if int(l) % 2 is 0 else 0
