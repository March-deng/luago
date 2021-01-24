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
func (l *luaState) Insert() {

}

//TODO: wait rotate to be implemented
func (l *luaState) Remove() {

}

func (l *luaState) Rotate(idx, n int) {
	//
	//t := l.stack.top - 1
	//i := l.stack.absIndex(idx) - 1
}

func (l *luaState) SetTop(idx int) {

}
