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

type FindTopicRequestJSON struct {
	UUID string `uri:"assignment_uuid"`
}

func (req *FindTopicRequestJSON) Parse(c *gin.Context) (*FindTopicRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindTopicRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		UUID: &req.UUID,
	}
}

type ExportAssignmentTopicResponseJSON struct {
	Topic           []string `json:"topic" bson:"topic"`
	Title           string   `json:"title" bson:"title"`
	Path            string   `json:"path"`
	ExportTeamTopic []ExportTeamTopic
}

type ExportTeamTopic struct {
	Code      string `json:"code" bson:"code" `
	Name      string `json:"name" bson:"name"`
	Team_type string `json:"team_type" bson:"team_type" `
}

func (u ExportAssignmentTopicResponseJSON) Parse(c *gin.Context, input *entities.ExportAssignmentTopic) *ExportAssignmentTopicResponseJSON {

	var ExportTeamTopics []ExportTeamTopic

	for _, value := range input.ExportTeamTopic {

		ExportTeamTopic := ExportTeamTopic{
			Code:      value.Code,
			Name:      value.Name,
			Team_type: value.Team_type,
		}

		ExportTeamTopics = append(ExportTeamTopics, ExportTeamTopic)

	}

	topic := &ExportAssignmentTopicResponseJSON{
		Path:            input.Path,
		Title:           input.Title,
		Topic:           input.Topic,
		ExportTeamTopic: ExportTeamTopics,
	}

	// if err := f.SaveAs("simple.xlsx"); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("title", topic.Title)
	fm := excelize.NewFile()
	sheetNameM := "Sheet1"

	fm.SetCellValue(sheetNameM, "A1", "ทีมที่")
	fm.SetCellValue(sheetNameM, "B1", "ประเภทผู้เข้าร่วมแข่งขัน")
	fm.SetCellValue(sheetNameM, "C1", "รหัสทีม")
	fm.SetCellValue(sheetNameM, "D1", "ชื่อทีม")
	fm.SetCellValue(sheetNameM, "E1", "รวมคะแนน")
	fm.SetCellValue(sheetNameM, "F1", "ลำดับ")

	asc := len(topic.Topic)
	as := ""
	ch := 71
	chasc := ch + asc
	idx := 0
	for ch = 71; ch < chasc; ch++ {
		as = fmt.Sprintf("%c", ch)

		// index := 0
		convertRow := fmt.Sprintf("%s%d", as, 1)
		fm.SetCellValue(sheetNameM, convertRow, topic.Topic[idx])
		// fmt.Println("ch :", ch)

		// asend := fmt.Sprintf("%c", ch+1)
		// convertEnd := fmt.Sprintf("%s%d", asend, 1)
		// fm.SetCellValue(sheetNameM, convertEnd, "รวม")

		idx += 1
	}
	// fmt.Println("ch :", ch)
	character := 71
	id := 1
	for index, value := range topic.ExportTeamTopic {

		if value.Team_type == "STUDENT" {
			value.Team_type = "ประเภทสถาบันการศึกษา"
		} else {
			value.Team_type = "ประเภทประชาชนทั่วไป"
		}

		// id += 1
		for i := 1; i <= 1; i++ {
			id += 1

			// fmt.Println("i :", i, "id", id, "index", index, "value.Code", topic.ExportTeamTopic[index].Code)
			// fmt.Println("index :", index-1)

			convertRow := fmt.Sprintf("A%d", id)
			fm.SetCellValue(sheetNameM, convertRow, id-1)
			convertRow = fmt.Sprintf("B%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.Team_type)
			convertRow = fmt.Sprintf("C%d", id)
			fm.SetCellValue(sheetNameM, convertRow, topic.ExportTeamTopic[index].Code)
			convertRow = fmt.Sprintf("D%d", id)
			fm.SetCellValue(sheetNameM, convertRow, value.Name)

			characterStop := fmt.Sprintf("%c", ch-1)
			// fmt.Println("characterStop", characterStop, id)

			characterStart := fmt.Sprintf("%c", character)
			// fmt.Println("characterStart", characterStart, id)

			characterCount := fmt.Sprintf("SUM(%s%d:%s%d)", characterStart, id, characterStop, id)
			// fmt.Println("characterCount", characterCount)
			//err := f.SetCellFormula("Sheet1", "A3", "=SUM(A1,B1)")

			convertRow = fmt.Sprintf("E%d", id)
			fm.SetCellFormula(sheetNameM, convertRow, characterCount)
			// fm.SetCellValue(sheetNameM, convertRow, "12")
			//=RANK(E2,$E$2:$E$3)
			convertRow = fmt.Sprintf("F%d", id)

			formularank := fmt.Sprintf("RANK(E%d,$E$2:$E%d)", id, len(topic.ExportTeamTopic)+1)
			// fmt.Println("formularank :", formularank)
			fm.SetCellFormula(sheetNameM, convertRow, formularank)

		}
	}

	fileName := fmt.Sprintf("%s.xlsx", topic.Title)

	fileExt := filepath.Ext(fileName)

	originalFileName := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	filenames := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt

	dst := filepath.Join(env.DataPath, "assignments", "template", filenames)
	// fmt.Println("dst", dst)

	if err := fm.SaveAs(dst); err != nil {
	}

	// aa := fmt.Sprintf("%s/v1/files/%s", env.BaseUrl, filenames)

	// topic.Path = fmt.Sprintf("%s/v1/files/%s", env.BaseUrl, filenames)

	// fmt.Println("fileExt", fileExt)
	// fmt.Println("originalFileName", originalFileName)
	// fmt.Println("filenames", filenames)
	fileBytes, err := ioutil.ReadFile(dst)
	if err != nil {
		utils.JSONErrorResponse(c, err)
	}
	// topic.Path = fmt.Sprintf("%d", len(fileBytes))
	c.Writer.WriteHeader(http.StatusOK)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", `attachment; filename=`+filenames)
	c.Header("Content-Length", fmt.Sprintf("%d", len(fileBytes)))
	c.Writer.Write(fileBytes)

	return topic

}
