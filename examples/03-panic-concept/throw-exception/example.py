# -*- coding: utf-8 -*-

def div(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b


print("div(42, 0) = %d" % div(42, 0))
print("OK")  # Не напечатается!
