package state

//GetTop
func (l *luaState) GetTop() int {
	return l.stack.top
}

func (l *luaState) AbsIndex(idx int) int {
	return l.stack.absIndex(idx)
}

func (l *luaState) CheckStack(n int) bool {
	l.stack.check(n)
	//TODO: what if expansion failed
	return true
}

func (l *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		l.stack.pop()
	}

}

func (l *luaState) Copy(from, to int) {
	val := l.stack.get(from)
	l.stack.set(to, val)

}

func (l *luaState) PushValue(idx int) {
	val := l.stack.get(idx)
	l.stack.push(val)
}

func (l *luaState) Replace(idx int) {
	val := l.stack.pop()
	l.stack.set(idx, val)
}

//TODO: wait rotate to be implemented
func (l *luaState) Insert(idx int) {
	l.Rotate(idx, 1)
}

//TODO: wait rotate to be implemented
func (l *luaState) Remove(idx int) {
	l.Rotate(idx, -1)
	l.Pop(1)
}

//TODO: 证明其正确性
func (l *luaState) Rotate(idx, n int) {

	t := l.stack.top - 1
	i := l.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = i - n - 1
	}
	l.stack.reverse(i, m)
	l.stack.reverse(m+1, t)
	l.stack.reverse(i, t)
}

func (l *luaState) SetTop(idx int) {
	newTop := l.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow")
	}
	n := l.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			l.stack.pop()
		}
	} else {
		for i := 0; i > n; i-- {
			l.stack.push(nil)
		}
	}
}

func (l *luaState) PushNil() {
	l.stack.push(nil)
}

func (l *luaState) PushBoolean(b bool) {
	l.stack.push(b)
}

func (l *luaState) PushInteger(n int64) {
	l.stack.push(n)
}

func (l *luaState) PushNumber(n float64) {
	l.stack.push(n)
}

func (l *luaState) PushString(s string) {
	l.stack.push(s)
}
