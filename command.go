package icli

import (
	"errors"
	"fmt"
)

//Command interface sets the common methods that a Command should have
type Command interface {
	SetFn( func(...string) error)
	GetFn()  func(...string) error
	SetName(string)
	GetName()	string
	SetDescription(string)
	GetDescription() string
	SetUsage(string)
	GetUsage() string
}

type commander struct {
	cmap map[string]Command
}

//get returns a command given its name. In case it doesn't exist, it throws an error
func (cmds *commander) get(cmdName string) (Command,error){
	if cmd, ok := cmds.cmap[cmdName]; ok{
		return cmd, nil
	}
	return nil,errors.New(fmt.Sprintf("Command %v doesn't exists",cmdName))
}

//add adds a command to the commander
func (cmds *commander) add(cmd Command) error{
	if len(cmds.cmap) == 0 {
		cmds.cmap = make(map[string]Command)
	}
	cmds.cmap[cmd.GetName()] = cmd
	return nil
}



