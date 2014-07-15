import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        if l.rstrip() == "":
            continue
        s, t = l.split(',')
        print s.rfind(t.rstrip())
