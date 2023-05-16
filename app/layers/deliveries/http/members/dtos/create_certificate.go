package dtos

import (
	"bytes"
	"html/template"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
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

const templateString = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<style>
    @font-face {
    font-family: testfont;
    src: url("/Cloud-Light\ 300.otf") format("opentype");
    }
    *
    {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    body {
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
        background: #fff;
    }
    .card{
        position: relative;
        width: 1000px;
        height: 1000px;
         
        border-radius: 20px;
        box-shadow: 0 15px 35px rgba(0,0,0,0.25);
    }
    .card .poster {
        position: relative;
        overflow: hidden;
    }
    .card .poster::before
    {
        content: '';
        position: absolute;
        bottom: -180px;
        width: 100%;
        height: 100%;
        transition: 0.5s;
        z-index: 1;
    }
    .card .poster img {
        width: 100%;
    }
    .card .details {
        position: absolute;
        bottom: 0;
        left: 0;
        padding: 100px 800px;
        width: 100%;
        z-index: 2;
        padding-bottom: 350px;
        transition: 0.5s;
    }
    .card .deta {
        position: absolute;
        bottom: 0;
        left: 0;
        padding: 320px 100px;
        width: 100%;
        z-index: 2;
        transition: 0.5s;
    }
    .card .deta img {
        width: 900px;
        margin-left: -35px;
        transform: translateY(-220px);
        padding-bottom: 0.1px;
    }
    .deta h1 {
        transform: translateY(-164px);
        text-align: center;
        font-size: 28px;
    }
    .log {
         
        scale: 0.6;

    }
    .deta img {
        scale: 0.4;
    }
    .deta h2 {
        text-align: center;
        transform: translateY(-163px);
        font-size: 22px;
        padding-bottom: 21px;
         
    }
    .deta h1 {
        text-align: center;
        font-size: 28px;
    }
     
    .details img {
        padding:  10px 39px;
         
        scale: 1.7;
    }
    .dela h1 {
        text-align: center;
        font-size: 21px;
        font-weight: 900;
        line-height: 28px;
    }
    .dela {
        padding-bottom: 40px;
    }
    .veta h1 {
        font-size: 18px;
    }
    .og {
        scale: 0.6;
         
    }
    .oz {
        scale: 0.3;
    }
    .veta .dex {
        position: absolute;
        bottom: 0;
         
         
        padding-left: 10px;
        width: 100%;
         
        z-index: 2;
        padding-bottom: 135px;
        transition: 0.5s;
    }
    .zog {
         
    }
    .veta .gex {
        position: absolute;
        bottom: 0;
         
         
        padding-left: 80px;
        width: 100%;
         
        z-index: 2;
        padding-bottom: 5px;
        transition: 0.5s;
    }
    .beta .gex {
        position: absolute;
        bottom: 0;
         
         
        padding-left: 80px;
        width: 100%;
         
        z-index: 2;
        padding-bottom: 190px;
        transition: 0.5s;
    }
    .deta-1 h1 {
        font-family: myFont;
        font-size: 20px;
    }
    .dex h3 {
        font-size: 16px;
        padding: 40px 10px;
        text-align: center;
        margin-left: -440px;
        transform: translateY(-165px);
         
    }
    .dex h2 {
        font-size: 16px;
        margin-left: 90px;
        padding: 1px 20px;
        text-align: center;
        transform: translateY(-250px);
    }
     
     
</style>
<body>
    <div class="card">
        <div class="poster">
            <img src="./background evat certificate.png">
        </div>
        <div class="details">
            <img src="./iz.png" class="logo">
            
        </div>
        <div class="deta">
            <img src="https://news.evat.online/content/images/2021/08/333-1.png" class="log">
            
            <h1>สมาคมยานยนต์ไฟฟ้าไทยร่วมกับ</br>การไฟฟ้าฝ่ายผลิตแห่งประเทศไทย</br><h2>ขอมอบเกียรติบัตรนี้เพื่อแสดงว่า</h2></h1>
            <h2> {{.Prefix}} {{.Firstname}} {{.Lastname}} </h2>
            <div class="dela">
                <h1>ได้เข้าร่วมกิจกรรม</br>โครงการแข่งขันรถจักรยานยนต์ไฟฟ้าดัดแปลงเพื่อธุรกิจแห่งอนาคต</br>(EVAT x EGAT Electric Motorcycle Conversion Contest for Business Opportunity)</br>วันที่ 19-21 พฤษภาคม พ.ศ. 2566</h1>
            </div>
            <div class="deta-1">
                <h1>ให้ไว้ ณ วันที่ 21 พฤษภาคม พ.ศ. 2566</h1>
                    
            </div>
            <div class="veta">
                
                        <div class="dex">
                            <img 
                                src="./signature.jpg" 
                                class="og"
                            >
                            <h3>นายกฤษฏา อุตตโมทย์</br>นายกสมาคมยานยนต์ไฟฟ้าไทย</h3>
                            <h2>นายบุญญนิตย์ วงศ์รักมิตร</br>ผู้ว่าการการไฟฟ้าฝ่ายผลิดแห่งประเทศไทย</h2>
                        </div>
                    
            </div>
        </div> 
    </div>
</body>
</html>
`

func (res *ResponseCertificate) Parse(c *gin.Context, input *entities.Member) ([]byte, error) {
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
