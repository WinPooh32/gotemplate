package ringfixed

// template type Ring(DType)
type DType int

// Ring is a fixed-size circular buffer.
type Ring struct {
	buf []DType
	cap int
	len int
	ptr int
}

// Make makes a new instance of the ring buffer with given capacity.
func Make(cap int) Ring {
	return Ring{
		buf: make([]DType, cap),
		cap: cap,
		len: 0,
		ptr: 0,
	}
}

// Len returns count of items.
func (r *Ring) Len() int {
	return r.len
}

// Cap returns maximum capacity of the ring buffer.
func (r *Ring) Cap() int {
	return r.cap
}

// Full returns true if the buffer is full.
func (r *Ring) Full() bool {
	return r.len >= r.cap
}

// ForcePushBack adds new item to the end of buffer
// and deletes first item if buffer is full.
func (r *Ring) ForcePushBack(value DType) {
	if r.Full() {
		r.PopFront()
	}
	r.PushBack(value)
}

// Back returns the last item of the buffer.
func (r *Ring) Back() (value DType) {
	return r.buf[(r.ptr+r.len-1)%r.cap]
}

// Front returns the first item of the buffer.
func (r *Ring) Front() (value DType) {
	return r.buf[r.ptr]
}

// PushBack adds new item to the end of the buffer.
func (r *Ring) PushBack(value DType) {
	end := (r.ptr + r.len) % r.cap
	r.buf[end] = value
	r.len++
}

// PopFront deletes first item and returns it.
func (r *Ring) PopFront() (value DType) {
	value = r.buf[r.ptr]
	r.ptr = (r.ptr + 1) % r.cap
	r.len--
	return value
}

// Copy makes copy of the ring buffer to dst slice.
// Returns count of copied items.
func (r *Ring) CopyTo(dst []DType) (n int) {
	if len(dst) < r.cap {
		panic("dst length must be equal or greater ring size")
	}
	end := r.ptr + r.len
	if end <= r.cap {
		copy(dst, r.buf[r.ptr:end])
		return r.len
	}
	p := end % r.cap
	copy(dst, r.buf[r.ptr:r.cap])
	copy(dst[r.cap-r.ptr:], r.buf[0:p])
	return r.cap - r.ptr + p
}
