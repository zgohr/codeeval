import sys
import re
import string

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        s = re.sub(r'\W+', '', l.rstrip().lower())
        others = []
        for x in string.lowercase:
            try:
                s.index(x)
            except ValueError:
                others.append(x)
        if len(others) is 0:
            print 'NULL'
        else:
            print ''.join(others)
