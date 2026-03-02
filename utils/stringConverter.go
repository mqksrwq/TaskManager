package utils

import (
	. "TaskManager/task"
	"fmt"
	"strings"
)

// ConvertToString - метод для конвертации динамического массива в строку
func ConvertToString(tasks []*Task) string {
	builder := new(strings.Builder)
	for _, t := range tasks {
		builder.WriteString(fmt.Sprintf("\nID задачи: %d\n"+
			"Название задачи: %s\n"+
			"Описание задачи: %s\n"+
			"Выполнена: %s\n", t.GetId(), t.GetName(),
			t.GetDescription(), boolToString(t.GetStatus())))
	}
	return builder.String()
}
