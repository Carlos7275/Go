package mylibrary;

func Pow(x float32, n int) float32 {
	if n == 0 {
		return 1
	}

	half := Pow(x, n/2)

	if n%2 == 0 {
		return half * half
	}

	if n > 0 {
		return x * half * half
	}
	return (half * half) / x
}
