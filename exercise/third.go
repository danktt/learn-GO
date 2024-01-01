package main

import "fmt"


var x int = 42
var y string = "James Bound"
var z bool = true



func main() {
	s :=  fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Printf("%v %T \n", x, x  )
	fmt.Printf("%v %T \n", y, y  )
	fmt.Printf("%v %T \n", z, z  )
	
	fmt.Println(s)
}
// Na prática: exercício #2
// Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
// Identificador "x" deverá ter tipo int
// Identificador "y" deverá ter tipo string
// Identificador "z" deverá ter tipo bool
// Na função main:
// Demonstre os valores de cada identificador
// O compilador atribuiu valores para essas variáveis. Como esses valores se chamam? => Res: O nome é valor ZERO


// Utilizando a solução do exercício anterior:
// Em package-level scope, atribua os seguintes valores às variáveis:
// para "x" atribua 42
// para "y" atribua "James Bond"
// para "z" atribua true
// Na função main:
// Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
// Demonstre a variável "s".