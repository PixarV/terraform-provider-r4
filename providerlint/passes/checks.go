package passes

import (
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT001"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT002"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT003"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT004"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT005"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSAT006"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSR001"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSR002"
	"github.com/PixarV/terraform-provider-ritt/providerlint/passes/AWSV001"
	"golang.org/x/tools/go/analysis"
)

var AllChecks = []*analysis.Analyzer{
	AWSAT001.Analyzer,
	AWSAT002.Analyzer,
	AWSAT003.Analyzer,
	AWSAT004.Analyzer,
	AWSAT005.Analyzer,
	AWSAT006.Analyzer,
	AWSR001.Analyzer,
	AWSR002.Analyzer,
	AWSV001.Analyzer,
}
