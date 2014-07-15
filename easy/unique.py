import sys

if __name__ == "__main__":
    f = open(sys.argv[1])
    for l in f:
        print ','.join(str(x) for x in
                       set((int(x) for x in l.split(','))))
