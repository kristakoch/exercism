package erratum

// Use opens a resource and calls Frob on it,
// handling errors properly.
func Use(o ResourceOpener, input string) (err error) {
	var res Resource

	res, err = o()

	// Rerun the ResourceOpener in the case of TransientErrors.
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}
	defer res.Close()

	// Defer the recovery, assigning the error to the
	// returned one. FrobErrors receive special handling.
	defer func() {
		if r := recover(); nil != r {
			switch rerr := r.(type) {
			case FrobError:
				res.Defrob(rerr.defrobTag)
			}
			err = r.(error)
		}
	}()

	res.Frob(input)

	return nil
}
