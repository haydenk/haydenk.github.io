+++
title = 'file_reader_writer.scala'
date = 2025-04-27T13:59:10Z
type = "snippet"
tags = ['scala']
+++

```scala
import java.io.{BufferedReader, BufferedWriter, FileReader, FileWriter}

def withFileWriter(filename: String) (handler: BufferedWriter => Unit): Unit = {
  val writer = new BufferedWriter(new FileWriter(filename))
  try handler(writer)
  finally writer.close()
}

def withFileReader(filename: String) (handler: BufferedReader => Unit): Unit = {
  val reader = new BufferedReader(new FileReader(filename))
  try handler(reader)
  finally reader.close()
}

withFileWriter("Hello.txt") {
  writer => { writer.write("Hello\n"); writer.write("World!") }
}

var result: String = ""

withFileReader("Hello.txt") {
  reader => { result = reader.readLine() + "\n" + reader.readLine() }
}

assert(result == "Hello\nWorld!")
```
