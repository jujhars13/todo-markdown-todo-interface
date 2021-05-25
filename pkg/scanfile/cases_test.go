package scanfile

// various test cases of markdown permutations
// can be a bit tricky to read because of multiline strings
var testCases = []struct {
	testMarkdownString string
	expected           []int
}{
	{``, []int{}},
	{`- [x] todo item done 1`, []int{}},
	{`- [ ] todo item 1`, []int{1}},
	{`random string
  - [ ] todo item 1 with random strings`, []int{1}},
	{`random string
  - [ ] todo item 1 with random strings before and after
  moar random strings`, []int{1}},
	{`- [ ] todo item 1
- [ ] todo item 2
- [ ] todo item 3
`, []int{1, 2, 3}},
	{`- [ ] todo item 1
- [ ] todo item 2
- [ ] todo item 3
- [ ] todo item 4
- [ ] todo item 5
- [ ] todo item 6
`, []int{1, 2, 3, 4, 5, 6}},
	{`- [ ] todo item 1
  random string
  random string two
- [ ] todo item 2
- [ ] todo item 3
`, []int{1, 4, 5}},
	{`- [x] todo item 1
  random string
  random string two
- [ ] todo item 2
random string two
- [ ] todo item 3
`, []int{4, 6}},
	{`- [ ] todo item 1
  random string
  random string two
- [x] todo item 2
random string two
- [x] todo item 3
- [ ] todo item 3
`, []int{1, 7}},
}
