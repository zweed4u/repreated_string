#!/usr/bin/python3


def repeatedString(s, n):
    if s == "a":
        return n
    if "a" not in s:
        return 0
    length_of_string = len(s)
    a_count = 0
    for letter in s:
        if letter == "a":
            a_count += 1

    # handle multiplier
    a_count *= int(n / length_of_string)

    leftover = n % length_of_string
    remaining = s[:leftover]
    for letter in remaining:
        if letter == "a":
            a_count += 1
    print(a_count)
    return a_count


repeatedString("a", 1000000000000)
repeatedString("aba", 10)
repeatedString("abcac", 10)
