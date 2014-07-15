import sys


f = open(sys.argv[1])
for l in f:
    a, b = l.split(',')
    b = b.strip()
    idx = a.find(b)
    print 1 if idx != -1 and idx + len(b) == len(a) else 0
