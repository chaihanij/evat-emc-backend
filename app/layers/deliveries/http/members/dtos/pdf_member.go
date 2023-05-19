package dtos

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

type ResponseCertificateNew ResponseCertificate

func (res *ResponseCertificateNew) Parse(c *gin.Context, input *entities.Member) *ResponseCertificateNew {
	copier.Copy(res, input)

	var templ *template.Template
	var err error
	data := struct {
		Prefix    string
		Firstname string
		Lastname  string
	}{
		Prefix:    res.Prefix,
		Firstname: res.FirstName,
		Lastname:  res.LastName,
	}

	if templ, err = template.ParseFiles("/var/app/template/index.html"); err != nil {
		// fmt.Println("err ;", err)
		log.WithError(err).Errorln("Parse File Error")

	}

	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		// fmt.Println("err ;", err)
		log.WithError(err).Errorln("Execute template")

	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.WithError(err).Errorln("New PDF Generator Error")

	}
	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.WithError(err).Errorln("Create PDF Error")

	}

	return res

}
