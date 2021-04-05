package timeutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Clock_Slowness(t *testing.T) {
	theTime := time.Now()
	clock := NewClock(theTime)
	slowness := clock.Slowness

	assert.NotEmpty(t, slowness, "slowness")
	fmt.Println(slowness)
}

func Test_Clock_Now(t *testing.T) {
	theTime := time.Now()
	clock := NewClock(theTime)
	now := clock.Now()

	assert.NotEmpty(t, now, "now")
	fmt.Println(now)
}

func Test_Clock_AlarmAt(t *testing.T) {
	theTime := time.Now()
	clock := NewClock(theTime)
	ch := clock.AlarmAt(time.Now().Add(5 * time.Second))

	fmt.Println(time.Now())
	fmt.Println("waiting...")
	attime := <-ch

	assert.NotEmpty(t, attime, "attime")
	fmt.Println(attime)
}

func Test_Time_MustParseTime(t *testing.T) {

	var theTime time.Time

	theTime, _ = time.Parse("15시 04분", "14시 00분")
	assert.Equal(t, theTime, MustParseTime("오후2시"))
	assert.Equal(t, theTime, MustParseTime("오후2시 00분"))
	assert.Equal(t, theTime, MustParseTime("14시 00분 "))
	assert.Equal(t, theTime, MustParseTime("14시 00분"))
	assert.Equal(t, theTime, MustParseTime("14시"))
}
