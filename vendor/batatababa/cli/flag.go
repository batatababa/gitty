package cli

// Flag Flag data structure
type Flag struct {
	ShortName   string
	LongName    string
	Description string
}

// FlagSlice Slice of Flag type def
type FlagSlice []Flag

func (slice FlagSlice) toColumnArray() (colArr [][]string) {
	if slice != nil {
		for _, flag := range slice {
			colArr = append(colArr, []string{flag.ShortName, flag.LongName, flag.Description})
		}
	}

	return
}
