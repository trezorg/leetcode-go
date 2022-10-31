package substring

type circleBuffer struct {
	start  int
	end    int
	length int
	values []int
	cap    int
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *circleBuffer) addPos(pos int) int {
	prevPos := c.values[c.start]
	c.values[c.end] = pos
	if c.start == c.end {
		c.start = (c.start + 1) % c.cap
	}
	c.end = (c.end + 1) % c.cap
	c.length = min(c.length+1, c.cap)
	return prevPos
}

func (c *circleBuffer) min() int {
	return c.values[c.start]
}

func (c *circleBuffer) isFull() bool {
	return c.length == c.cap
}

func newcircleBuffer(cap int) *circleBuffer {
	return &circleBuffer{
		cap:    cap,
		values: make([]int, cap, cap),
		length: 0,
		start:  0,
	}
}

type posDict struct {
	pos    map[byte]*circleBuffer
	minPos int
}

func (p *posDict) addPos(pos int, v byte) {
	cb, ok := p.pos[v]
	if !ok {
		return
	}
	prevPos := cb.addPos(pos)
	if prevPos == p.minPos {
		var min int
		for _, t := range p.pos {
			min = t.min()
			break
		}
		for _, t := range p.pos {
			if min > t.min() {
				min = t.min()
			}
		}
		p.minPos = min
	}
}

func (p posDict) found() bool {
	for _, v := range p.pos {
		if !v.isFull() {
			return false
		}
	}
	return true
}

func newPosDict(t string) posDict {
	symbols := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		symbols[t[i]]++
	}
	pos := make(map[byte]*circleBuffer)
	for k, v := range symbols {
		pos[k] = newcircleBuffer(v)
	}
	return posDict{
		minPos: 0,
		pos:    pos,
	}
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	pd := newPosDict(t)
	l, r, lm, rm := 0, 0, 0, 0
	for r < len(s) && !pd.found() {
		pd.addPos(r, s[r])
		r++
	}
	if !pd.found() {
		return ""
	}
	l = pd.minPos
	lm, rm = l, r
	for r < len(s) {
		pd.addPos(r, s[r])
		r++
		if l < pd.minPos {
			l = pd.minPos
			if r-l < rm-lm {
				lm = l
				rm = r
			}
		}
	}
	return s[lm:rm]
}
