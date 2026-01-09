package pkg

type Value struct {
	Type Type
	Data interface{}
}

func IntValue(value int) Value {
	return Value{
		Type: TYPE_INT,
		Data: value,
	}
}

func BoolValue(value bool) Value {
	return Value{
		Type: TYPE_BOOL,
		Data: value,
	}
}
