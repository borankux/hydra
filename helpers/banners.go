package helpers

var art string = "\n    __              __          \n   / /_  __  ______/ /________ _\n  / __ \\/ / / / __  / ___/ __ `/\n / / / / /_/ / /_/ / /  / /_/ / \n/_/ /_/\\__, /\\__,_/_/   \\__,_/  \n      /____/                    \n"
var clear string = "\n██╗  ██╗██╗   ██╗██████╗ ██████╗  █████╗ \n██║  ██║╚██╗ ██╔╝██╔══██╗██╔══██╗██╔══██╗\n███████║ ╚████╔╝ ██║  ██║██████╔╝███████║\n██╔══██║  ╚██╔╝  ██║  ██║██╔══██╗██╔══██║\n██║  ██║   ██║   ██████╔╝██║  ██║██║  ██║\n╚═╝  ╚═╝   ╚═╝   ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝\n                                         \n"

func GetBanner(name string) string {
	switch name {
	case "art":
		return art
	case "clear":
		return clear
	default:
		return clear
	}
}
