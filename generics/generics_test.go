package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "greetings")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(StackOfInts)

		// check stack is empty
		AssertTrue(t, stackOfInts.IsEmpty())

		// add a thing, then check its not empty
		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		// add another thing, pop it back again
		stackOfInts.Push(456)
		value, ok := stackOfInts.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, value, 456)

		value, ok = stackOfInts.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, value, 123)

		AssertTrue(t, stackOfInts.IsEmpty())

	})
	t.Run("string stack", func(t *testing.T) {
		stackOfStrings := new(StackOfStrings)

		// check stack is empty
		AssertTrue(t, stackOfStrings.IsEmpty())

		// add a thing, then check its not empty
		stackOfStrings.Push("123")
		AssertFalse(t, stackOfStrings.IsEmpty())

		// add another thing, pop it back again
		stackOfStrings.Push("456")
		value, ok := stackOfStrings.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, value, "456")

		value, ok = stackOfStrings.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, value, "123")

		AssertTrue(t, stackOfStrings.IsEmpty())
	})

	t.Run("interface stack DX is horrid", func(t *testing.T) {
		stackOfInts := new(StackWithoutGenericsOfInts)
		stackOfInts.Push(1)
		stackOfInts.Push(2)

		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		// AssertEqual(t, firstNum+secondNum, 3) THIS DOESNT WORK BEACUSE YOU CANT ADD ANY WITH ANY

		// get our ints from out interface{}
		reallyFirstNum, ok := firstNum.(int)
		AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

		reallySecondNum, ok := secondNum.(int)
		AssertTrue(t, ok) // and again!

		AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	})

	t.Run("integer stack with generics", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()

	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()

	if got {
		t.Errorf("got %v, want false", got)
	}
}
