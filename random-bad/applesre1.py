
duplicates = ['a','b','c','d','d','d','e','a','b','f','g','g','h']
# new_set = set(duplicates)
# duplicates = list(new_set)
# print(duplicates.sort())

d = {}

for letter in duplicates:
    if letter not in d:
        d[letter] = 1
    else:
        d[letter] += 1

for k in d:
    print(k)

# Write a function to which returns an infinite number sequence and print the numbers in main

def infinite_sequence():
    counter = 0
    while True:
        counter += 1
        yield counter


for i in infinite_sequence():
     print(i, end=" ")