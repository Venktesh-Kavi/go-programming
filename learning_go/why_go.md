## Why Go

### Personal Standpoint

* Wanted to learn a programming language to do quick prototypes, understand low level system internals (like networks, system programming, databases)
* Have been a java programmer for large part of the career. Though I have an inclination on compiled and staticly typed languages, java being centered around object oriented programming and making it excessively verbose and time consuming to accomplanish simple tasks (sample can be simple json parsing).
* Wanted to learn a language which feels like an interpreted language providing all the code sugars and can help me achive tasks in a simple way.


### Technically Why Go?

#### Compiled and Interpreted languages

* Examples of compiled languages are java, go
* Intepreted languages (python, ruby)
* Most interpreted languages are dynamically types (duck typed).

#### What does compiled mean?
* All of the code is compiled into an intermediary/machine code before it is executed.
* In Java the java compiler (javac) compiles the code to byte code. The JVM and JIT (Just in Time interpreter) interpret the code to machine code (JIT caches most frequently used code to direct machine code)
* In Go there is no intermediary step of byte code conversion, the compiled code is in machine code. This is were go gains most of its performance.

#### What does interpreted language mean?\

* Interpreted languaages are read line by line by the interpreted which is usally written in a low-level language (eg.., python uses CPython as the interpretor, ruby uses MRI (Matz Ruby Interpreter written in C)
* Intepreted languages tend to be slow as there is an intermediary interpretation layer required.

#### Concurrency & Parallelism Handling in Compiled and Interpreted Languages

* Concurrency is the feel of multiple tasks running simaltaneously. (The reason its a feel, is because the CPU context switches depending on the scheduler).
* Parallelism is diving a task into smaller units and acting on it parallely. (parallelism can be achieved only if have multiple cores).

* Large part of the 20 century languages were developed in a time where there only single core processors. Concurrency and distributing tasks were a luxury.
* Over the period of time as we entered the 21st century the chip development was more skewed towards increasing the number of cores rather than building a single CPU with a large core (transistors) (This could be due to clock frequencies, thermal desipation and other factors).
* For most the languages in the 20th century concurrency/parallelism was a after thought.
* Enter Go (21st century centered around ease of concurrent programming)


Java
* Java programs run as a single process. Java achieves concurreny and parallelism via threads and higher level constructs like executor service.
* If the machine has a single core, then java threads use context switching to achieve concurrency. In a multi core system this is called parallelism.
* Java threads are native OS threads, so they tend to occupy higher memory compared to green thread (virtual threads introduced in java 21)

Python/Ruby

* Python and Ruby also have threads.
* Both these interpreted languages suffer lack of parallelism due to GIL (Global Interpreter Lock)
* When we mean lack of parallelism, this means CPU bound tasks cannot be parallelism, IO tasks can be parallelised.

```
# IO based task example. Here we sleep on the thread, so control is delegated to another thread to achieve parallelism
# Define a method to be executed in a separate thread
def do_work
  puts "Thread #{Thread.current.object_id} started."
  sleep(2) # Simulate some work
  puts "Thread #{Thread.current.object_id} finished."
end

# Create three threads and start them
threads = []
3.times do
  threads << Thread.new { do_work }
end

# Wait for all threads to complete
threads.each(&:join)

puts "All threads have completed."
```

```
## CPU Bound Task Example, GIL ensures that only one ruby thread can execute ruby bytecode at a time.
def multiply_numbers(a, b)
  a * b
end

def run_multiplications_in_threads(numbers)
  results = {}
  threads = []

  numbers.each do |pair|
    a, b = pair
    threads << Thread.new do
      results[[a, b]] = multiply_numbers(a, b)
    end
  end

  threads.each(&:join)
  results
end

# Example usage
numbers = [[2, 3], [5, 7], [11, 13], [17, 19]]
start_time = Time.now
results = run_multiplications_in_threads(numbers)
end_time = Time.now

puts "Results:"
results.each { |pair, result| puts "#{pair[0]} * #{pair[1]} = #{result}" }
puts "Total time: #{(end_time - start_time).round(2)} seconds"

````
* This would the reason why in puma server in ruby on rails, we setup the number of processes. We run N number of processes to utilise the multiple cores in the CPU, but each process is still bound to the GIL.

