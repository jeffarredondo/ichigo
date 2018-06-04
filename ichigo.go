package ichigo

import (
	"errors"
)

func maxOfSlice(aSlice []float64) float64 {
	maxValue := 0.0
	for i := 0; i < len(aSlice); i++ {
		if aSlice[i] > maxValue {
			maxValue = aSlice[i]
		}
	}

	return maxValue
}

func minOfSlice(aSlice []float64) float64 {
	maxValue := 100000.0
	for i := 0; i < len(aSlice); i++ {
		if aSlice[i] < maxValue {
			maxValue = aSlice[i]
		}
	}

	return maxValue
}

//Tenkansen returns the Tenkan-sen conversion line which is the
//	average of the 9 period high and 9 period low
func Tenkansen(close []float64) []float64 {

	tenkansen := make([]float64, (len(close) + 26))

	for i := 8; i < len(close); i++ {
		tenkansen[i] = (maxOfSlice(close[(i-8):i]) + (minOfSlice(close[(i - 8):i]))) / 2
	}
	return tenkansen
}

//Kijunsen is the Kijun-sen base line for the average of the 26
//	period high and 26 period low
func Kijunsen(close []float64) []float64 {

	kijunsen := make([]float64, (len(close) + 26))

	for i := 25; i < len(close); i++ {
		kijunsen[i] = (maxOfSlice(close[(i-25):i]) + (minOfSlice(close[(i - 25):i]))) / 2
	}
	return kijunsen
}

//SenkouSpanA is the average of the Tenkan-sen and Kijun-sen
//	this is a rolling daily average posted 26 periods in the future
func SenkouSpanA(conversion, base []float64) []float64 {

	spanA := make([]float64, (len(conversion) + 26))

	for i := 25; i < len(spanA); i++ {
		spanA[i] = (conversion[i-25] + base[i-25]) / 2

	}

	return spanA
}

//SenkouSpanB is the average of the 52 period high and 52 period low
//	This is placed 26 days in the future
func SenkouSpanB(close []float64) []float64 {

	spanB := make([]float64, (len(close) + 26))

	for i := 51; i < len(spanB); i++ {
		spanB[i] = (maxOfSlice(close[(i-51):i]) + minOfSlice(close[(i-51):i])) / 2

	}

	return spanB
}

//ChikouSpan close lagged 26 period
func ChikouSpan(close []float64) []float64 {
	chikou := make([]float64, len(close)+26)

	for i := 25; i < len(close); i++ {
		chikou[i-25] = close[i]
	}
	return chikou
}

//IchiMe takes a slice of Close and builds the ichimoku cloud.
// This function requires at least 52 days of data. I can make this
// work with less data, but going with this for now
func IchiMe(close []float64) ([][]float64, error) {
	//multidimensional array will be n + 26 days long and 5 wide
	var n int
	var ichicloud [][]float64

	n = len(close) + 26

	if n < 52 {
		return ichicloud, errors.New("Not enough data to form ichimoku cloud")
	}

	//Tenkan-sen
	tenkansen := Tenkansen(close)

	//Kijun-sen
	kijunsen := Kijunsen(close)

	//Senkou Span A
	senkouA := SenkouSpanA(tenkansen, kijunsen)

	//Senkou Span B
	senkouB := SenkouSpanB(close)

	//ChikouSpan
	chikou := ChikouSpan(close)

	//ToDo:
	//Span Thicky Thiccness

	//Slope Span A

	//Bull v Bear

	//Conversion Position (-1, 0, 1)

	//Rolling count of conversion positions

	var tempSlice []float64

	for i := 0; i < n; i++ {
		tempSlice = []float64{}
		tempSlice = append(tempSlice, tenkansen[i])
		tempSlice = append(tempSlice, kijunsen[i])
		tempSlice = append(tempSlice, senkouA[i])
		tempSlice = append(tempSlice, senkouB[i])
		tempSlice = append(tempSlice, chikou[i])

		ichicloud = append(ichicloud, tempSlice)
	}

	return ichicloud, nil

}
