class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0

        count = 0
        db = {}
        self.length_of_longest_substring(s, count, db)

        biggest_number = 0
        for i, _ in db.items():
            if db[i]["count"] > biggest_number:
                biggest_number = db[i]["count"]

        return biggest_number

    def length_of_longest_substring(self, s: str, count=0, db={}):
        if len(s) == 0:
            return

        for c in s:
            if db.get(count) is None:
                db[count] = {"data": {c: True}, "count": 1}
            else:
                if not db[count]["data"].get(c):
                    db[count]["data"].update({c: True})
                    db[count]["count"] += 1
                else:
                    break

        count += 1
        return self.length_of_longest_substring(s[1:], count, db)


print(Solution().lengthOfLongestSubstring("tmmzuxt"))