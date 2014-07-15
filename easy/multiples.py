import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        x, n = list(int(b) for b in l.split(','))
        i = count = 0
        while(i < x):
            count += 1
            i = n * count
        print i
