package dtos

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type ResponseCertificateNew ResponseCertificate

// func (res *ResponseCertificateNew) Parse(c *gin.Context, input *entities.Member) *ResponseCertificateNew {
// 	copier.Copy(res, input)

// 	var templ *template.Template
// 	var err error
// 	data := struct {
// 		Prefix    string
// 		Firstname string
// 		Lastname  string
// 	}{
// 		Prefix:    res.Prefix,
// 		Firstname: res.FirstName,
// 		Lastname:  res.LastName,
// 	}

// 	if templ, err = template.ParseFiles("/var/app/template/index.html"); err != nil {
// 		// fmt.Println("err ;", err)
// 		log.WithError(err).Errorln("Parse File Error")

// 	}

// 	var body bytes.Buffer
// 	if err = templ.Execute(&body, data); err != nil {
// 		// fmt.Println("err ;", err)
// 		log.WithError(err).Errorln("Execute template")

// 	}

// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.WithError(err).Errorln("New PDF Generator Error")

// 	}
// 	// read the HTML page as a PDF page
// 	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

// 	// enable this if the HTML file contains local references such as images, CSS, etc.
// 	page.EnableLocalFileAccess.Set(true)

// 	// add the page to your generator
// 	pdfg.AddPage(page)
// 	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
// 	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
// 	pdfg.Dpi.Set(300)

// 	err = pdfg.Create()
// 	if err != nil {
// 		log.WithError(err).Errorln("Create PDF Error")

// 	}

// 	return res

// }

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

	if templ, err = template.ParseFiles("/app/data/template/index.html"); err != nil {
		fmt.Println("err ;", err)
	}

	buf := new(bytes.Buffer)

	if err = templ.Execute(buf, data); err != nil {
		fmt.Println("err ;", err)
	}

	body := buf.String()

	filename := fmt.Sprintf("%s", res.FirstName)

	mockFile := `/app/data/template/` + filename + `.html`
	if err1 := ioutil.WriteFile(mockFile, []byte(body), 0644); err1 != nil {
		fmt.Println("err1 GenPDF : ", err1.Error())
	}

	f, err := os.Open(mockFile)
	if err != nil {
		fmt.Println("err openFile : ", err.Error())
	}

	defer os.Remove(mockFile)
	defer f.Close()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("err pdfg :", err)
	}
	pageReader := wkhtmltopdf.NewPageReader(f)

	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)

	pageReader.PageOptions.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(pageReader)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		fmt.Println("error2 create PDF : ", err)
	}
	err = pdfg.WriteFile(`/app/data/template/` + filename + `.pdf`)
	if err != nil {
		fmt.Println("err WriteFile ;", err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Disposition", `attachment; filename=`+res.FirstName+".pdf")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(pdfg.Bytes())))
	c.Writer.Write(pdfg.Bytes())

	return res
}
