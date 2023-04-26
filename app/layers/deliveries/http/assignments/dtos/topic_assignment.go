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
	fm.SetColWidth(sheetNameM, "A", "A", 20)
	fm.SetCellValue(sheetNameM, "B1", "ประเภทผู้เข้าร่วมแข่งขัน")
	fm.SetColWidth(sheetNameM, "B", "B", 20)
	fm.SetCellValue(sheetNameM, "C1", "รหัสทีม")
	fm.SetColWidth(sheetNameM, "C", "C", 20)
	fm.SetCellValue(sheetNameM, "D1", "ชื่อทีม")
	fm.SetColWidth(sheetNameM, "D", "D", 20)
	// fm.SetCellValue(sheetNameM, "E1", "รวมคะแนน")
	// fm.SetColWidth(sheetNameM, "E", "E", 20)
	// fm.SetCellValue(sheetNameM, "F1", "ลำดับ")
	// fm.SetColWidth(sheetNameM, "F", "F", 20)

	// fill := excelize.Fill{
	// 	Type:    "pattern",
	// 	Color:   []string{"#fef2cb"},
	// 	Pattern: 1,
	// }

	// borderleft := excelize.Border{

	// 	Type:  "left",
	// 	Style: 1,
	// 	Color: "#000000",
	// }
	// bordertop := excelize.Border{

	// 	Type:  "top",
	// 	Style: 1,
	// 	Color: "#000000",
	// }
	// borderbottom := excelize.Border{

	// 	Type:  "bottom",
	// 	Style: 1,
	// 	Color: "#000000",
	// }
	// borderright := excelize.Border{

	// 	Type:  "right",
	// 	Style: 1,
	// 	Color: "#000000",
	// }
	// style := excelize.Style{
	// 	Border: []excelize.Border{
	// 		borderleft,
	// 		bordertop,
	// 		borderbottom,
	// 		borderright,
	// 	},
	// 	Fill: fill,
	// 	// Border: []excelize.Border{
	// 	// 	border,
	// 	// },
	// }
	// styleID, _ := fm.NewStyle(&style)
	// fm.SetCellStyle("Sheet1", "A1", "F1", styleID)

	asc := len(topic.Topic)
	as := ""
	ch := 69
	chasc := ch + asc
	idx := 0
	for ch = 69; ch < chasc; ch++ {
		as = fmt.Sprintf("%c", ch)

		// index := 0
		convertRow := fmt.Sprintf("%s%d", as, 1)
		fm.SetCellValue(sheetNameM, convertRow, topic.Topic[idx])
		fm.SetColWidth(sheetNameM, as, as, 40)

		fill := excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#fef2cb"},
			Pattern: 1,
		}

		borderleft := excelize.Border{

			Type:  "left",
			Style: 1,
			Color: "#000000",
		}
		bordertop := excelize.Border{

			Type:  "top",
			Style: 1,
			Color: "#000000",
		}
		borderbottom := excelize.Border{

			Type:  "bottom",
			Style: 1,
			Color: "#000000",
		}
		borderright := excelize.Border{

			Type:  "right",
			Style: 1,
			Color: "#000000",
		}
		style := excelize.Style{
			Border: []excelize.Border{
				borderleft,
				bordertop,
				borderbottom,
				borderright,
			},
			Fill: fill,
			// Border: []excelize.Border{
			// 	border,
			// },
		}
		styleID, _ := fm.NewStyle(&style)
		fm.SetCellStyle("Sheet1", "A1", convertRow, styleID)

		// fm.SetColWidth()

		// fmt.Println("ch :", ch)

		// asend := fmt.Sprintf("%c", ch+1)
		// convertEnd := fmt.Sprintf("%s%d", asend, 1)
		// fm.SetCellValue(sheetNameM, convertEnd, "รวม")

		idx += 1
	}
	// fmt.Println("ch :", ch)
	// character := 71
	id := 1

	student := []string{}
	population := []string{}

	for _, v := range topic.ExportTeamTopic {

		if v.Team_type == "STUDENT" {
			student = append(student, v.Name)
		} else {
			population = append(population, v.Name)

		}
	}

	for index, value := range topic.ExportTeamTopic {
		if value.Team_type == "STUDENT" {
			value.Team_type = "ประเภทสถาบันการศึกษา"
		} else {
			value.Team_type = "ประเภทประชาชนทั่วไป"
		}

		for i := 1; i <= 1; i++ {
			id += 1

			// convertRow := fmt.Sprintf("A%d", id)
			// fm.SetCellValue(sheetNameM, convertRow, id-1)
			// convertRow = fmt.Sprintf("B%d", id)
			// fm.SetCellValue(sheetNameM, convertRow, value.Team_type)
			// convertRow = fmt.Sprintf("C%d", id)
			// fm.SetCellValue(sheetNameM, convertRow, topic.ExportTeamTopic[index].Code)
			// convertRow = fmt.Sprintf("D%d", id)
			// fm.SetCellValue(sheetNameM, convertRow, value.Name)

			// characterStop := fmt.Sprintf("%c", ch-1)

			// characterStart := fmt.Sprintf("%c", character)

			// characterCount := fmt.Sprintf("SUM(%s%d:%s%d)", characterStart, id, characterStop, id)

			// convertRow = fmt.Sprintf("E%d", id)
			// fm.SetCellFormula(sheetNameM, convertRow, characterCount)
			// convertRow = fmt.Sprintf("F%d", id)

			// formularank := fmt.Sprintf("RANK(E%d,$E$2:$E%d)", id, len(topic.ExportTeamTopic)+1)
			// fm.SetCellFormula(sheetNameM, convertRow, formularank)

			if value.Team_type == "ประเภทสถาบันการศึกษา" {

				convertRow := fmt.Sprintf("A%d", id)
				fm.SetCellValue(sheetNameM, convertRow, id-1)
				convertRow = fmt.Sprintf("B%d", id)
				fm.SetCellValue(sheetNameM, convertRow, value.Team_type)
				convertRow = fmt.Sprintf("C%d", id)
				fm.SetCellValue(sheetNameM, convertRow, topic.ExportTeamTopic[index].Code)
				convertRow = fmt.Sprintf("D%d", id)
				fm.SetCellValue(sheetNameM, convertRow, value.Name)

				characterStop := fmt.Sprintf("%c", ch-1)

				// characterStart := fmt.Sprintf("%c", character)

				// characterCount := fmt.Sprintf("SUM(%s%d:%s%d)", characterStart, id, characterStop, id)

				// convertRow = fmt.Sprintf("E%d", id)
				// fm.SetCellFormula(sheetNameM, convertRow, characterCount)
				// convertRow = fmt.Sprintf("F%d", id)

				// formularank := fmt.Sprintf("RANK(E%d,$E$2:$E%d)", id, len(student)+1)
				// fm.SetCellFormula(sheetNameM, convertRow, formularank)

				fill := excelize.Fill{
					Type:    "pattern",
					Color:   []string{"#e2edd9"},
					Pattern: 1,
				}

				borderleft := excelize.Border{

					Type:  "left",
					Style: 1,
					Color: "#000000",
				}
				bordertop := excelize.Border{

					Type:  "top",
					Style: 1,
					Color: "#000000",
				}
				borderbottom := excelize.Border{

					Type:  "bottom",
					Style: 1,
					Color: "#000000",
				}
				borderright := excelize.Border{

					Type:  "right",
					Style: 1,
					Color: "#000000",
				}

				style := excelize.Style{
					Border: []excelize.Border{
						borderleft,
						bordertop,
						borderbottom,
						borderright,
					},
					Fill: fill,
					// Border: []excelize.Border{
					// 	border,
					// },
				}
				styleID, _ := fm.NewStyle(&style)
				highlightend := fmt.Sprintf("%s%d", characterStop, len(student)+1)
				fm.SetCellStyle("Sheet1", "A2", highlightend, styleID)

			}
			if value.Team_type == "ประเภทประชาชนทั่วไป" {
				convertRow := fmt.Sprintf("A%d", id)
				fm.SetCellValue(sheetNameM, convertRow, id-1)
				convertRow = fmt.Sprintf("B%d", id)
				fm.SetCellValue(sheetNameM, convertRow, value.Team_type)
				convertRow = fmt.Sprintf("C%d", id)
				fm.SetCellValue(sheetNameM, convertRow, topic.ExportTeamTopic[index].Code)
				convertRow = fmt.Sprintf("D%d", id)
				fm.SetCellValue(sheetNameM, convertRow, value.Name)

				characterStop := fmt.Sprintf("%c", ch-1)

				// characterStart := fmt.Sprintf("%c", character)

				// characterCount := fmt.Sprintf("SUM(%s%d:%s%d)", characterStart, id, characterStop, id)

				// convertRow = fmt.Sprintf("E%d", id)
				// fm.SetCellFormula(sheetNameM, convertRow, characterCount)
				// convertRow = fmt.Sprintf("F%d", id)

				//character
				// formularank1 := fmt.Sprintf("RANK(E%d,$E$%d:$E%d)", id, len(student), len(student)+1)

				// formularank := fmt.Sprintf("RANK(E%d,$E$%d:$E%d)", id, len(student)+2, len(population)+len(student)+1)
				// fm.SetCellFormula(sheetNameM, convertRow, formularank)

				fill := excelize.Fill{
					Type:    "pattern",
					Color:   []string{"#fbe3d4"},
					Pattern: 1,
				}
				borderleft := excelize.Border{

					Type:  "left",
					Style: 1,
					Color: "#000000",
				}
				bordertop := excelize.Border{

					Type:  "top",
					Style: 1,
					Color: "#000000",
				}
				borderbottom := excelize.Border{

					Type:  "bottom",
					Style: 1,
					Color: "#000000",
				}
				borderright := excelize.Border{

					Type:  "right",
					Style: 1,
					Color: "#000000",
				}

				style := excelize.Style{
					Border: []excelize.Border{
						borderleft,
						bordertop,
						borderbottom,
						borderright,
					},
					Fill: fill,
					// Border: []excelize.Border{
					// 	border,
					// },
				}
				styleID, _ := fm.NewStyle(&style)

				highlightstart := fmt.Sprintf("A%d", len(student)+2)
				highlightend := fmt.Sprintf("%s%d", characterStop, len(population)+len(student)+1)
				// fmt.Println("___ :", len(population)+len(student))
				// fmt.Println("highlightend :", highlightend)
				fm.SetCellStyle("Sheet1", highlightstart, highlightend, styleID)

			}

		}
	}

	// startRow := 1
	// // endRow := len(rows)
	// err = fm.RemoveRow(sheetNameM, startRow)
	// // err = fm.RemoveRows(sheetNameM, startRow, endRow)
	// if err != nil {
	// 	fmt.Println(err)

	// }

	fileName := fmt.Sprintf("%s-%v-%v-%v.xlsx", topic.Title, time.Now().Day(), time.Now().Month(), time.Now().Year())

	fileExt := filepath.Ext(fileName)

	originalFileName := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	filenames := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v-%v-%v", time.Now().Day(), time.Now().Month(), time.Now().Year()) + fileExt

	dst := filepath.Join(env.DataPath, "assignments", "template", filenames)

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

	return topic

}
