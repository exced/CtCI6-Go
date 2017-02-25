package chapter2

func removeDupsWithBuffer(l *List) {
	buffer := make(map[int]bool)
	for e := l.Front(); e != nil; e = e.Next() {
		if buffer[e.Value] {
			l.MoveAfter(e, e.Next())
		} else {
			buffer[e.Value] = true
		}
	}
}
