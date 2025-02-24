package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Validator[T any] func(T) error

type SimilarMatches struct {
	StringA        string
	StringB        string
	CutoffScore    float64
	ShortenedMatch bool
}

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | ~float32 | ~float64
}

// @ || Validator functions ||
func Validate[T any](data T, validators ...Validator[T]) error {
	for _, validator := range validators {
		err := validator(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func MustNotBeEmptyString(s string) error {
	if utf8.RuneCountInString(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("string cannot be empty")
	}
	return nil
}

func MustBeEqual[T comparable](actualVal, expectedVal T) error {
	if actualVal != expectedVal {
		return fmt.Errorf("expected value of %v to be equal to %v", actualVal, expectedVal)
	}
	return nil
}

func MustBeGreaterThanOrEqual[T Number](actualVal, minVal T) error {
	if actualVal <= minVal {
		return fmt.Errorf("expected value of %v to be greater than or equal to %v", actualVal, minVal)
	}
	return nil
}

func MustBeGreaterThan[T Number](actualVal, minVal T) error {
	if minVal > actualVal {
		return fmt.Errorf("expected value of %v to be greater than %v", actualVal, minVal)
	}
	return nil
}

func MustBeInRange[T Number](actualVal, minVal, maxVal T) error {
	if minVal > actualVal || actualVal > maxVal {
		return fmt.Errorf("expected value of %v to be between %v and %v", actualVal, minVal, maxVal)
	}
	return nil
}

func (s *SimilarMatches) Validate() error {
	return Validate(s, func(m *SimilarMatches) error {
		return MustNotBeEmptyString(m.StringA)
	},
		func(m *SimilarMatches) error {
			return MustNotBeEmptyString(m.StringB)
		},
		func(m *SimilarMatches) error {
			return MustBeInRange(m.CutoffScore, 0.0, 1.0)
		})
}

//@ || Helper functions ||

// * Need to use RuneCountInString instead of len due to Japanese Characters
func longestWord(strA, strB string) int {
	if utf8.RuneCountInString(strA) > utf8.RuneCountInString(strB) {
		return utf8.RuneCountInString(strA)
	}
	return utf8.RuneCountInString(strB)
}

func shortestSlice(a, b []rune) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}

// @ || Algorithmic functions ||
func calculateJaroWinklerScore(strA, strB string) (float64, error) {
	score, err := calculateJaroScore(strA, strB)
	if err != nil {
		return 0.0, err
	}
	jw := score + (0.1 * float64(getPrefixLength(strA, strB)) * (1.0 - score))
	return jw, nil
}

func calculateJaroScore(strA, strB string) (float64, error) {
	a := strings.ToLower(strA)
	b := strings.ToLower(strB)
	commonRunesA := jaroScoreCommonRunes(a, b)
	commonRunesB := jaroScoreCommonRunes(b, a)
	transpositionValue := jaroScoreTransposition(commonRunesA, commonRunesB)
	matchingAValue := float64(len(commonRunesA))
	matchingBValue := float64(len(commonRunesB))
	js := 0.33 * ((matchingAValue / float64(utf8.RuneCountInString(a))) +
		(matchingBValue / float64(utf8.RuneCountInString(b))) +
		((matchingAValue - transpositionValue) / matchingAValue))
	return js, nil
}

func jaroScoreCommonRunes(strA, strB string) []rune {
	runesA := []rune(strA)
	runesB := []rune(strB)
	var commonChars []rune
	runeMatched := make([]bool, utf8.RuneCountInString(strB))
	limit := longestWord(strA, strB) / 2
	for i, r := range runesA {
		end := i + limit + 1
		start := i - limit
		if end > utf8.RuneCountInString(strB) {
			end = utf8.RuneCountInString(strB)
		}
		if start < 0 {
			start = 0
		}
		for j := start; j < end; j++ {
			if runeMatched[j] {
				continue
			}
			if r == runesB[j] {
				commonChars = append(commonChars, r)
				runeMatched[j] = true
				break
			}
		}
	}
	return commonChars
}

func jaroScoreTransposition(a, b []rune) float64 {
	t := 0.0
	len := shortestSlice(a, b)
	for i := 0; i < len; i++ {
		if a[i] != b[i] {
			t++
		}
	}
	return t
}

func getPrefixLength(strA, strB string) int {
	runesA := []rune(strA)
	runesB := []rune(strB)
	p := 0
	len := shortestSlice(runesA, runesB)
	for i := 0; i < len; i++ {
		if (runesA[i] != runesB[i]) || (p == 4) {
			break
		}
		p++
	}
	return p
}

// @ || General functions ||

// @ if first rune != passed acceptable first rune break
// @ something like first rune in value of [s]
func IsSimilar(sm SimilarMatches) (bool, float64, error) {
	if sm.ShortenedMatch && sm.StringA[0] != sm.StringB[0] {
		return false, 0.0, nil
	}
	score, err := calculateJaroWinklerScore(sm.StringA, sm.StringB)
	if err != nil {
		return false, 0.0, err
	}
	if score >= sm.CutoffScore {
		return true, score, nil
	}

	return false, 0.0, nil
}

// @ add acceptable first rune
func NewQuery(strA, strB string, cutoffScore float64, shortenedMatch bool) SimilarMatches {
	return SimilarMatches{
		StringA:        strA,
		StringB:        strB,
		CutoffScore:    cutoffScore,
		ShortenedMatch: shortenedMatch,
	}
}
