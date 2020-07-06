package disruptor

import (
	"sync/atomic"
)

type Sequenced interface {
	Next() int64
	NextInc(value int64) int64
	TryNext() int64
	TryNextInc(value int64) int64
	Publish(sequence int64)
}

func NewRingBuffer(bufferSize int, factory func() interface{}) *RingBuffer {
	ringBuffer := &RingBuffer{
		entries:    nil,
		bufferSize: 0,
		sequencer:  nil,
	}
	for i := 0; i < bufferSize; i++ {
		if cap(ringBuffer.entries) > bufferSize {
			ringBuffer.entries[i] = factory()
		} else {
			ringBuffer.entries = append(ringBuffer.entries, factory())
		}
	}
	return ringBuffer
}

type RingBuffer struct {
	entries    []interface{}
	bufferSize int
	sequencer  *MultiProducerSequencer
}

func (r *RingBuffer) Next() int64 {
	return r.sequencer.Next()
}

func (r *RingBuffer) NextInc(value int64) int64 {
	return r.sequencer.NextInc(value)
}

func (r *RingBuffer) PublishEvent(translator func(interface{}, int64)) {
	seq := r.sequencer.Next()
	translator(r.get(seq), seq) //TODO
}

func (r *RingBuffer) get(sequence int64) interface{} {
	return nil
}

type Sequence struct {
	value int64
}

func NewSequence(value int64) *Sequence {
	return &Sequence{
		value: value,
	}
}

func (s *Sequence) Get() int64 {
	return atomic.LoadInt64(&(s.value))
}

func (s *Sequence) Set(v int64) {
	atomic.SwapInt64(&(s.value), v)
}

func (s *Sequence) CompareAndSet(expectedValue int64, newValue int64) bool {
	return atomic.CompareAndSwapInt64(&(s.value), expectedValue, newValue)
}

func (s *Sequence) IncrementAndGet() int64 {
	return s.AddAndGet(1)
}

func (s *Sequence) AddAndGet(increment int64) int64 {
	var newValue int64
	for {
		current := s.Get()
		newValue = current + increment
		if s.CompareAndSet(current, newValue) {
			break
		}
	}
	return newValue
}

type MultiProducerSequencer struct {
	sequence        Sequence
	BufferSize      int
	AvailableBuffer []int
}

func (m *MultiProducerSequencer) Next() int64 {
	return m.NextInc(int64(1))
}

func (m *MultiProducerSequencer) NextInc(value int64) int64 {
	if value < 1 {
		panic("value must be > 0")
	}
	var next int64
	for {
		current := m.sequence.Get()
		next = current + value

		if m.sequence.CompareAndSet(current, next) {
			break
		}
	}
	return next
}

func NewMultiProducerSequencer() *MultiProducerSequencer {

	return nil
}
