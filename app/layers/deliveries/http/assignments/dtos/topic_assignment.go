package dtos

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
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
	Topic []string `json:"topic" bson:"topic"`
}

func (u ExportAssignmentTopicResponseJSON) Parse(c *gin.Context, input *entities.ExportAssignmentTopic) *ExportAssignmentTopicResponseJSON {
	topic := &ExportAssignmentTopicResponseJSON{
		Topic: input.Topic,
	}

	// if err := f.SaveAs("simple.xlsx"); err != nil {
	// 	log.Fatal(err)
	// }
	// sheetNameM := "Sheet1"
	dataexcle := [][]string{

		{"", ""},
		{"_id"},
	}
	fm := excelize.NewFile()
	sheetNameM := "Sheet1"
	fm.SetSheetName("Sheet1", sheetNameM)

	for _, value := range topic.Topic {

		// dataexcle = append(dataexcle, value)
		dataexcle = append(dataexcle, []string{value})

		// fmt.Println(index+1, "value :", value)
		// convertRow := fmt.Sprintf("A%d", index+1)
		// fmt.Println("convertRow", convertRow)
		// fm.SetCellValue("Sheet1", convertRow, value)

		// f.GetCellValue("Sheet1", value)

		// f.SetSheetRow("Sheet1", convertRow, value)

		// if err := f.SetSheetRow("Sheet1", startcell, &row); err != nil {
		// 	fmt.Println("err at setsheetrow =>", err)
		// 	return err
		// }

	}
	for i, row := range dataexcle {
		startcell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println("err at startcell =>", err)
		}
		if err := fm.SetSheetRow(sheetNameM, startcell, &row); err != nil {
			fmt.Println("err at setsheetrow =>", err)
		}
	}

	if err := fm.SaveAs("output.xlsx"); err != nil {
		fmt.Println("Error saving file:", err)
	}

	return topic

}
