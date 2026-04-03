# GO VARIABLES
# What are variables?

 Variables are containers for storing data values.

 # SO WE HAVE 4 TYPES OF GO VARIBLE TYPES WHICH ARE:

 int: which are used in storing (numbers such as 2, 4, -4, -8)

 float32: which stores floating point numbers, with decimals, such as 19.99 or -19.99.

 strings: which are used for storing texts such as "hello" string values are always surrounded by double quotes.

 bool: used in storing values in 2 states: true or false.


 # They are various ways of declaring a variable which are:

 If you want to declare the variable you say....

 var<> keyword for declaring a variable.
 Money<> variable name.
 bool <> variable type.

 Putting them together it is
   var Money bool

   The other way of declaring a variable is by using the short variable declaration

   Money <> name of variable
   := <> short variable declaration
   student1 <> value

   SO let's take an example:

   package main

   import ("fmt")

   var student1 bool = "ben"
   var student2 = "ben"
   x := 1

   func main(){
    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)
   }
   The example above declares a variable with an initial value

   So let's declare some variable without actually an initial value to it:

   package main

   import ("fmt")

   func main() {
    var a string
    var b int
    var c bool

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
   }