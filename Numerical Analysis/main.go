/*
Create a struct that can keep track of the following NumericalAnalyser stats for a stream of input data:
  - sum
  - avg
  - median
  - max
  - min
*/
package main

type NumericalAnalyser struct {
	Count   int
	Sum     float32
	Avg     float32
	Min     float32
	Max     float32
	Median  float32
	maxHeap []float32 // keep track of the smallest numbers seen
	minHeap []float32 // keep track of the largest numbers seen
}

func NewNumerical() *NumericalAnalyser {
	analyser := new(NumericalAnalyser)
	return analyser
}

func (analyser *NumericalAnalyser) insert(x float32) {
	// update count
	analyser.Count += 1

	// update sum
	analyser.Sum += x

	// update avg
	analyser.Avg = analyser.Sum / float32(analyser.Count)

	// update min
	if x < analyser.Min {
		analyser.Min = x
	}

	// update max
	if x > analyser.Max {
		analyser.Max = x
	}

	//TODO: work on median with the heaps

}

func (analyser *NumericalAnalyser) remove(x float32) bool {
	//TODO: handle removal of elements from the stream
	return false
}
