package icli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct{
	reader *bufio.Reader

	commander commander

	helper *helpCmd

	errColor Color
	descColor Color
	usageColor Color

	inputPointer string
	welcomeMessage string
}

//NewCLI returns a CLI with an embedded helper and exit handler.
//The default input pointer is '>>>'
//The default error color is RED
func NewCLI() *CLI{
	handler := &CLI{}
	handler.reader = bufio.NewReader(os.Stdin)
	handler.SetErrorColor(RED)
	handler.inputPointer = ">>>"

	helper := newHelper()
	handler.helper = helper
	handler.AddCmd(helper)

	exitCmd := &exitCmd{BasicCommand{Name: "exit",Fn: fnExit,Description: "Finish the cli execution.",Usage: "exit"}}
	handler.AddCmd(exitCmd)

	return handler
}

//NewEmptyCLI returns an empty CLI handler.
func NewEmptyCLI() *CLI{
	handler := &CLI{}
	handler.reader = bufio.NewReader(os.Stdin)
	handler.inputPointer = ">>>"

	return handler
}

//Run starts the CLI.
func (h *CLI) Run() {
	fmt.Println(h.welcomeMessage)
	for {
		fmt.Print(h.inputPointer)
		input, _ := h.reader.ReadString('\n')
		if len(input) == 1 {
			continue
		}
		h.checkCommands(input)
	}
}

//SetErrorColor sets the error color in the CLI.
func (h *CLI) SetErrorColor(color Color) {
	h.errColor = color
}

//SetWelcomeMessage sets the welcome message for when the cli is started.
func (h *CLI) SetWelcomeMessage(msg string){
	h.welcomeMessage = msg
}

//AddCmd adds a command to the cli. The command should implements Command interface.
//it also adds that command to the helper if it's set.
func (h *CLI) AddCmd(cmd Command){
	h.commander.add(cmd)
	if h.helper != nil {
		h.helper.updateCmd(cmd)
	}
}

//AddCmds adds multiple command to the cli. The commands should implement Command interface.
//it also adds every command to the helper if it's set.
func (h *CLI) AddCmds(cmds ...Command){
	for _, cmd := range cmds{
		h.commander.add(cmd)
		if h.helper != nil {
			h.helper.updateCmd(cmd)
		}
	}
}


//checkCommands check in the Run loop execution if the commands is valid and execute it.
func (h *CLI) checkCommands(input string){
	call := strings.Fields(input)
	cmdName := call[0]


	cmd, err := h.commander.get(cmdName)
	if err != nil {
		fmt.Printf("%v %v %v\n",h.errColor, err.Error(), RESET)
	}else{
		errCmd := cmd.GetFn()(call[1:]...)
		if errCmd != nil {
			fmt.Printf("%v %v %v\n",h.errColor, errCmd, RESET)
		}
	}
}