package teams

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gopkg.in/gomail.v2"
)

func (r repo) SendEmailRegister(data string, activate_code string) error {
	log.Debugln("Send Email")
	log.Debugln("activate_code :", activate_code)

	nameUser := data

	var (
		email = entities.SendEmail{}
		html  = `<!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <link rel="preconnect" href="https://fonts.googleapis.com">
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
        <link href="https://fonts.googleapis.com/css2?family=Prompt&display=swap" rel="stylesheet">
            <title>Document</title>
        </head>
        <style>
        
            body {
                background-size: cover;
                background-position: center;
                position: relative;
                text-align: center;
                display: flex;
                justify-content: center;
                font-family: 'Prompt', sans-serif;
            }
            .head-img img {
                padding-top: 60px;
            }
            .txt-bb {
                padding-bottom: 20px;
            }
            .txt-button {
        
            padding: 15px;
            align-items: center;    
            color: white;
            width: 35%;
            border-radius: 30px;
            background-color: #417630;
            font-size: 17px;
            border: none;
            transform: translateY(-260px);
            }
            .head-main {
                background-color: white;
            }
            .head2-img {
                width: 100%;
                transform: translateY(-270px);
            }
            .txt-link {
                padding-bottom: 20px;
                transform: translateY(-240px);
            }
            .txt-link span{
                font-size: 17px;
                color: #417630;;
            }
            .head2-img img {
                padding-top: 80px;
            }
            .continner {
                background-color: #F8FDF8;
                height: 800px;
                padding: 30px;
                border-radius: 20px;
                margin-top: 200px;
            }
            .txt-decs p {
                font-size: 40px;
                color: #417630;
                transform: translateY(-240px);
            }
            .setting-password {
                cursor: pointer;
            }
            .txt-customer p {
                font-family: 'font-awesome';
                padding-right: 28rem;
                transform: translateY(-220px);
                font-size: 18px;
                font-weight: 500;
            }
            .txt-a p {
                padding-top: 3px;
                padding-bottom: 30px;
                text-align: left;
                transform: translateY(-270px);
            }
            .txt-footer {
                text-align: left;
                font-weight: 100;
                padding-bottom: 9px;
                line-height: 10px;
                transform: translateY(-220px);
            }
            .contact {
                font-weight: 100;
                text-align: left;
                color: #202527;
        
            }
            
        </style>
        <body>
            <div class="head-main">
                <div class="head-second">
                    <div class="head-img">
                        <img src="./logo.png" alt="logo">
                    </div>
                    <div class="desc">
                        <p>โครงการแข่งขันรถจักรยานยนต์ไฟฟ้าดัดแปลงเพื่อธุรกิจแห่งอนาคต</p>
                    </div>
                    <div class="continner">
                        <div class="head2-img">
                            <img src="./hero.svg" alt="logo">
                        </div>
        
                        <div class="txt-customer">
                            <p>สวัสดีคุณ ` + nameUser + `  </p>
                        </div>
                        
                        <div class="txt-decs">
                            <p>การสมัครโครงการได้รับการอนุมัติแล้ว</p>
                        </div>
                        <div class="txt-a">
                            <p>เพื่อดำเนินการให้เสร็จสมบูรณ์ กรุณาตั้งค่ารหัสผ่านโดยกดที่ปุ่มหรือ<br/>กดลิงก์ด้านล่างนี้ ขอบพระคุณครับ</p>
                        </div>
                        <div class="txt-bb">
                            <button class="txt-button">ตั้งค่ารหัสผ่าน</button>
                        </div>
                        <div class="txt-link">
                            <p>หรือกดลิงก์:&nbsp;<span> <a href="https://emc.evat.or.th/dev/change-password?` + activate_code + `" target="_blank" >http://evat.com/123445</a></span></p>
                        </div>
                        <div class="txt-footer">
                            <p>ด้วยความเคารพ</p>
                            <p>สมาคมโครงการแข่งขันรถจักรยานยนต์ไฟฟ้า</p>
                            <p>ดัดแปลงเพื่อธุรกิจแห่งอนาคต</p>
                            <br/>
                            <p>อาคารเคเอกซ์ ชั้น 12เลขที่ 110/1 ถนนกรุงธนบุรี</p>
                            <p>แขวงบางลำภูล่าง เขตคลองสาน กรุงเทพมหานคร 10600</p>
                            <p>อีเมล: contact@evat.or.th. โทร: 086-390-3339</p>
                        </div>
                        <div class="contact">
                            
                        </div>
                    </div>
        
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

	// email.Sender = "pr@evat.or.th"
	// email.Password = "evatOMS100%"
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
