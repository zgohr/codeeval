import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        l = l.rstrip()
        for i, x in enumerate(l):
            if l.count(str(i)) != int(x):
                print 0
                break
        else:
            print 1
