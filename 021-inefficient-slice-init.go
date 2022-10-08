package main

type SliceFoo struct {
	a string
}

type SliceBar struct {
	b string
}

func fooToBar(f SliceFoo) SliceBar {
	return SliceBar{
		b: f.a,
	}
}

// slice initialization with zero length will incur performance penalty of expanding the backing array
func inefficientConvert(foos []SliceFoo) []SliceBar {
	bars := make([]SliceBar, 0)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// initializing a slice with the known output length and then accessing elements by index is most performant
func efficientConvert(foos []SliceFoo) []SliceBar {
	bars := make([]SliceBar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}
