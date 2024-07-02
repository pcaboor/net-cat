package server

import (
	"github.com/gookit/color"
)

func PrintWelcome() string {
	return color.White.Sprintf("Welcome to TCP-Chat!\n") +
		color.White.Sprintf("         _nnnn_\n") +
		color.White.Sprintf("        dGGGGMMb\n") +
		color.White.Sprintf("       @p~qp~~qMb\n") +
		color.White.Sprintf("       M|@||@) M|\n") +
		color.LightYellow.Sprintf("       @,----.JM|\n") +
		color.LightYellow.Sprintf("      JS^\\__/  qKL\n") +
		color.White.Sprintf("     dZP        qKRb\n") +
		color.White.Sprintf("    dZP          qKKb\n") +
		color.White.Sprintf("   fZP            SMMb\n") +
		color.White.Sprintf("   HZM            MMMM\n") +
		color.White.Sprintf("   FqM            MMMM\n") +
		color.LightYellow.Sprintf(" __| \".        |\\dS\"qML\n") +
		color.LightYellow.Sprintf(" |    `.       | `' \\Zq\n") +
		color.LightYellow.Sprintf("_)      \\.___.,|     .'\n") +
		color.HiYellow.Sprintf("\\____   )"+color.White.Sprintf("MMMMMP")+color.HiYellow.Sprintf("|   .'\n")) +
		color.LightYellow.Sprintf("     `-'       `--'\n")

}
