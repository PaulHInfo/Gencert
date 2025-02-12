package certpdf

import (
	"fmt"
	"gencert/cert"
	"path"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	/*err := os.Mkdir(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}*/
	p = &PdfSaver{OutputDir: outputDir}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()
	//design (Worst part)
	background(pdf)
	header(pdf, &cert)
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "i", 40)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	footer(pdf)
	//save
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Println("PDF saved")
	return nil
}
func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	w, h := pdf.GetPageSize()
	pdf.ImageOptions("img/bg.png", 0, 0, w, h, false, opts, 0, "")

}
func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	m := 30.0
	x := 0.0
	w := 30.0
	fileName := "img/smile.png"
	pdf.ImageOptions(fileName, x+m, 20, w, 0, false, opts, 0, "")
	wp, _ := pdf.GetPageSize()
	x = wp - w
	pdf.ImageOptions(fileName, x-m, 20, w, 0, false, opts, 0, "")
	pdf.SetFont("Arial", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}
func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	iw := 50.0
	filename := "img/valide.png"
	pw, ph := pdf.GetPageSize()
	x := pw - iw - 20.0
	y := ph - iw - 10
	pdf.ImageOptions(filename, x, y, iw, 0, false, opts, 0, "")
}
