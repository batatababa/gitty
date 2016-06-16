package cli

// Option Option data structure
type Option struct {
	ShortName   string
	LongName    string
	Value       string
	Description string
}

//OptionSlice Slice of Option type def
type OptionSlice []Option

func (slice OptionSlice) toColumnArray() (colArr [][]string) {
	if slice != nil {
		for _, opt := range slice {
			colArr = append(colArr, []string{opt.ShortName, opt.LongName, opt.Description, opt.Value})
		}
	}

	return
}
