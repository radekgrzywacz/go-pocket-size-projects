package habit

import (
	"strings"
	"time"
	
	"github.com/google/uuid"

)

func validateAndCompleteHabit(h Habit) (Habit, error) {
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{field: "name", reason: "is empty"}
	}

	if h.WeeklyFrequency == 0 {
		h.WeeklyFrequency = 1
	}

	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h, nil
}
