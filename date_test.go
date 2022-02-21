package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {

	t.Run("GIVEN WHEN THEN", func(t *testing.T) {
		testDate := "2022-02-14"
		date := Date{}
		testDates, _ := time.Parse(layout, testDate)
		date.Time = &testDates
		parsedDate, err := date.MarshalJSON()
		parsedDateStr := string(parsedDate)
		println(parsedDateStr)
		asserts := assert.New(t)
		asserts.Nil(err)
		expectedParsedDate := "\"" + testDate + "\""
		asserts.Equal(parsedDateStr, expectedParsedDate)
	})

}
