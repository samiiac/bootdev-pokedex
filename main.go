package main 

type config struct{
	commandRegistry map[string]command
	previous        string
	next            string
}

func main(){
	c :=  config{
	commandRegistry : getCommands(),
	previous : "",
	next : "",

 }

	startRepl(&c)
}