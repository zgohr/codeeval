import sys

for l in open(sys.argv[1]):
    l = l.strip()
    nums, s = l.split(';')
    nums = [int(x) for x in nums.split(',')]
    pairs = []
    for i, n in enumerate(nums):
        for n2 in nums[i + 1:]:
            if n + n2 == int(s):
                pairs.append("%d,%d" % (n, n2))

    print ';'.join(pairs) if len(pairs) > 0 else "NULL"
