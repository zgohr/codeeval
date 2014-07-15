import sys
import math

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        b = l.rstrip()
        count = 0
        while(count < 1000):
            b = sum([int(math.pow(int(x), 2)) for x in b])
            if b == 1:
                print 1
                break
            b = str(b)
            count += 1
        else:
            print 0
