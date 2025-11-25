+++
title = 'recursive_messages.scala'
date = 2025-04-27T14:01:21Z
type = "snippet"
tags = ['scala']
+++
<!--more-->

```scala
class Msg(val id: Int, val parent: Option[Int], val text: String)
def printMessages(messages: Array[Msg]): Unit = {
  for (message <- messages) {
    var indent: Int = 0
    message.parent match {
      case Some(parent) => indent = parent + 2
      case None => indent = 0
    }
    println(" " * indent + s"#${message.id} ${message.text}")
  }
}

printMessages(Array[Msg](
  new Msg(0, None, "Hello"),
  new Msg(1, Some(0), "World"),
  new Msg(2, None, "I am Cow"),
  new Msg(3, Some(2), "Here me moo"),
  new Msg(4, Some(2), "Here I stand"),
  new Msg(5, Some(2), "I am Cow"),
  new Msg(6, Some(5), "Here me moo, moo"),
))
```

**Output:**

```text
#0 Hello
  #1 World
#2 I am Cow
    #3 Here me moo
    #4 Here I stand
    #5 I am Cow
       #6 Here me moo, moo
```
