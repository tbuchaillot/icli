# ICLI
Interactive CLI Builder written in Go

#Usage
   - Create a new CLI.
   - Assign any Command you want.
   - Run the CLI :) .

#Customization
- `SetWelcomeMessage` allows you to set message shown on the cli start.
- `SetErrorColor` allows you to set the error message color.
- The lib have defined 7 colors in the colors file.
- `AddCmd` or `AddCmds` allows you to add any command that implements `Command` interface. 
- The library has already a `BasicCommand` struct that implements `Command` interface. 

#Examples
```go

import (
    "fmt"
    "github.com/tbuchaillot/icli"
)
func main(){
	//Creates a new interactive cli
	cli := icli.NewCLI()
	//Set the welcome message for when we open the cli
	cli.SetWelcomeMessage(fmt.Sprintf("%v Welcome to interactive CLI!%v",icli.BLUE,icli.RESET))

	//Creates a new BasicCommand just for demo pourpose.
	testCmd := &icli.BasicCommand{
		Name:"test",
		Description: "Used to demo this :D ",
		Usage: "test <ARGS>",
		Fn: TestCmdFn,
	}

	//Add that command to the cli
	cli.AddCmd(testCmd)

	//Run the cli :)
	cli.Run()
}

//TestCmdFn is the function for when test command is executed. It's going to print the arguments given.
func TestCmdFn(args ...string) error{
	fmt.Println("test called with args:",args)
	return nil
}

```
