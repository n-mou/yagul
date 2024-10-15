package list_test

import (
	legacy "container/list"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/n-mou/yagul/list"
)

type listTuple[T any] struct {
	New    *list.List[T]
	Legacy *legacy.List
}

type testCase struct {
	name     string
	expected string
	result   string
}

func (c testCase) run(t *testing.T) {
	t.Logf("Testing case: %v", c.name)
	if c.result == c.expected {
		t.Logf("\tEXPECTED: %v", c.expected)
		t.Logf("\tRESULT: %v", c.result)
		t.Log("SUCCESS")
	} else {
		t.Errorf("ERROR:\n\tEXPECTED: %v\n\tGOT: %v", c.expected, c.result)
	}
}

func stringifyTestCase(name string, lists listTuple[string], t *testing.T) {
	tc := testCase{
		name,
		containerListToString(lists.Legacy),
		lists.New.String(),
	}
	tc.run(t)
}

func lengthTestCase(name string, lists listTuple[string], t *testing.T) {
	tc := testCase{
		name,
		strconv.Itoa(lists.Legacy.Len()),
		strconv.Itoa(lists.New.Len()),
	}
	tc.run(t)
}

func testLists() listTuple[string] {
	testListSlice := []string{"Longing", "Rusted", "Furnace", "Daybreak", "Seventeen", "Benign", "Nine", "Homecoming", "One", "Freight car"}
	testLists := listTuple[string]{
		list.New(testListSlice...),
		legacy.New(),
	}
	for _, i := range testListSlice {
		testLists.Legacy.PushBack(i)
	}
	return testLists
}

func testSuite() []listTuple[string] {
	emptyLists := listTuple[string]{
		list.New[string](),
		legacy.New(),
	}

	singleNodeLists := listTuple[string]{
		list.New("one"),
		legacy.New(),
	}
	singleNodeLists.Legacy.PushBack("one")

	twoNodesLists := listTuple[string]{
		list.New("one", "two"),
		legacy.New(),
	}
	twoNodesLists.Legacy.PushBack("one")
	twoNodesLists.Legacy.PushBack("two")

	testLists := testLists()
	return []listTuple[string]{emptyLists, singleNodeLists, twoNodesLists, testLists}
}

func containerListToString(l *legacy.List) string {
	el := l.Front()
	sl := make([]string, 0, l.Len())
	for el != nil {
		sl = append(sl, fmt.Sprintf("%v", el.Value))
		el = el.Next()
	}
	return "[" + strings.Join(sl, " -> ") + "]"
}

var tsNames = []string{"empty list", "single node list", "two nodes list", "test list"}

func TestNew(t *testing.T) {
	ts := testSuite()
	for i := range ts {
		// I know strings.Title is deprecated due to issues with Unicode characters but there's
		// no alternative in the standard library, also it can handle ASCII without any problems
		stringifyTestCase(fmt.Sprintf("%v intialization", strings.Title(tsNames[i])), ts[i], t)
	}
}
func TestBack(t *testing.T) {
	ts := testSuite()

	for i := range ts {
		expected := "<nil>"
		result := "<nil>"
		if ts[i].Legacy.Back() != nil {
			expected = ts[i].Legacy.Back().Value.(string)
		}
		if ts[i].New.Back() != nil {
			result = ts[i].New.Back().Value
		}

		tCase := testCase{
			tsNames[i] + " back",
			expected,
			result,
		}
		tCase.run(t)
	}
}

func TestFront(t *testing.T) {
	ts := testSuite()

	for i := range ts {
		expected := "<nil>"
		result := "<nil>"
		if ts[i].Legacy.Front() != nil {
			expected = ts[i].Legacy.Front().Value.(string)
		}
		if ts[i].New.Front() != nil {
			result = ts[i].New.Front().Value
		}

		tCase := testCase{
			tsNames[i] + " front",
			expected,
			result,
		}
		tCase.run(t)
	}
}

func TestInit(t *testing.T) {
	ts := testSuite()
	for i := range ts {
		expected := containerListToString(ts[i].Legacy.Init())

		tc := testCase{
			strings.Title(tsNames[i]) + " Init() function",
			expected,
			ts[i].New.Init().String(),
		}
		tc.run(t)
		lengthTestCase(fmt.Sprintf("%v Init() length check", strings.Title(tsNames[i])), ts[i], t)
	}
}

func getNthElement(old *legacy.List, nu *list.List[string], index int) (*legacy.Element, *list.Element[string]) {
	// Auxiliar function of the next one
	oldEl := old.Front()
	newEl := nu.Front()
	for i := 0; i < index; i++ {
		oldEl = oldEl.Next()
		newEl = newEl.Next()
	}
	return oldEl, newEl
}

func testInsertCases(
	names []string,
	modifier func(*legacy.List, *legacy.Element, *list.List[string], *list.Element[string]),
	t *testing.T,
) {
	// Common code that runs on any test of list functions that take any *Element
	// as argument.
	indices := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 5},
		{0, 0, 1, 9},
	}

	for i := range indices {
		ts := testSuite() // Lazy reset
		for j := range ts {
			if j == 0 {
				continue // Empty lists don't have elements to work with
			}

			oldEl, newEl := getNthElement(ts[j].Legacy, ts[j].New, indices[i][j])
			modifier(ts[j].Legacy, oldEl, ts[j].New, newEl)

			stringifyTestCase(fmt.Sprintf(names[i], tsNames[j]), ts[j], t)
			lengthTestCase(fmt.Sprintf("Checking length of %v", tsNames[j]), ts[j], t)
		}
	}
}

func TestInsertAfter(t *testing.T) {
	ts := testSuite()

	for i := range ts {
		if i == 0 {
			continue // Can't insert anything in an empty list
		}
		ts[i].Legacy.InsertAfter("New Node", ts[i].Legacy.Front())
		ts[i].New.InsertAfter("New Node", ts[i].New.Front())

		tc := testCase{
			fmt.Sprintf("Insert new node after the first element in %v", tsNames[i]),
			containerListToString(ts[i].Legacy),
			ts[i].New.String(),
		}
		tc.run(t)
	}

	names := []string{
		"Insert new element after the 1st element of %v",
		"InsertAfter() new element in the middle of %v",
		"Insert new element after the last element of %v",
	}
	modifier := func(ol *legacy.List, oe *legacy.Element, nl *list.List[string], ne *list.Element[string]) {
		ol.InsertAfter("New Element", oe)
		nl.InsertAfter("New Element", ne)
	}

	testInsertCases(names, modifier, t)
}

func TestInsertBefore(t *testing.T) {
	ts := testSuite()

	for i := range ts {
		if i == 0 {
			continue // Can't insert anything in an empty list
		}
		ts[i].Legacy.InsertBefore("New Element", ts[i].Legacy.Front())
		ts[i].New.InsertBefore("New Element", ts[i].New.Front())

		tc := testCase{
			fmt.Sprintf("Insert new element before the first element in %v", tsNames[i]),
			containerListToString(ts[i].Legacy),
			ts[i].New.String(),
		}
		tc.run(t)
	}

	names := []string{
		"Insert new element before the 1st element of %v",
		"InsertBefore() new element in the middle of %v",
		"Insert new element before the last element of %v",
	}
	modifier := func(ol *legacy.List, oe *legacy.Element, nl *list.List[string], ne *list.Element[string]) {
		ol.InsertBefore("New Element", oe)
		nl.InsertBefore("New Element", ne)
	}

	testInsertCases(names, modifier, t)
}

func TestLen(t *testing.T) {
	ts := testSuite()
	expectedLengths := []string{"0", "1", "2", "10"}
	for i := range ts {
		tc := testCase{
			tsNames[i] + " length",
			expectedLengths[i],
			strconv.Itoa(ts[i].New.Len()),
		}
		tc.run(t)
	}
}

func testMoveCases(
	cmd func(listTuple[string], *legacy.Element, *legacy.Element, *list.Element[string], *list.Element[string]),
	verb string,
	t *testing.T,
) {
	from := []int{0, 3, 9}
	to := []int{0, 3, 9}
	fromS := []string{"the first element", "the 4th element", "the last element"}
	toS := []string{"the beginning", "the 4th place", "the end"}

	for i := range from {
		for j := range to {
			tl := testLists() // All tests are performed only on the test list
			oFrom, nFrom := getNthElement(tl.Legacy, tl.New, from[i])
			oTo, nTo := getNthElement(tl.Legacy, tl.New, to[j])
			cmd(tl, oFrom, oTo, nFrom, nTo)

			stringifyTestCase(
				fmt.Sprintf("Move %v %v %v of the list", fromS[i], verb, toS[j]),
				tl,
				t,
			)
			lengthTestCase("Cheking test list length", tl, t)
		}
	}
}

func TestMoveAfter(t *testing.T) {
	testMoveCases(
		func(tl listTuple[string], of, ot *legacy.Element, nf, nt *list.Element[string]) {
			tl.Legacy.MoveAfter(of, ot)
			tl.New.MoveAfter(nf, nt)
		},
		"after",
		t,
	)
}

func TestMoveBefore(t *testing.T) {
	testMoveCases(func(tl listTuple[string], of, ot *legacy.Element, nf, nt *list.Element[string]) {
		tl.Legacy.MoveBefore(of, ot)
		tl.New.MoveBefore(nf, nt)
	},
		"before",
		t,
	)
}

func testMoveToCases(
	cmd func(listTuple[string], *legacy.Element, *list.Element[string]),
	verb string,
	t *testing.T,
) {
	places := []int{0, 3, 6, 9}
	placesS := []string{"the first element", "the 4th element", "the 7th element", "the last element"}
	for i := range places {
		tl := testLists()
		oEl, nEl := getNthElement(tl.Legacy, tl.New, places[i])
		cmd(tl, oEl, nEl)
		stringifyTestCase(
			fmt.Sprintf("Moving %v to %v in test list", placesS[i], verb),
			tl,
			t,
		)
		lengthTestCase("Cheking test list length", tl, t)
	}
}

func TestMoveToBack(t *testing.T) {
	testMoveToCases(
		func(tl listTuple[string], oe *legacy.Element, ne *list.Element[string]) {
			tl.Legacy.MoveToBack(oe)
			tl.New.MoveToBack(ne)
		}, "back", t)
}

func TestMoveToFront(t *testing.T) {
	testMoveToCases(
		func(tl listTuple[string], oe *legacy.Element, ne *list.Element[string]) {
			tl.Legacy.MoveToFront(oe)
			tl.New.MoveToFront(ne)
		}, "front", t)
}

func testPushCases(cmd func(tl listTuple[string]), adverb string, t *testing.T) {
	ts := testSuite()
	for i := range ts {
		cmd(ts[i])
		stringifyTestCase(fmt.Sprintf("Push %v new element in %v", adverb, tsNames[i]), ts[i], t)
		lengthTestCase(fmt.Sprintf("Cheking length of %v", tsNames[i]), ts[i], t)
	}
}

func testPushListCases(cmd func(listTuple[string], *legacy.List, *list.List[string]), adverb string, t *testing.T) {
	tl := testLists()
	listSlice := []string{"Another", "Completely", "Different", "List"}
	secondNewList := list.New(listSlice...)
	secondOldList := legacy.New()
	for _, i := range listSlice {
		secondOldList.PushBack(i)
	}
	cmd(tl, secondOldList, secondNewList)

	stringifyTestCase(fmt.Sprintf("Push%vList()", adverb), tl, t)
	lengthTestCase("Cheking test list length", tl, t)
}

func TestPushBack(t *testing.T) {
	testPushCases(func(tl listTuple[string]) {
		tl.Legacy.PushBack("New Node")
		tl.New.PushBack("New Node")
	}, "back", t)
}
func TestPushBackList(t *testing.T) {
	testPushListCases(func(tl listTuple[string], ol *legacy.List, nl *list.List[string]) {
		tl.Legacy.PushBackList(ol)
		tl.New.PushBackList(nl)
	}, "Back", t)
}

func TestPushFront(t *testing.T) {
	testPushCases(func(tl listTuple[string]) {
		tl.Legacy.PushFront("New Node")
		tl.New.PushFront("New Node")
	}, "front", t)
}

func TestPushFrontList(t *testing.T) {
	testPushListCases(func(tl listTuple[string], ol *legacy.List, nl *list.List[string]) {
		tl.Legacy.PushFrontList(ol)
		tl.New.PushFrontList(nl)
	}, "Front", t)
}

func TestRemove(t *testing.T) {
	ts := testSuite()

	for i := range ts {
		if i == 0 {
			continue // Unable to remove elements from empty string
		}
		oe := ts[i].Legacy.Front()
		ne := ts[i].New.Front()

		ts[i].Legacy.Remove(oe)
		ts[i].New.Remove(ne)

		stringifyTestCase(fmt.Sprintf("Removing the first element of %v", tsNames[i]), ts[i], t)
		lengthTestCase(fmt.Sprintf("Cheking length of %v", tsNames[i]), ts[i], t)
	}

	positions := []int{3, 6, 9}
	positionsS := []string{"the 4th element", "the 7th element", "the last element"}
	for i := range positions {
		tl := testLists()

		oe, ne := getNthElement(tl.Legacy, tl.New, positions[i])
		tl.Legacy.Remove(oe)
		tl.New.Remove(ne)

		stringifyTestCase(
			fmt.Sprintf("Remove %v of %v", positionsS[i], tsNames[i]),
			tl,
			t,
		)
		lengthTestCase(fmt.Sprintf("Cheking length of %v", tsNames[i]), tl, t)
	}

}
