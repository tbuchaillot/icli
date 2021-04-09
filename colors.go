package icli

type Color string

const(
	GREEN Color = "\033[32m"
	YELLOW Color = "\033[33m"
	BLUE Color = "\033[34m"
	RED Color = "\033[31m"
	PURPLE Color = "\033[35m"
	GRAY Color = "\033[37m"
	WHITE Color = "\033[97m"
	RESET Color = "\033[0m"
)
