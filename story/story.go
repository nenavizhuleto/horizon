package story

import (
	"container/list"
	"fmt"
)

type Story[About any] struct {
	started bool
	ended   bool

	chronology *list.List
}

func New[About any]() *Story[About] {
	return &Story[About]{
		chronology: list.New(),
	}
}

func (s *Story[About]) value(el *list.Element) About {
	return el.Value.(About)
}

func (s *Story[About]) Explain() []About {
	var (
		explanation = make([]About, 0)
		beginning   = s.beginning()
		ending      = s.ending()
		chapter     *list.Element
	)

	explanation = append(explanation, s.Beginning())

	chapter = beginning.Next()
	for chapter != nil && chapter != ending {
		explanation = append(explanation, s.value(chapter))
		chapter = chapter.Next()
	}

	explanation = append(explanation, s.Ending())

	return explanation
}

func (s *Story[About]) TimeSpan() int {
	return s.chronology.Len()
}

func (s *Story[About]) Repeat(repeat func(About) bool) {
	if !s.started {
		return
	}

	chapter := s.beginning()
	s.End()
	for chapter != nil {
		if !repeat(s.value(chapter)) {
			break
		}
		chapter = chapter.Next()
	}
}

func (s *Story[About]) Begin() *Story[About] {
	s.started = true
	s.ended = false
	return s
}

func (s *Story[About]) BeginWith(beginning About) *Story[About] {
	if !s.started {
		s.Begin()
		s.chronology.PushBack(beginning)
	}
	return s
}

func (s *Story[About]) beginning() *list.Element {
	return s.chronology.Front()
}

func (s *Story[About]) Beginning() About {
	return s.value(s.beginning())
}

func (s *Story[About]) ContinueWith(chapter About) *Story[About] {
	if s.started && !s.ended {
		s.chronology.PushBack(chapter)
	}
	return s
}

func (s *Story[About]) End() *Story[About] {
	s.ended = true
	return s
}

func (s *Story[About]) EndWith(ending About) *Story[About] {
	if !s.ended {
		s.End()
		s.chronology.PushBack(ending)
	}
	return s
}

func (s *Story[About]) ending() *list.Element {
	return s.chronology.Back()
}

func (s *Story[About]) Ending() About {
	return s.value(s.ending())
}

func (s *Story[About]) Describe() string {
	var (
		about     string = "\n"
		beginning *list.Element
		ending    *list.Element
		chapter   *list.Element
	)

	if s.chronology.Len() == 0 {
		return about
	}

	if s.started {
		beginning = s.chronology.Front()
	}

	if s.ended {
		ending = s.chronology.Back()
	}

	about = fmt.Sprintf("beginning: %+v\n", s.value(beginning))

	chapter = beginning.Next()
	for chapter != nil && chapter != ending {
		text := s.value(chapter)
		about = fmt.Sprintf("%schapter: %+v\n", about, text)
		chapter = chapter.Next()
	}

	about = fmt.Sprintf("%sending: %+v\n", about, s.value(ending))

	return about
}
