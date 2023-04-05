package teams

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gopkg.in/gomail.v2"
)

func (r repo) SendEmailRegister(data string) error {
	log.Debugln("Send Email")

	nameUser := data

	var (
		email = entities.SendEmail{}
		html  = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<style>

    body {
        background-color: powderblue;
        text-align: center;
        display: flex;
        justify-content: center;
    }
    .continner {
        width: 80%;
    }
    .txt-decs {
        /* Layout Properties */
        /* UI Properties */

        text-align: left;
        color: #417630;
    }
    .setting-password {
        cursor: pointer;
    }
    .txt-customer {
        text-align: left;
    }
    .txt-a {
        text-align: left;
    }
    .txt-footer {
        text-align: left;
        color: #202527;
    }
    .contact {
        text-align: left;
        color: #202527;

    }
    
</style>
<body>

    <div class="continner">

        <div class="head-img">
            <img src="./logo.png" alt="logo">
        </div>
        <div class="desc">
            <p>โครงการแข่งขันรถจักรยานยนต์ไฟฟ้าดัดแปลงเพื่อธุรกิจแห่งอนาคต</p>
        </div>
        <div>
            <img src="./desc.png" alt="">
        </div>

        <div class="txt-customer">
            <p>สวัสดีคุณ ` + nameUser + `  </p>
        </div>
        
        <div class="txt-decs">
            <p>การสมัครโครงการได้รับการอนุมัติแล้ว</p>
        </div>
        <div class="txt-a">
            <p>เพื่อดำเนินการให้เสร็จสมบูรณ์ กรุณาตั้งค่ารหัสผ่านโดยกดที่ปุ่มหรือ</p>
            <p>กดลิงก์ด้านล่างนี้ ขอบพระคุณครับ</p>
        </div>
        <div class="setting-password" >
            <img src="https://cdn-icons-png.flaticon.com/512/59/59408.png" alt="verify">
        </div>
        <div class="txt-footer">
            <p>ด้วยความเคารพ</p>
            <p>สมาคมโครงการแข่งขันรถจักรยานยนต์ไฟฟ้า</p>
            <p>ดัดแปลงเพื่อธุรกิจแห่งอนาคต</p>
        </div>
        <div class="contact">
            <p>อาคารเคเอกซ์ ชั้น 12เลขที่ 110/1 ถนนกรุงธนบุรี</p>
            <p>แขวงบางลำภูล่าง เขตคลองสาน กรุงเทพมหานคร 10600</p>
            <p>อีเมล: contact@evat.or.th. โทร: 086-390-3339</p>
        </div>

    </div>

</body>
</html>`
	)

	m := gomail.NewMessage()

	subject := "ยืนยันการสมัครเข้าแข่งขัน"
	content := html
	receiver := data
	cc := ""

	email.Sender = "sanch_ai@hotmail.com"
	email.Password = "0877380568"

	mailto := strings.Split(receiver, ",")
	mailcc := strings.Split(cc, ",")

	receiverto := make([]string, len(mailto))
	for i, recipient := range mailto {
		receiverto[i] = m.FormatAddress(recipient, "")
	}

	addresses := make([]string, len(mailcc))
	for i, recipient := range mailcc {
		addresses[i] = m.FormatAddress(recipient, "")
	}

	m.SetHeader("To", receiverto...)
	m.SetHeader("From", email.Sender)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	if cc != "" {
		m.SetHeader("Cc", addresses...)
	}
	d := gomail.NewDialer("smtp.office365.com", 587, email.Sender, email.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error", err)
		return err

	}

	return nil

}
