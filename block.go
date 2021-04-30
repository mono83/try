package try

// Block defines traditional try-catch-finally block
// implemented in Golang.
// Original https://dzone.com/articles/try-and-catch-in-golang
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception is an alias for interface{}
type Exception interface{}

// Throw throws an exception (anything)
func Throw(up Exception) {
	panic(up)
}

// Do performs operation using try-catch-finally
func (b Block) Do() {
	if b.Finally != nil {
		defer b.Finally()
	}
	if b.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}
	b.Try()
}
