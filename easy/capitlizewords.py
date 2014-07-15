import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        print ' '.join([x[0].capitalize() + x[1:] for x in l.split()])
