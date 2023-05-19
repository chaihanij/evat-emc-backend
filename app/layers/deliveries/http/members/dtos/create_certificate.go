package dtos

import (
	"bytes"
	"html/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type RequestCertificate struct {
	MemberUUID string `uri:"member_uuid"`
}

type ResponseCertificate struct {
	Prefix       string    `json:"prefix"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	BirthDay     time.Time `json:"birth_day" `
	NationalId   string    `json:"national_id" `
	Checkin_date time.Time `json:"checkin_date"`
	Is_checkin   bool      `json:"is_checkin"`
	Is_data      bool      `json:"is_data"`
	Is_image     bool      `json:"is_image"`
	Is_national  bool      `json:"is_national"`
}

func (req *RequestCertificate) Parse(c *gin.Context) (*RequestCertificate, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil

}

func (req *RequestCertificate) ToEntity() *entities.MemberFilter {
	return &entities.MemberFilter{
		UUID: &req.MemberUUID,
	}
}

func (res *ResponseCertificate) Parse(c *gin.Context, input *entities.Member) ([]byte, error) {
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
		return nil, err
	}

	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		// fmt.Println("err ;", err)
		log.WithError(err).Errorln("Execute template")
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.WithError(err).Errorln("New PDF Generator Error")
		return nil, err
	}

	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)
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
		return nil, err
	}

	return pdfg.Bytes(), nil

}
