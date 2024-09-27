## Golang Docs

**💥1. Getting Started ⇒**

1️⃣ Create a file **main.go** and open it in integrated terminal by just doing right click and press "Open in Integrated terminal".

2️⃣ Now run command in terminal **"go mod init example/hello"**

3️⃣ Now type in **main.go** file our first "Hello World" program.

    package main

    import "fmt"

    func main( ) {
    	fmt.Println("Hello World!");
    }

4️⃣ Like other languages C++, java in go this is our main function:

      func main( ) {
    	fmt.Println("Hello World!");
    }

5️⃣ Now to run the file we just need to give command **"go run fileName"**.

6️⃣ Command we use to find the go environment **"go env GOPATH"**

### 💥Data Types in Go

---

### 1. Basic Data Types

**1️⃣ Integers ⇒**

    1. Signed Integers : int8, int16, int32, int64

    2. Unsigned Integers : unit8(byte), unit16, unit32, unit64, unit

**2️⃣ Floating :** float32, float64(default is float64)

**3️⃣ Boolean :** bool (true or false)

**4️⃣ Strings :** string (seq. of characters, immutable)

**5️⃣ Complex :** complex64, complex128

### 2. Derived Data Types

**1️⃣ Array :** Fixed size collection of elements of same type.

    Example: var arr[5] int

**2️⃣ Slice :** Dynamically-Sized array, more flexible than array.

    Example: var slice[ ] int

**3️⃣ Map :** Key-value pair collections, like dictionaries.

    Example: map[string] int

**4️⃣ Struct :** Collection of fields of different data types. In other language we say object.

    Example:

    type Person struct {

        Name string

        Age int
    }

### 3. Pointer

Stores the memory address of a variable.

    Example: var ptr *int

### 4. Function

Functions themselves are data types in Go, and they can be assigned to variables.

### 5. Interface

An abstract type used to define a set of methods.

    Example:

    type Shape interface {

      Area() float64

    }

### 6. Channel

A mechanism for goroutines to communicate with each other.

    Example: chan int

**💥 Declaring variables ⇒** To declare variable in go we use **var keyword**. It is also known as explicit type declaration.

    Example:

    var name string = "Harsh"

    var age int = 32

Even go allow us to declare a variable without specifying the type.

    Example:

    var name = "harsh"

    var age = 32

We can also use short declaration **(:=)** to declare variable.

    Example:

    name := "Harsh"

    age := 25

**💥 Checking data type of any variable :** To check data type in go we use format specifier **"%T"**.

    Example:

    func main() {

      var username string = "Harsh Mehra"

      fmt.Println(username);

      fmt.Printf("Variable is of type: %T \n", username);
    }

    output :Harsh Mehra
            Variable is of type: string

**💥 Variable name in Capital :** Starting global variable name with **capital letter** means this variable can be accessed in another packages. This is similar to **public** we see in other languages.

    Example:

    var ExportedPackage string = "I am visible outside my package"
