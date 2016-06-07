package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Raw struct {
	column0 string
	column1 string
	column2 string
	column3 string
}

func printRaws(raws []*Raw) {
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "column0", "column1", "column2", "column3")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----")
	for _, t := range raws {
		fmt.Fprintf(tw, format, t.column0, t.column1, t.column2, t.column3)
	}
	tw.Flush() // calculate column widths and print table
}

type memorizingSort struct {
	t        []*Raw
	memories []func(x, y *Raw) bool
}

func (x memorizingSort) Len() int      { return len(x.t) }
func (x memorizingSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x memorizingSort) Less(i, j int) bool {
	result := false
	for _, less := range x.memories {
		result = less(x.t[i], x.t[j])
	}
	return result
}

var m = []func(x, y *Raw) bool{
	func(x, y *Raw) bool { return x.column0 < y.column0 },
	func(x, y *Raw) bool { return x.column1 < y.column1 },
	func(x, y *Raw) bool { return x.column2 < y.column2 },
	func(x, y *Raw) bool { return x.column3 < y.column3 },
}

var raws = []*Raw{
	{"A", "3", "い", "a"},
	{"A", "1", "う", "b"},
	{"B", "1", "あ", "c"},
	{"C", "2", "あ", "a"},
}

func main() {

	sort.Sort(memorizingSort{raws, []func(x, y *Raw) bool{
		func(x, y *Raw) bool { return x.column0 < y.column0 }}})
	printRaws(raws)

	sort.Sort(memorizingSort{raws, []func(x, y *Raw) bool{
		func(x, y *Raw) bool { return x.column0 < y.column0 },
		func(x, y *Raw) bool { return x.column1 < y.column1 }}})
	printRaws(raws)

	sort.Sort(memorizingSort{raws, []func(x, y *Raw) bool{
		func(x, y *Raw) bool { return x.column0 < y.column0 },
		func(x, y *Raw) bool { return x.column1 < y.column1 },
		func(x, y *Raw) bool { return x.column2 < y.column2 }}})
	printRaws(raws)

	sort.Stable(memorizingSort{raws, []func(x, y *Raw) bool{
		func(x, y *Raw) bool { return x.column0 < y.column0 },
		func(x, y *Raw) bool { return x.column1 < y.column1 },
		func(x, y *Raw) bool { return x.column2 < y.column2 }}})
	printRaws(raws)
}
