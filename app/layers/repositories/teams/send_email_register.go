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
		  <title>Verify Email</title>
		</head>
		
		<body>
		  <div class="cov-form" style='width: 800px; margin: 0 auto;'>
			<div class="header-email">
			  <div class="img-logo"
				style='background-image: url("https://evspms.pttor.com/ptt_image/other/Rectangle_10@2x.png"); background-size: 100% 100%; background-repeat: no-repeat;'>
				<img src="https://evspms.pttor.com/ptt_image/logo/evlogo1.jpg" alt="" style='width: 150px;'>
			  </div>
			</div>
			<div class="body-email" style='min-height: 750px;'>
			  <div class="cov-header" style='text-align: right'>
				<div class="txt-detail" style='margin-top: 30px'>
				  <p class="txt" style="font-size: 16px; text-align: left; color: rgba(144,144,144,1);">
					สวัสดีคุณ  ` + nameUser + ` 
				  </p>
				</div>
				  <div class="txt-detail" style='margin-top: 30px'>
				  <p class="txt" style="font-size: 16px; text-align: left; color: rgba(144,144,144,1);">
				  ขอต้อนรับสู่บริการชาร์จไฟฟ้าสำหรับรถยนต์ไฟฟ้าด้วยแอปพลิเคชัน อีวี สเตชั่น พลัส
				  <br/><br/>
		
				  คุณได้ลงทะเบียนสมัครใช้งานผ่านช่องทางออนไลน์เรียบร้อยแล้ว ขั้นตอนต่อไป กรุณากดปุ่มด้านล่างเพื่อตั้งรหัสผ่านใหม่ 
				  พร้อมยืนยันการสมัครใช้บริการของคุณอีกครั้ง ทั้งนี้ บริษัทจะใช้อีเมลนี้ในการส่งเอกสารใบกำกับภาษี / ใบเสร็จรับเงินในรูปแบบอิเล็กทรอนิกส์
				  สำหรับค่าบริการ อีวี สเตชั่น พลัส ที่คุณได้รับในโอกาสถัดไป
				  </p>
				</div>
				<div class="txt-detail" style='margin-top: 30px'>
				  <p class="txt" style="font-size: 16px; text-align: left; color: rgba(144,144,144,1);">
				  หากคุณไม่ได้เป็นผู้สมัครบริการนี้ ไม่ต้องทำการใดๆ หรือตอบกลับอีเมลเพื่อแจ้งให้เราทราบ
				  </p>
				</div>
			  </div>
			  <div class="txt-detail" style='margin-top: 30px'>
				<p class="txt" style="font-size: 16px; text-align: left; color: rgba(144,144,144,1);">
				  ด้วยความเคารพ <br />
				  ศูนย์ปฏิบัติการ อีวี สเตชั่น พลัส <br/>
				  บริษัท​ ปตท.​ น้ำมันและการค้าปลีก​ จำกัด​ (มหาชน) <br/>
				  โทร 02-061-9519
				</p>
			  </div>
              <hr />
               <hr />
			</div>
		  </div>
		</body>
		
		</html>`
	)

	m := gomail.NewMessage()

	subject := "topic"
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
