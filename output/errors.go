package output

import "github.com/fatih/color"

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")
}
