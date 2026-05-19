package string_helper

func InterfaceArrToStringArr(params []interface{}) []string {
	var result []string
	for i := range params {
		result = append(result, params[i].(string))
	}

	return result
}
