# Go and OOO course

## struct in go
example:
```
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}
````
how to use it:
```
func main() {
	contaDoDiegues := ContaCorrente{
		titular:       "Diegues",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}
	fmt.Println(contaDoDiegues)
}
```

## Zero Initializatrion
* bool = false
* int = 0
* float = 0
* string = ""
* struct = {}
    
A Null pointer in go is **nil**

## problema ao criar novos pacotes, para resolver rodar: 
* go mod init
* go mod tidy