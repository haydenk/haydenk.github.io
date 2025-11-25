+++
title = 'flexible_fizz_buzz.scala'
date = 2025-04-27T14:01:01Z
type = "snippet"
tags = ['scala']
+++
<!--more-->

```scala
def flexibleFizzBuzz(start: Int, end: Int) (callback: String => Unit): Unit = {
    for (i <- Range(start, end)) {
      var output: String = ""
      if (i % 3 == 0) {
        output += "Fizz"
      }
      if (i % 5 == 0) {
        output += "Buzz"
      }
      if (output.isEmpty) {
        output = s"${i}"
      }
      callback(output)
    }
}

flexibleFizzBuzz(start = 1, end = 50) {
  i => println(i)
}
```

**Output:**

```text
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
```
