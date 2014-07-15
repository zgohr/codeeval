import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        s, chars = l.split(', ')
        for c in chars.strip():
            s = s.replace(c, '').strip()
        print s
