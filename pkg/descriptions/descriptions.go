package descriptions

import (
	"fmt"

	"github.com/Checkmarx/kics/pkg/model"
)

var (
	descClient HTTPDescription = &Client{}
)

// RequestAndOverrideDescriptions - Requests CIS descriptions and override default descriptions
func RequestAndOverrideDescriptions(summary *model.Summary) error {
	descriptionIDs := make([]string, 0)
	for idx := range summary.Queries {
		descriptionIDs = append(descriptionIDs, summary.Queries[idx].DescriptionID)
	}

	descriptionMap, err := descClient.RequestDescriptions(descriptionIDs)
	if err != nil {
		return err
	}

	for idx := range summary.Queries {
		if descriptionMap[summary.Queries[idx].DescriptionID].DescriptionID == "" {
			continue
		}
		descriptionID := summary.Queries[idx].DescriptionID

		summary.Queries[idx].CISDescriptionID = descriptionMap[descriptionID].DescriptionID
		summary.Queries[idx].CISDescriptionTitle = descriptionMap[descriptionID].DescriptionTitle
		summary.Queries[idx].CISDescriptionText = descriptionMap[descriptionID].DescriptionText
		summary.Queries[idx].CISRationaleText = descriptionMap[descriptionID].RationaleText
		summary.Queries[idx].CISBenchmarkName = descriptionMap[descriptionID].BenchmarkName
		summary.Queries[idx].CISBenchmarkVersion = descriptionMap[descriptionID].BenchmarkVersion

		summary.Queries[idx].CISDescriptionIDFormatted = fmt.Sprintf(
			"CIS Security - %s v%s - Rule %s",
			descriptionMap[descriptionID].BenchmarkName,
			descriptionMap[descriptionID].BenchmarkVersion,
			descriptionMap[descriptionID].DescriptionID,
		)
		summary.Queries[idx].CISDescriptionTextFormatted = fmt.Sprintf(
			"%s %s",
			descriptionMap[descriptionID].DescriptionText,
			descriptionMap[descriptionID].RationaleText,
		)
	}
	return nil
}
