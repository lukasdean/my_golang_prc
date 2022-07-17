package prc_temp

// CToF 把摄氏温度转化为华氏温度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 把华氏温度转化为摄氏温度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
