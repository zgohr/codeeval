import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        l = l.rstrip()
        print ''.join([x.upper() if x.islower() else x.lower() for x in l])
