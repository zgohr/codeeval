import sys
for l in open(sys.argv[1]):
    print l.split(' ')[-2]
