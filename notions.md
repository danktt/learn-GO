## Cap 1
<nil> in GO means: null in JS

We can use _ in GO to ignore a return value, the name is **Blank Identifier**

## Cap2

 - We can use := to declare and assign a variable, but we can't use it to reassign a variable
 - We use := just inside a func scope, outside we use =

 Zero values:
  - int: 0
  - float: 0.0
  - bool: false
  - string: ""
  - pointer: nil
  - function: nil
  - interface: nil
  - slice: nil
  - channel: nil
  - map: nil


  Format printing: documentação.
    Grupo #1: Print -> standard out
      func Print(a ...interface{}) (n int, err error)
      func Println(a ...interface{}) (n int, err error)
      func Printf(format string, a ...interface{}) (n int, err error)
  Format verbs. (%v %T)
    Grupo #2: Print -> string, pode ser usado como variável
      func Sprint(a ...interface{}) string
      func Sprintf(format string, a ...interface{}) string
      func Sprintln(a ...interface{}) string
    Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
    func Fprint(w io.Writer, a ...interface{}) (n int, err error)
    func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
    func Fprintln(w io.Writer, a ...interface{}) (n int, err error)



## Cap3
