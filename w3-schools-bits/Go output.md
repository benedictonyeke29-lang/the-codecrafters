# GO OUTPUTS FUMNCTIONS
 
 Golang makes use of 3 functions to output or display text which are:

  Print: This functions prints only outputs in their default format only.

  func main() {
  var a,b string = "Ben","Edict"

  fmt.Print(a)
  fmt.Print(b)
    }

    Result:
    BenEdict


  #  And if we want the outputs to be on a different line we an do:

    func main() {
  var a,b string = "Ben","Edict"

  fmt.Print(a, "\n")
  fmt.Print(b, "\n")
    }

    Result:
    Ben
    Edict


 # NOTE 
  "\n" is used for creating new lines


  #  We can use one Print function to print multiple variables just like this:

  package main

    import (
	"fmt"
    )

    var i, j string = "Ben", "Edict"

    func main() {

	fmt.Print(i, "\n", j)
    }
 
 
  #  Maybe if you intend to add space in between the string we use " ". Let's try an example

    package main

    import (
    "fmt"
    )

    var i, j string = "Ben", "Edict"

    func main() {

    fmt.Print(i, " ", j)
    }


 The Print function also insert space in between argument if neither are strings:

    package main
    import ("fmt")

    func main() {

    var i,j = 3,4
    fmt.Print(i,j)
    }

    Result:
    3 4

 # THE Println function
  The Println