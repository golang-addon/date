package date

import "time"

const layout = "2006-01-02"

type Date struct {
	*time.Time
}

func (dt Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.Time.Format(layout) + `"`), nil
}

func (dt *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var err error
	value := string(data)
	n := len(value)
	t, err := time.Parse(layout, value[1:n-1])
	dt.Time = &t
	return err
}
