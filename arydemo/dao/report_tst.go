package dao

import (
	"context"
	"got/arydemo/model"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllReportTest is a function to get a slice of record(s) from report_test table in the go_by_example database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllReportTest(ctx context.Context, page, pagesize int64, order string) (results []*model.ReportTest, totalRows int, err error) {

	resultOrm := DB.Model(&model.ReportTest{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetReportTest is a function to get a single record from the report_test table in the go_by_example database
// error - ErrNotFound, db Find error
func GetReportTest(ctx context.Context, argId uint32) (record *model.ReportTest, err error) {
	record = &model.ReportTest{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddReportTest is a function to add a single record to report_test table in the go_by_example database
// error - ErrInsertFailed, db save call failed
func AddReportTest(ctx context.Context, record *model.ReportTest) (result *model.ReportTest, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateReportTest is a function to update a single record from report_test table in the go_by_example database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateReportTest(ctx context.Context, argId uint32, updated *model.ReportTest) (result *model.ReportTest, RowsAffected int64, err error) {

	result = &model.ReportTest{}
	db := DB.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteReportTest is a function to delete a single record from report_test table in the go_by_example database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteReportTest(ctx context.Context, argId uint32) (rowsAffected int64, err error) {

	record := &model.ReportTest{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
