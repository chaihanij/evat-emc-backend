package dtos

import (
	"fmt"
	"mime/multipart"

	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"

	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UploadScoreAssingment struct {
	AssignmentUUID   string                `uri:"assignment_uuid"`
	Document         *multipart.FileHeader `swaggerignore:"true" form:"document"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
	UpdatedBy        string                `swaggerignore:"true"`
}

type MyRow struct {
	ColumnA string
	ColumnB string
	ColumnC string
}

func (req *UploadScoreAssingment) Parse(c *gin.Context) (*UploadScoreAssingment, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	document, err := c.FormFile("document")
	if err != nil {
		fmt.Println("err:", err)
		return nil, errors.InternalError{Message: err.Error()}
	}
	// fmt.Println("abc :", document)

	req.Document = document
	fileExt := filepath.Ext(document.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(document.Filename), filepath.Ext(document.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	dir := filepath.Join("assignment_teams", "documents")
	os.MkdirAll(filepath.Join(env.DataPath, dir), os.ModePerm)
	dst := filepath.Join(env.DataPath, dir, filename)
	if err := c.SaveUploadedFile(document, dst); err != nil {
		log.WithError(err).Debugln("UpdateMemberImageRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	fmt.Println("fileExt :", fileExt)
	fmt.Println("dtr", dir)
	fmt.Println("dst", dst)
	fmt.Println("document :", document.Size)
	fmt.Println("originalFileName :", originalFileName)

	f, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println("err :", err)
	}

	// c1 := f.GetCellValue("Sheet", "A1")
	// fmt.Println("c1", c1)
	// c2 := f.GetSheetIndex("Sheet")
	// fmt.Println("c2", c2)
	// c3 := f.GetSheetName(1)
	// fmt.Println("c3 :", c3)
	// c4 := f.GetColWidth("Sheet", "A1")
	// fmt.Println("c4", c4)
	// c5 := f.GetRowHeight("Sheet", 1)
	// fmt.Println("C5", c5)

	rows := f.GetRows("Raw data")

	// for _, row := range c6 {
	// 	// 	fmt.Println("i", i)
	// 	for _, colCell := rangerow {

	// 		fmt.Println("key", colCell)
	// 	}
	// }

	for i, row := range rows {
		fmt.Println("number :", i , "row :", row)
		// for _, r := range row {
		// 	fmt.Println("i", i, "r :", r)
		// }

	}

	

	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FilePath = filepath.Join(dir, filename)
	req.FileFullPath = dst
	log.WithField("value", req).Debugln("SendAssignmentDocumentRequestJSON")

	jwtRawData, ok := c.Get(constants.JWTDataKey)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTRestoreFail}
	}

	jwtData, ok := jwtRawData.(entities.JwtData)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTInvalidStructure}
	}

	if jwtData.UID == "" {
		return nil, errors.ParameterError{Message: constants.UserUIDMissing}
	}
	req.UpdatedBy = jwtData.UID
	return req, nil
}
