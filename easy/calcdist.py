import sys
import re
import math


if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        s = [int(x) for x in re.sub(r'\(|\)|\,', '', l).split()]
        print int(math.sqrt((math.pow(s[2] - s[0], 2)) +
                            (math.pow(s[3] - s[1], 2))))
