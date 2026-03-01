package utils

import (
	. "TaskManager/task"
	"fmt"
	"strings"
)

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

func boolToString(b bool) string {
	if b {
		return "да"
	}
	return "нет"
}
