package console

import (
	"fmt"
)

//PrintEmo : function to print emojis
func PrintEmo(num int, emo string, a ...interface{}) {
	for i := 0; i < num; i++ {
		fmt.Printf(emo + " ")
	}
	fmt.Printf("%s\n", a...)
}

//Log : function to print log
func Log(a ...interface{}) {
	fmt.Printf("%s\n", a...)
}

//Pizza : function to print pizza
func Pizza(num int, a ...interface{}) {
	PrintEmo(num, "🍕", a...)
}

//Beer : function to print beer
func Beer(num int, a ...interface{}) {
	PrintEmo(num, "🍺", a...)
}

//Unicorn : function to print unicorn
func Unicorn(num int, a ...interface{}) {
	PrintEmo(num, "🦄", a...)
}

//Poop : function to print poop
func Poop(num int, a ...interface{}) {
	PrintEmo(num, "💩", a...)
}

//Laugh : function to print laugh
func Laugh(num int, a ...interface{}) {
	PrintEmo(num, "😂", a...)
}

//Amazing : function to print amazing
func Amazing(num int, a ...interface{}) {
	PrintEmo(num, "😍", a...)
}

//Angry : function to print angry
func Angry(num int, a ...interface{}) {
	PrintEmo(num, "😡", a...)
}

//Sad : function to print sad
func Sad(num int, a ...interface{}) {
	PrintEmo(num, "😣", a...)
}

//Afraid : function to print afraid
func Afraid(num int, a ...interface{}) {
	PrintEmo(num, "😨", a...)
}

//Blush : function to print blush
func Blush(num int, a ...interface{}) {
	PrintEmo(num, "😊", a...)
}

//Grin : function to print grin
func Grin(num int, a ...interface{}) {
	PrintEmo(num, "😁", a...)
}

//Clap : function to print clap
func Clap(num int, a ...interface{}) {
	PrintEmo(num, "👏", a...)
}

//Earth : function to print earth
func Earth(num int, a ...interface{}) {
	PrintEmo(num, "🌎", a...)
}

//Spades : function to print spades
func Spades(num int, a ...interface{}) {
	PrintEmo(num, "♠", a...)
}

//Warning : function to print warning
func Warning(num int, a ...interface{}) {

	PrintEmo(num, "⚠", a...)
}

//Hamburger : function to print hamburger
func Hamburger(num int, a ...interface{}) {
	PrintEmo(num, "🍔", a...)
}

//Crown : function to print crown
func Crown(num int, a ...interface{}) {
	PrintEmo(num, "👑", a...)
}

//Palette : function to print palette
func Palette(num int, a ...interface{}) {
	PrintEmo(num, "🎨", a...)
}

//Bomb : function to print bomb
func Bomb(num int, a ...interface{}) {
	PrintEmo(num, "💣", a...)
}

//Octopus : function to print octopus
func Octopus(num int, a ...interface{}) {
	PrintEmo(num, "🐙", a...)
}

//Snake : function to print snake
func Snake(num int, a ...interface{}) {
	PrintEmo(num, "🐍", a...)
}

//PandaFace : function to print panda face
func PandaFace(num int, a ...interface{}) {
	PrintEmo(num, "🐼", a...)
}

//CatFace : function to print cat face
func CatFace(num int, a ...interface{}) {
	PrintEmo(num, "🐱", a...)
}

//Zap : function to print zap
func Zap(num int, a ...interface{}) {
	PrintEmo(num, "⚡", a...)
}

//Exclamation : function to print exclamation
func Exclamation(num int, a ...interface{}) {
	PrintEmo(num, "❗", a...)
}

//Question : function to print question mark
func Question(num int, a ...interface{}) {
	PrintEmo(num, "❓", a...)
}

//Heart : function to print heart
func Heart(num int, a ...interface{}) {
	PrintEmo(num, "️❤", a...)
}
