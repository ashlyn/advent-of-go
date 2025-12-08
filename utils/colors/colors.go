package colors

// Reset resets the console color to default
var Reset  = "\033[0m"
// Red changes the console color to red
var Red    = "\033[31m"
// Green changes the console color to green
var Green  = "\033[32m"
// Yellow changes the console color to yellow
var Yellow = "\033[33m"
// Blue changes the console color to blue
var Blue   = "\033[34m"
// Purple changes the console color to purple
var Purple = "\033[35m"
// Cyan changes the console color to cyan
var Cyan   = "\033[36m"
// Gray changes the console color to gray
var Gray   = "\033[37m"
// White changes the console color to white
var White  = "\033[97m"

func formatColorString(color string, str string) string {
	return color + str + Reset
}

// RedString returns the given string formatted in red color
func RedString(str string) string {
	return formatColorString(Red, str)
}

// GreenString returns the given string formatted in green color
func GreenString(str string) string {
	return formatColorString(Green, str)
}

// YellowString returns the given string formatted in yellow color
func YellowString(str string) string {
	return formatColorString(Yellow, str)
}

// BlueString returns the given string formatted in blue color
func BlueString(str string) string {
	return formatColorString(Blue, str)
}

// PurpleString returns the given string formatted in purple color
func PurpleString(str string) string {
	return formatColorString(Purple, str)
}

// CyanString returns the given string formatted in cyan color
func CyanString(str string) string {
	return formatColorString(Cyan, str)
}

// GrayString returns the given string formatted in gray color
func GrayString(str string) string {
	return formatColorString(Gray, str)
}

// WhiteString returns the given string formatted in white color
func WhiteString(str string) string {
	return formatColorString(White, str)
}