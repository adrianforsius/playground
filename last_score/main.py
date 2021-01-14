num = 0
blocks = [5, -2, 4, "Z", "X", 9, "+", "+"]
blocks_modified = [5, -2, 4, "Z", "X", 9, "+", "+"]

last_scores = []
for i, sym in enumerate(blocks):
    if sym == "+":
        vals = []
        if len(last_scores) >= 2:
            last_scores.append(last_scores[-1] + last_scores[-2])
        continue
    if sym == "X":
        if len(last_scores) >= 1 and last_scores[-1]:
            last_scores.append(int(last_scores[-1] * 2))
        continue
    if sym == "Z":
        last_scores.pop()
        continue
    last_scores.append(int(sym))

print(last_scores)
print(sum(last_scores))
