def all_same(arr):
	return len(set(arr)) == 1

def is_sumable(arr, target):
	data = {}
	
	for i, v in enumerate(arr):
		if data.get(v) is None:
			data[target - v] = i
		else:
			return True
			
	return False
	
def nonConstructibleChange(coins):
    # Write your code here.
	if (len(coins) == 0):
		return 1
	
	if all_same(coins):
		return (len(coins) * coins[0]) + 1
	
	coins.sort()
	
	if coins[0] != 1:
		return 1

	for target in range(1, coins[len(coins) - 1]):
		if target == 1:
			continue

		if is_sumable(coins, target) == False:
			return target


print(nonConstructibleChange([5, 7, 1, 1, 2, 3, 22]))