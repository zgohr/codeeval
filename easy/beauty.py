import sys
import re
from operator import itemgetter

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        s = re.sub(r'\W+', '', l.rstrip().lower())
        max_val = 26
        total = 0
        value = {}
        for x in set(s):
            value[x] = s.count(x)
        for k, v in sorted(value.items(), key=itemgetter(1), reverse=True):
            total += v * max_val
            max_val -= 1
        print total
