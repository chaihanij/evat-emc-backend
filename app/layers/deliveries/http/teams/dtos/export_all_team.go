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

func (m *ExportAllTeamResponseJSON) Parse(c *gin.Context, data []entities.Team) *ExportAllTeamResponseJSON {
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

	sheetNameM := "Sheet1"
	fm.SetCellValue(sheetNameM, "A1", "ลำดับ")
	fm.SetColWidth(sheetNameM, "A", "A", 20)
	// fm.SetCellValue(sheetNameM, "A1", "รหัสทีม")
	// fm.SetColWidth(sheetNameM, "A", "A", 20)
	fm.SetCellValue(sheetNameM, "B1", "รหัสทีม")
	fm.SetColWidth(sheetNameM, "B", "B", 20)
	fm.SetCellValue(sheetNameM, "C1", "ประเภททีม")
	fm.SetColWidth(sheetNameM, "C", "C", 20)
	fm.SetCellValue(sheetNameM, "D1", "ชื่อทีม")
	fm.SetColWidth(sheetNameM, "D", "D", 20)
	fm.SetCellValue(sheetNameM, "E1", "สถานันการศึกษา")
	fm.SetColWidth(sheetNameM, "E", "E", 20)
	//สถานันการศึกษา

	// fmt.Println("team :", teams)
	id := 1
	for index, value := range teams {
		fmt.Println("index :", index, "value :", value)

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
		}

	}

	fileName := fmt.Sprintf("team-%v-%v-%v.xlsx", time.Now().Day(), time.Now().Month(), time.Now().Year())
	fileExt := filepath.Ext(fileName)
	originalFileName := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	filenames := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v-%v-%v", time.Now().Day(), time.Now().Month(), time.Now().Year()) + fileExt

	dst := filepath.Join(env.DataPath, "teams", "export", filenames)

	if err := fm.SaveAs(dst); err != nil {
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
