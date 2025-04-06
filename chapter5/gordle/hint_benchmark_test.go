package gordle

import "testing"

func BenchmarkFeedback_StringConcat1(b *testing.B) {
	fb := feedback{absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkFeedback_StringConcat2(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkFeedback_StringConcat3(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkFeedback_StringConcat4(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkFeedback_StringConcat5(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.StringConcat()
	}
}

func BenchmarkStringBuilder1(b *testing.B) {
	fb := feedback{absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder2(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder3(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder4(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter}
	for n := 0; n < b.N; n++ {
		_ = fb.String()
	}
}

func BenchmarkStringBuilder5(b *testing.B) {
	fb := feedback{absentCharacter, wrongPosition, correctPosition, absentCharacter, wrongPosition}
	for n := 0; n < b.N; n++ {
		_ = fb.String()

	}
}
