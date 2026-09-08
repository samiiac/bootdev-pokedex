package main 

import ("fmt"
 "strings"
 "bufio"
 "os"
)


func cleanInput(text string)[] string {
	return strings.Fields(strings.ToLower(text))
}


func startRepl(c *config) {
	scanner:= bufio.NewScanner(os.Stdin)
	for true {
    fmt.Print("Pokedex > ")
	scanner.Scan()
    input := cleanInput(scanner.Text())
	if len(input) == 0 {
     continue
	}
	args := ""
	if len(input) > 1 {
		args = input[1]
	}
	
	if command,ok := c.commandRegistry[input[0]];ok {
     err := command.callback(c,args)
	 if err != nil {
		fmt.Println(err)
	 }
	}else{
	fmt.Println("Unknown Command")
	}
	 
	}
}


