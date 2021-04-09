package icli


import (
	"fmt"
	"os"
	"strings"
)

//BasicCommand implements Command interface. It usually good idea to use it as a default command.
type BasicCommand struct{
	//Fn is the func to be executed when that command is called
	Fn   func(...string) error
	//Name is the command name.
	Name string
	//Description is the description of what that command does.
	Description string
	//Usage is the description of the usage of that command.
	Usage	string

	//TODO
	Options []Options
}

//TODO Options
type Options struct{
	Name string
	Description	string

	Required bool
}

func (cmd *BasicCommand) SetFn(fn func(...string) error){
	cmd.Fn = fn
}

func (cmd *BasicCommand) GetFn()  func(...string) error{
	return cmd.Fn
}

func (cmd *BasicCommand) SetName(name string){
	cmd.Name = name
}

func (cmd *BasicCommand) GetName()	string{
	return cmd.Name
}

func (cmd *BasicCommand) SetDescription(desc string){
	cmd.Description = desc
}

func (cmd *BasicCommand) GetDescription() string{
	return cmd.Description
}

func (cmd *BasicCommand) SetUsage(usage string){
	cmd.Usage = usage
}

func (cmd *BasicCommand) GetUsage() string{
	return cmd.Usage
}


//helpCmd is the built in helper for the commands.
type helpCmd struct {
	BasicCommand
	cmds []Command
}

//newHelper returns the built in - basic helper
func newHelper() *helpCmd{
	helper := &helpCmd{}
	helper.Name ="help"
	helper.Description= "It shows the help of commands"
	helper.Usage = "help or help <CMD>"
	helper.Fn = helper.fnHelp

	return helper
}

//fnHelp is the func used when help command is executed
func (helper *helpCmd) fnHelp(args ...string) error{
	if len(args) == 0 {
		cmdsName := []string{}
		for _, cmd :=  range helper.cmds{
			cmdsName = append(cmdsName, cmd.GetName())
		}

		fmt.Printf(" Available commands:%v %v %v.\n",GREEN, strings.Join(cmdsName,","),RESET)
	} else {
		cmdName := args[0]

		for _, cmd := range helper.cmds{
			if cmdName == cmd.GetName(){
				fmt.Printf(" %vDescription:%v %v\n",GREEN,RESET,cmd.GetDescription())
				fmt.Printf(" %vUsage:%v %v\n",GREEN,RESET,cmd.GetUsage())
				return nil
			}
		}
		fmt.Printf(" %vInvalid command.%v\n",RED,RESET)
	}

	return nil
}

//updateCmd is the helper function to maintain the list commands updated
func (helper *helpCmd) updateCmd (cmd Command){
	helper.cmds = append(helper.cmds,cmd)
}

//exitCmd is the build-in basic exit command
type exitCmd struct {
	BasicCommand
}

//fnExit is the basic exit function for exit command
func fnExit(args ...string) error{
	fmt.Println("\033[34mGoodbye!\033[0m")
	os.Exit(0)
	return nil
}