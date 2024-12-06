## Notes on Types in Go

* Go is statically typed, with built in and user defined types.

``` go
// types of declarations
type Person struct {
    FirstName   string
    LastName    string
    Age         int
}

type CustomInt  int
type Convertor  func(string) Score
type TeamScores map[string]Score
```

### Method vs Functions

* Methods are different from functions varying only in the aspect of receivers.
* No function/method overloading is supported

### Pointer vs Value Receivers

* The method receivers can be pointer or value type.

Rules:

* The method modifies the receiver, we must use a pointer receiver.
* If method needs to handle nil instances.
* If method doesn't modify the receiver, use value receiver.

***Note: The receiver being pointer or value depends on other methods. Its generally a good practice
to have consistency, if any method
has a pointer receiver then all methods of this particular type can have pointer receivers even if
they don't modify the receiver***

* When we use a pointer receiver with a local variable which is value type, go automatically
  converts them to a pointer type. Indicated the `cmd/main.go`, `Counter` example.
* This is untrue for functions, if a value type is passed to a function and if a pointer receiver is
  called, we are invoking the method only the ***copy***.

### Getters and Setters

* Do not right getters and setters for structs, unless you need to meet them to a interface.
* Go encourages to directly access the fields, reserve methods for business logic.
* Exceptions are when you want to update a set of fields in one go or when the update is not a
  straight forward assignment to a field.

### Code Method For NIL Instances

* When we call method with the pointer receiver being nil, what happens?.
    * Go allows us to do this unlike other languages.
    * For value type receivers Go panics. (As there is no value being pointed by the pointer.)
* In some cases accepting nil values makes the code simpler. Eg.., Binary Tree, (Create a binary
  tree directly with inserts.)

### Method Values and Expressions

``` go
type Adder struct {
    start   int
}

func (a *Adder) AddTo(val int) int {
    return a.start + val
}

func main() {
    a := Adder{10}
    a.AddTo(20) // output 30
    
    fa := a.AddTo
    fa(20) // output 30 (method value)
    
    fae := Adder.AddTo
    fae(Adder, 20) // output 30 (method exp)
}
// Method expressions are particularly useful in dependency injections
```

### When to use Methods and Functions?

Since methods can be converted functions, when should one choose between the two?

* Package level state should be immutable.
* Use function when the logic depends only the input parameters.
* Any time your logic depends on values that are configured at startup or changed while your program
  is running, those values should be stored in a struct and that logic should be implemented as a
  method

### Type Declarations Aren't Inheritance

* Most language term subtyping as inheritance. (State and methods of parent type are made available
  for the child type, and the child type can be substituted for the parent type.)
* In Go this is not the case. There is not hierarchy between the two types. In the below example
  child types cannot be substituted in place of the parent types.

``` go
type HighScore Score
type Employee Person 

var i int = 100
var s Score = 200
var hs HighScore = 300
hs = s // compilation error
s = i // compilation error

s = Score(i) // ok 
hs = HighScore(s) // ok
```

### When to Declare User Defined Types on Other Types (Built in/User Defined)

* Types are kind of documentation of your code.
* Makes it clearer for someone reading the code.
* Assume we have method which accepts a percentage, instead of passing on a int, if we have type
  percentage, the probability of making an error is less.
* When we have the same underlying data but require different operations to be done, prefer two
  separate types here.

### Favour Composition Over Inheritance

* Go doesn't have support for inheritance. Composition is built in.
* Any field or method declared on the embedded struct are promoted to the containing struct and can
  be invoked directly on it.
* Any type can be embedded (struct/user defined types/function type)
* If the containing type has the same name as the embedded type. If a field X is present in Inner
  and Outer. access X like this o.x, o.Inner.X

``` go
// example
type Employee struct {
  Id  string
  Name  string 
}

func (e *Employee) Description() string {
  return fmt.SPrintf("%s, %s", e.Id, e.Name)
}

type Manager struct {
  Employee
  Reportees []Employee
}

func main() {
  m := Manager{
    Employee: Employee{
      Id: "12341",
      Name: "FOO",
      },
      Reportees: []Employee{}      
    }
  
  // following are allowed because of embeeding Employee in Manager struct.
  fmt.Println(m.Id)
  fmt.Println(m.Description())
```

### Embedding is Not Inheritance

* No other programming languages support embedding.
* In the above eg.., Manager type cannot be assigned to Employee.
* If you need to access the employee field perform (m.Employee). var e Employee = m // compilation
  error
* Dynamic dispatch is not supported in concrete types in Go.
    * If you have a method in embedded field, which calls another method in the embedded fields.
      Assume that the second method called has the same name as the method in the containing type.
    * The embedded field will not invoke the containing types method.

### The Real Star in Go is its Interface Handling (Btr than Concurrency)

* interfaces in go are usually named with `er` ending. (Eg.., io.Reader, io.Closer, json.Marshaller)

``` go 
type Stringer interface {
  String() string
}
```

* Go's interfaces are implemented implicitly.
* Any concrete type which implements the methods sets of an interface, implements the interface
  implicitly
* The implicit behaviour bring both type safety and decoupling. Bridging both static and dynamically
  typed languages.

### Interfaces are Type safe duck typing

* In the dynamically typed programming language world, (ruby, python, js) don't have interfaces.
* They use duck-typing. If it walks like a duck, quacks like a duck. It is a duck.
* The concept is you have pass an instance of a type to a method, as long as the function can invoke
  the method it expects.

``` ruby 
class Logic:

def process(self, data):
  # business logic

def program(logic):
  logic.process(data)

logicToUse = Logic()
program(logicToUse)
```

* Duck typing makes it hard to know exactly what functionality is expected. As new developers move
  to a project or existing dev forgets about the code. They have navigate to understand what the
  code does.

Java Developers, define an explicit interface and an implementation and refer only to the interface
in the client code.

``` java
public interface Logic {
String process(String data);
}
public class LogicImpl implements Logic { public String process(String data) {
            // business logic
} }
public class Client {
private final Logic logic;
// this type is the interface, not the implementation
public Client(Logic logic) { this.logic = logic;
}
public void program() {
// get data from somewhere this.logic(data);
} }
public static void main(String[] args) { Logic logic = new LogicImpl(); Client client = new Client(logic); client.program();
}
```

* Dynamic language developers look at the explicit interfaces in Java and don’t see how you can
  possibly refactor your code over time when you have explicit dependencies. Switching to a new
  implementation from a different provider means rewriting your code to depend on a new interface.

* Go’s developers decided that both groups are right. If your application is going to grow and
  change
  over time, you need flexibility to change implementation. However, in order for people to
  understand
  what your code is doing (as new people work on the same code over time), you also need to specify
  what the code depends on. That’s where implicit interfaces come in. Go code is a blend of the
  previous two styles:

``` go
type Logic interface {
  Process(data string) string
}

type LogicProvider struct {}

func (l LogicProvider) Process(data string) string {
  // business logic
}

type client struct {
  L Logic
}

func (c Client) Program() {
  c.L.Process(data)
}

func main() {
  c := Client {
      L: LogicProvider{},
    }
  c.Program()  
}
```

### Using Standard Interfaces Encourage Decorator Pattern

``` go
func Process(r io.Reader) error {}

f, err := os.Open("foo.txt") // os.File returned implementes io.Reader interfaces Read Method
if err != nil {
  return err
}
defer f.Close()
return process(f) // os.File type is sent in place of io.Reader
return nil

f, err := os.Open("foo.zip")
if err != nil {
  return err
}
defer f.Close()
gz, err := gzip.NewReader(f) // io.File type being passed inplace of io.Reader.
if err != nil {
  return err
}
defer gz.Close()
return process(gz)

// same code that was reading from a uncompressed file is reading from a compressed file.
```

* Decorator pattern is typically used to add supplementary functionalities to exposed
  interface/functions from libraries.
* It is common in Go to write factory functions, which takes in an instance of an interface &
  returns another type that implements the interface.
