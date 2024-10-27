package react

type reactor struct {
	registry []*computeCell
}
type inputCell struct {
	value int
	r     *reactor
}
type computeCell struct {
	compute   func() int
	callbacks map[int]func(int)
	prevVal   int
}
type canceler struct {
	id           int
	registeredTo *computeCell
}

// New is a factory for Reactors.
func New() Reactor {
	var r reactor

	r.registry = make([]*computeCell, 0)

	return &r
}

// CreateInput creates an input cell.
func (r *reactor) CreateInput(i int) InputCell {
	var ic inputCell

	ic.value = i
	ic.r = r

	return &ic
}

// Createcompute creates a compute 1 cell.
func (r *reactor) CreateCompute1(
	c Cell,
	compute func(int) int,
) ComputeCell {
	return r.createCompute(
		func() int { return compute(c.Value()) },
	)
}

// Createcompute creates a compute 2 cell.
func (r *reactor) CreateCompute2(
	c1, c2 Cell,
	compute func(int, int) int,
) ComputeCell {
	return r.createCompute(
		func() int { return compute(c1.Value(), c2.Value()) },
	)
}

func (r *reactor) createCompute(compute func() int) *computeCell {
	var cc computeCell

	cc.compute = compute
	cc.callbacks = make(map[int]func(int))
	cc.prevVal = cc.Value()
	r.registry = append(r.registry, &cc)

	return &cc
}

// SetValue sets the value on an input cell.
func (ic *inputCell) SetValue(i int) {
	ic.value = i
	ic.r.update()
}

// update triggers callbacks and prevVal updates in the reactor.
func (r *reactor) update() {
	for _, cc := range r.registry {
		if len(cc.callbacks) == 0 || cc.Value() == cc.prevVal {
			continue
		}
		for _, cb := range cc.callbacks {
			cb(cc.Value())
		}
		cc.prevVal = cc.Value()
	}
}

// Value returns the value of an input cell.
func (ic *inputCell) Value() int {
	return ic.value
}

// Value returns the value of a compute cell.
func (cc *computeCell) Value() int {
	return cc.compute()
}

// AddCallback adds a callback function to a compute cell.
func (cc *computeCell) AddCallback(cb func(int)) Canceler {
	var c canceler

	c.id = len(cc.callbacks)
	cc.callbacks[c.id] = cb
	c.registeredTo = cc

	return &c
}

// Cancel removes a callback function from a compute cell.
func (c *canceler) Cancel() {
	delete(c.registeredTo.callbacks, c.id)
}
