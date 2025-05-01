/* Switch Statement :---> 
the switch statement in go used for multiple conditional execution .its a cleaner way to write a sequence of if - else if - else statements*/

package main

import "fmt"

func main() {
    fmt.Println("Basic Switch Statement:")
    day := 3 
    switch day{
        case 1:
            fmt.Println("Monday")
        case 2:
            fmt.Println("Tuesday")
        case 3:
            fmt.Println("Wednesday")
        case 4: 
            fmt.Println("Thursday")
        case 5: 
            fmt.Println("Friday")
        
        default:
            fmt.Println("Invalid Day")
        }
  
   