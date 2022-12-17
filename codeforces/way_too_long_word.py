n = input()
strs = []

for i in ragen(n):
	strs.append(input())

for s in strs:
	if len(s) > n:
		fc = s[0]
		lc = s[-1]
		c = s[1:-1]
		print(fc + len(c) + lc)
	else:
		print(s)
