package dtos

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateCertificateRequestJSON struct {
	Fname       string `uri:"fname" form:"fname"`
	Lname       string `uri:"lname" form:"lname"`
	DateCreated string `uri:"date_created" form:"date_created" `
	StartRace   string `uri:"start_race" form:"start_race" `
	StopRace    string `uri:"stop_race" form:"stop_race" `
}

func (req *CreateCertificateRequestJSON) Parse(c *gin.Context) (*CreateCertificateRequestJSON, error) {
	// body := ""
	err := c.ShouldBindUri(req)
	if err != nil {
		return nil, err
	}

	logrus.Debugln("req:", *req)

	templatePath := "http://127.0.0.1:8080/v1/files/04b198cc-05c6-4f1c-b875-35fa80e64679"
	templateCertificate, err := template.ParseFiles(templatePath)
	if err != nil {
		logrus.Debugln("err :", err)
	}
	fmt.Println("templateCertificate :", templateCertificate)
	// buf := new(bytes.Buffer)

	// // data
	// if err = templateCertificate.Execute(buf, *req); err != nil {
	// 	logrus.Debugln("templete error :", err)
	// }
	// body = buf.String()

	// mockFileName := fmt.Sprintf("%s.html", time.Now())

	// mockFile := mockFileName

	// if err1 := ioutil.WriteFile(mockFile, []byte(body), 0644); err1 != nil {
	// 	fmt.Println("err1 GenPDF : ", err1.Error())
	// }

	// f, err := os.Open(mockFile)
	// if err != nil {
	// 	fmt.Println("err openFile : ", err.Error())
	// }

	// defer os.Remove(mockFile)
	// defer f.Close()

	// pdfg, err := wkhtmltopdf.NewPDFGenerator()
	// if err != nil {
	// 	fmt.Println("error1 pdfg : ", err.Error())
	// }

	// pageReader := wkhtmltopdf.NewPageReader(f)
	// pageReader.PageOptions.EnableLocalFileAccess.Set(true)
	// pdfg.AddPage(pageReader)

	// pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// pdfg.Dpi.Set(300)
	// err = pdfg.Create()
	// if err != nil {
	// 	fmt.Println("error2 create PDF : ", err.Error())
	// }

	// filename := fmt.Sprintf("%s.pdf", time.Now())

	// outputPath := filename
	// file, _ := os.Open(outputPath)

	// fmt.Println("file :", file)

	return req, nil

}
