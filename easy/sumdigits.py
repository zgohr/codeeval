import sys

f = open(sys.argv[1])
for l in f:
    n = l.rstrip()
    print sum(int(x) for x in n)
