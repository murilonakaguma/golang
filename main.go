package main

import (
    "fmt"
)
func dividir(dividendo int, divisor int)  (int, string){
if divisor == 0 {
    return 0, "Erro na divisâo pro zero"
}
    return dividendo / divisor, ""
}
func operaçaoBasica(a int, b int) (int, int, int){
soma := a + b
multiplicacao := a * b
subtracao := a - b 
return soma, multiplicacao, subtracao
}
func main() {
  resultado, erro := dividir(10,2)
  if erro != "" {
    fmt.Println("erro")
  }else{
    fmt.Println("o resultado da divisao é:", resultado)
  }
  soma, mult, sub := operaçaoBasica(10,2)
  fmt.Println(soma)
  fmt.Println(mult)
  fmt.Println(sub)
}
 