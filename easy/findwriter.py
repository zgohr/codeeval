import sys


if __name__ == '__main__':
    f = open(sys.argv[1])
    for l in f:
        code, key = l.split('|')
        name = []
        for k in [int(x) - 1 for x in key.split()]:
            name.append(code[k])
        print ''.join(name)
