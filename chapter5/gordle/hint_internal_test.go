package gordle

import "testing"

func TestFeedback_String(t *testing.T) {
	testCases := map[string]struct {
		fb   feedback
		want string
	}{
		"thee correct": {
			fb:   feedback{correctPosition, correctPosition, correctPosition},
			want: "💚💚💚",
		},
		"one of each": {
			fb:   feedback{correctPosition, wrongPosition, absentCharacter},
			want: "💚🟡⬜",
		},
		"different order for one of each": {
			fb:   feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟡⬜💚",
		},
		"unknown position": {
			fb:   feedback{hint(88)},
			want: "💔",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.String()
			if got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}
