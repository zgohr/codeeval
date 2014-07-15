import sys

f = open(sys.argv[1])
for l in f:
    s = l.rstrip().split(' ')
    # Maximum repeatable length is half the number of values
    for size in range(1, len(s) / 2 + 1):
        # Loop over all values
        for i in range(0, len(s)):
            # Does the current slice match the previous
            curr = s[i:i + size]
            prev = s[i - size:i]
            if curr and prev and curr == prev:
                print ' '.join(s[i:i + size])
                break
        else:
            continue
        break
