package main

type TFExercise interface {
	info() string
}

func PrintValue(val any) {
	switch val.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case TFExercise:
		println("TFExercise")
	default:
		println("unknown type")
	}
}

func Add(a, b any) any {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case string:
		return a.(string) + b.(string)
	case TFExercise:
		return a.(TFExercise).info() + b.(TFExercise).info()
	default:
		return nil
	}
}

func GenericAdd[T int | float32 | string](a, b T) T {
	return a + b
}
