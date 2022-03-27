def avg(marks):
    assert len(marks) != 0, 'empty marks list'
    return round(sum(marks) / len(marks), 2)


# python marks.py
# python -O marks.py
marks = []
print(avg([]))
