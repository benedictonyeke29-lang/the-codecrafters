# GO CONSTANT

A Go const is only used when the variable declared has a fixed and unchangeable value that is when the const keyword is used the variable is said to be unchangeable and read-only.

When declaring the value of a const it's value must be always assigned to it. e.g
  Const PI = 3.14



  They are different rules in naming a const which are:

  Const name must be written in capital letters in order to differenciate between other variables.
  
  Const can  be used both outside and inside of a function body.

 # TYPES OF CONSTANTS

 We have 2 types of constants which are:
 TYPED CONSTANT
 UNTYPED CONSTANT

 TYPED CONSTANT
 They are declared with a type for example:
   const A int = 1


UNTYPED CONSTANT
They are declared without using a type for example:
   Const A = 1

 When a constant is declared it is not possible to change the value later

   func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}  

More than one constant declaration

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)