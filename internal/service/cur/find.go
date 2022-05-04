package cur

import (
	"github.com/PixarV/aws-sdk-go/aws"
	"github.com/PixarV/aws-sdk-go/service/costandusagereportservice"
)

func FindReportDefinitionByName(conn *costandusagereportservice.CostandUsageReportService, name string) (*costandusagereportservice.ReportDefinition, error) {
	input := &costandusagereportservice.DescribeReportDefinitionsInput{}

	var result *costandusagereportservice.ReportDefinition

	err := conn.DescribeReportDefinitionsPages(input, func(page *costandusagereportservice.DescribeReportDefinitionsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, reportDefinition := range page.ReportDefinitions {
			if reportDefinition == nil {
				continue
			}

			if aws.StringValue(reportDefinition.ReportName) == name {
				result = reportDefinition
				return false
			}
		}
		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
