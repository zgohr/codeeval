import sys


if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        for c in l:
            if l.count(c) is 1:
                print c
                break
