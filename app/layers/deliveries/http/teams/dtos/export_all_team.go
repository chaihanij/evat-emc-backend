package dtos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/utils"
)

type ExportTeamRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
	Name     *string `form:"name"`
	TeamType *string `form:"teamtype"`
}

func (req *ExportTeamRequestJSON) Parse(c *gin.Context) (*ExportTeamRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}
func (req *ExportTeamRequestJSON) ToEntity() *entities.TeamFilter {
	return &entities.TeamFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		TeamType: req.TeamType,
	}
}

type ExportAllTeamResponseJSON []TeamResponse

func (m *ExportAllTeamResponseJSON) Parse(c *gin.Context, data []entities.Team, members []entities.Member) *ExportAllTeamResponseJSON {
	var teams ExportAllTeamResponseJSON = ExportAllTeamResponseJSON{}
	for _, value := range data {
		team := &TeamResponse{
			UUID:      value.UUID,
			Code:      value.Code,
			Name:      value.Name,
			TeamType:  value.TeamType,
			Academy:   value.Academy,
			Detail:    value.Detail,
			Year:      value.Year,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			CreatedBy: value.CreatedBy,
			UpdatedBy: value.UpdatedBy,
			IsPaid:    value.IsPaid,
		}
		teams = append(teams, *team)
	}

	fm := excelize.NewFile()

	sheetNameM := "team"
	index, _ := fm.NewSheet("team")
	fm.SetCellValue(sheetNameM, "A1", "ลำดับ")
	fm.SetColWidth(sheetNameM, "A", "A", 20)
	fm.SetCellValue(sheetNameM, "B1", "รหัสทีม")
	fm.SetColWidth(sheetNameM, "B", "B", 20)
	fm.SetCellValue(sheetNameM, "C1", "ประเภททีม")
	fm.SetColWidth(sheetNameM, "C", "C", 20)
	fm.SetCellValue(sheetNameM, "D1", "ชื่อทีม")
	fm.SetColWidth(sheetNameM, "D", "D", 20)
	fm.SetCellValue(sheetNameM, "E1", "สถาบันการศึกษา")
	fm.SetColWidth(sheetNameM, "E", "E", 20)
	fm.SetCellValue(sheetNameM, "F1", "วันอนุมัติ")
	fm.SetColWidth(sheetNameM, "F", "F", 20)

	id := 1
	for _, value := range teams {

		var PaidDateTime interface{}
		PaidDateTime = ""

		theTime := time.Date(0001, 01, 01, 00, 00, 00, 100, time.UTC)

		if theTime.Unix() != value.PaidDateTime.Unix() {
			PaidDateTime = value.PaidDateTime
		}

		for i := 1; i <= 1; i++ {
			id += 1

			convertRow := fmt.Sprintf("A%d", id)
			fm.SetCellValue(sheetNameM, convertRow, id-1)

			//Code
			convertRow = fmt.Sprintf("B%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.Code)

			//teamtype
			convertRow = fmt.Sprintf("C%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.TeamType)

			//teamname
			convertRow = fmt.Sprintf("D%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.Name)

			//academy
			convertRow = fmt.Sprintf("E%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.Academy)

			convertRow = fmt.Sprintf("F%d", id)
			fm.SetCellValue(sheetNameM, convertRow, PaidDateTime)
		}

	}

	fm.DeleteSheet("Sheet1")
	fm.SetActiveSheet(index)

	fileName := fmt.Sprintf("team-%v-%v-%v.xlsx", time.Now().Day(), time.Now().Month(), time.Now().Year())
	fileExt := filepath.Ext(fileName)
	originalFileName := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	filenames := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v-%v-%v", time.Now().Day(), time.Now().Month(), time.Now().Year()) + fileExt

	dst := filepath.Join(env.DataPath, "teams", "export", filenames)

	if err := fm.SaveAs(dst); err != nil {
	}
	f, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println("Error opening file:", err)
		// return
	}
	sheetMember := "member"
	index2, _ := f.NewSheet("member")
	f.SetCellValue("member", "A1", "Hello from Sheet2!")

	f.SetCellValue(sheetMember, "A1", "ลำดับ")
	f.SetColWidth(sheetMember, "A", "A", 20)
	f.SetCellValue(sheetMember, "B1", "ชื่อ - นามสกุล")
	f.SetColWidth(sheetMember, "B", "B", 20)
	f.SetCellValue(sheetMember, "C1", "วันที่ลงทะเบียน")
	f.SetColWidth(sheetMember, "C", "C", 40)
	f.SetCellValue(sheetMember, "D1", "เลขบัตรประชาชน")
	f.SetColWidth(sheetMember, "D", "D", 25)
	f.SetCellValue(sheetMember, "E1", "ข้อมูล")
	f.SetColWidth(sheetMember, "E", "E", 20)
	f.SetCellValue(sheetMember, "F1", "รูปตรงบัตรประชาชน")
	f.SetColWidth(sheetMember, "F", "F", 20)
	f.SetCellValue(sheetMember, "G1", "ลงทะเบียน")
	f.SetColWidth(sheetMember, "G", "G", 20)

	idx := 1

	for _, value := range members {
		// idx += 1
		is_national := "ไม่ผ่าน"

		if *value.Check_national == true {
			is_national = "ผ่าน"
		}

		is_data := "ข้อมูลไม่ถูกต้อง"

		if *value.Is_check_data == true {
			is_data = "ข้อมูลถูกต้อง"
		}

		is_image := "รูปไม่ตรงบัตรประชาชน"

		if *value.Is_Check_image == true {
			is_image = "รูปตรงบัตรประชาชน"
		}

		is_checkin := "ยังไม่ได้ลงทะเบียน"
		if *value.Is_checkin == true {
			is_checkin = "ลงทะเบียนแล้ว"
		}

		var dateCheckIn interface{}
		dateCheckIn = ""

		// if *value.Checkin_date == ""{

		// }

		theTime := time.Date(0001, 01, 01, 00, 00, 00, 100, time.UTC)

		if theTime.Unix() != value.Checkin_date.Unix() {
			dateCheckIn = value.Checkin_date
		}

		for i := 1; i <= 1; i++ {
			idx += 1

			row := fmt.Sprintf("A%d", idx)
			f.SetCellValue(sheetMember, row, idx-1)

			row = fmt.Sprintf("B%d", idx)
			name := fmt.Sprintf("%s %s %s ", *value.Prefix, value.FirstName, value.LastName)
			f.SetCellValue(sheetMember, row, name)

			row = fmt.Sprintf("C%d", idx)
			f.SetCellValue(sheetMember, row, dateCheckIn)

			row = fmt.Sprintf("D%d", idx)
			f.SetCellValue(sheetMember, row, is_national)

			row = fmt.Sprintf("E%d", idx)
			f.SetCellValue(sheetMember, row, is_data)

			row = fmt.Sprintf("F%d", idx)
			f.SetCellValue(sheetMember, row, is_image)

			row = fmt.Sprintf("G%d", idx)
			f.SetCellValue(sheetMember, row, is_checkin)

		}
	}

	f.SetActiveSheet(index2)

	// Save the changes to the file
	if err := f.SaveAs(dst); err != nil {
		fmt.Println("Error saving file:", err)
		// return
	}
	fileBytes, err := ioutil.ReadFile(dst)
	if err != nil {
		utils.JSONErrorResponse(c, err)
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", `attachment; filename=`+filenames)
	c.Header("Content-Length", fmt.Sprintf("%d", len(fileBytes)))
	c.Writer.Write(fileBytes)

	return &teams
}
