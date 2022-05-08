package homework

func average(input [15]float32) (result float32) {
//Place your code here
var sum float32 = 0
	for _, number := range input {
		sum += number
	}	
	result = sum / 15
	return
}
