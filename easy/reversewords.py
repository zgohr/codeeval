import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        l = l.rstrip()
        if l == '':
            continue
        print ' '.join(l.split(' ')[::-1])
