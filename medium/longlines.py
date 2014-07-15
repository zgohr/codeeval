import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    count = None
    lines = []
    for l in f:
        s = l.rstrip()
        if count is None:
            count = int(s)
        else:
            lines.append(s)
    lines.sort(key=lambda s: len(s), reverse=True)
    print '\n'.join(lines[:count])
