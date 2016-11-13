package tempconv

// 摄氏转华氏
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 华氏转摄氏
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
