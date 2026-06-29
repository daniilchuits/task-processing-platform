package pdf

import (
	"log"
	"os"
	"worker/domain"
	"worker/interfaces"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFUpdate struct {
	Switcher interfaces.Switcher
	Update   interfaces.PDFUpdater
}

func (upd PDFUpdate) Work(userId int, filepath string) error {
	log.Println(1)
	if err := upd.Switcher.StatusProcessing(userId, filepath); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	log.Println(2)
	f, err := os.Open(filepath)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}
	log.Println("Stat:", stat)

	log.Println(3)
	ctx, err := api.ReadContext(f, nil)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	log.Println("ctx.PageCount =", ctx.PageCount)

	if ctx.XRefTable != nil {
		log.Println("xref.PageCount =", ctx.XRefTable.PageCount)
	}

	log.Println(4)
	pagesNum := ctx.PageCount
	log.Println("PageCount =", pagesNum)

	hasImage, err := hasImages(ctx)
	if err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}

	log.Println("HasImage =", hasImage)

	data := domain.PDFData{
		UserId:   userId,
		Filepath: filepath,
		Pages:    pagesNum,
		HasImage: hasImage,
	}

	log.Printf("PDFData: %+v\n", data)
	if err = upd.Update.PdfUpdate(data); err != nil {
		upd.Switcher.StatusFail(userId, filepath)
		return err
	}
	return nil
}
