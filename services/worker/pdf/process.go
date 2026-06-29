package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func hasImages(ctx *model.Context) (bool, error) {

	for i := 1; i <= ctx.PageCount; i++ {

		pageDict, _, _, err := ctx.PageDict(i, true)
		if err != nil {
			return false, err
		}

		resObj, found := pageDict["Resources"]
		if !found {
			continue
		}

		resDict, ok := resObj.(types.Dict)
		if !ok {
			continue
		}

		xObjectObj, found := resDict["XObject"]
		if !found {
			continue
		}

		xObjects, ok := xObjectObj.(types.Dict)
		if !ok {
			continue
		}

		for _, v := range xObjects {

			ref, ok := v.(types.IndirectRef)
			if !ok {
				continue
			}

			obj, err := ctx.DereferenceDict(ref)
			if err != nil {
				continue
			}

			subtype := obj.NameEntry("Subtype")

			if subtype != nil && *subtype == "Image" {
				return true, nil
			}
		}
	}

	return false, nil
}
