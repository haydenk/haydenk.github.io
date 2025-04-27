+++
title = 'fibonacci.py'
date = 2025-04-27T13:36:17Z
type = "snippet"
tags = ['python']
+++

```python
def fibonacci(n: int) -> int:
    abs_n: int = abs(n)
    if abs_n < 2:
        return 0
    if abs_n < 3:
        return 1
    return fibonacci(abs_n-1) + fibonacci(abs_n-2)
```
