package mocking

import (
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	testSpy := MockSpy{}

	want := []Event{
		{Name: "write", Content: "3\n"},
		{Name: "sleep", Content: "1s"},
		{Name: "write", Content: "2\n"},
		{Name: "sleep", Content: "1s"},
		{Name: "write", Content: "1\n"},
		{Name: "sleep", Content: "1s"},
		{Name: "write", Content: "Go!"},
	}

	countdown(&testSpy)

	assertSpy(t, testSpy, want)
}

func assertSpy(t testing.TB, m MockSpy, events []Event) {
	t.Helper()
	actualEvents := len(m.Events)
	requiredEvents := len(events)
	if actualEvents != requiredEvents {
		t.Errorf("exptected %d but got %d events", requiredEvents, actualEvents)
	}

	for i := range actualEvents {
		if m.Events[i] != events[i] {
			t.Errorf("exprected %s but got %s", events[i], m.Events[i])
		}
	}
}

type MockSpy struct {
	Events []Event
}

func (m *MockSpy) SpySleep(t time.Duration) {
	name := "sleep"
	content := t.String()
	event := Event{Name: name, Content: content}
	m.Events = append(m.Events, event)
}

func (m *MockSpy) SpyWrite(s string) {
	name := "write"
	event := Event{Name: name, Content: s}
	m.Events = append(m.Events, event)
}
