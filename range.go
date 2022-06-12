// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"math"
)

const defaultAxisDivideCount = 6

type axisRange struct {
	divideCount int
	min         float64
	max         float64
	size        int
	boundary    bool
}

type AxisRangeOption struct {
	Min         float64
	Max         float64
	Size        int
	Boundary    bool
	DivideCount int
}

func NewRange(opt AxisRangeOption) axisRange {
	max := opt.Max
	min := opt.Min

	max += math.Abs(max * 0.1)
	min -= math.Abs(min * 0.1)
	divideCount := opt.DivideCount
	r := math.Abs(max - min)

	// 最小单位计算
	unit := 2
	if r > 10 {
		unit = 4
	}
	if r > 30 {
		unit = 5
	}
	if r > 100 {
		unit = 10
	}
	if r > 200 {
		unit = 20
	}
	unit = int((r/float64(divideCount))/float64(unit))*unit + unit

	if min != 0 {
		isLessThanZero := min < 0
		min = float64(int(min/float64(unit)) * unit)
		// 如果是小于0，int的时候向上取整了，因此调整
		if min < 0 ||
			(isLessThanZero && min == 0) {
			min -= float64(unit)
		}
	}
	max = min + float64(unit*divideCount)
	return axisRange{
		divideCount: divideCount,
		min:         min,
		max:         max,
		size:        opt.Size,
		boundary:    opt.Boundary,
	}
}

func (r axisRange) Values() []string {
	offset := (r.max - r.min) / float64(r.divideCount)
	values := make([]string, 0)
	for i := 0; i <= r.divideCount; i++ {
		v := r.min + float64(i)*offset
		value := commafWithDigits(v)
		values = append(values, value)
	}
	return values
}

func (r *axisRange) getHeight(value float64) int {
	v := (value - r.min) / (r.max - r.min)
	return int(v * float64(r.size))
}

func (r *axisRange) getRestHeight(value float64) int {
	return r.size - r.getHeight(value)
}

func (r *axisRange) GetRange(index int) (float64, float64) {
	unit := float64(r.size) / float64(r.divideCount)
	return unit * float64(index), unit * float64(index+1)
}
func (r *axisRange) AutoDivide() []int {
	return autoDivide(r.size, r.divideCount)
}

func (r *axisRange) getWidth(value float64) int {
	v := value / (r.max - r.min)
	// 移至居中
	if r.boundary &&
		r.divideCount != 0 {
		v += 1 / float64(r.divideCount*2)
	}
	return int(v * float64(r.size))
}
