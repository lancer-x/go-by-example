package dao

import (
	"context"
	"time"

	"go-by-example/examples/lgorm/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllGrowthPublisherChannelConfig is a function to get a slice of record(s) from growth_publisher_channel_config table in the guazi_sem database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllGrowthPublisherChannelConfig(ctx context.Context, page, pagesize int64, order string) (results []*model.GrowthPublisherChannelConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&model.GrowthPublisherChannelConfig{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(int(offset)).Limit(int(pagesize))
	} else {
		resultOrm = resultOrm.Limit(int(pagesize))
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

// GetGrowthPublisherChannelConfig is a function to get a single record from the growth_publisher_channel_config table in the guazi_sem database
// error - ErrNotFound, db Find error
func GetGrowthPublisherChannelConfig(ctx context.Context, argId uint32) (record *model.GrowthPublisherChannelConfig, err error) {
	record = &model.GrowthPublisherChannelConfig{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddGrowthPublisherChannelConfig is a function to add a single record to growth_publisher_channel_config table in the guazi_sem database
// error - ErrInsertFailed, db save call failed
func AddGrowthPublisherChannelConfig(ctx context.Context, record *model.GrowthPublisherChannelConfig) (result *model.GrowthPublisherChannelConfig, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGrowthPublisherChannelConfig is a function to update a single record from growth_publisher_channel_config table in the guazi_sem database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateGrowthPublisherChannelConfig(ctx context.Context, argId uint32, updated *model.GrowthPublisherChannelConfig) (result *model.GrowthPublisherChannelConfig, RowsAffected int64, err error) {

	result = &model.GrowthPublisherChannelConfig{}
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

// DeleteGrowthPublisherChannelConfig is a function to delete a single record from growth_publisher_channel_config table in the guazi_sem database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteGrowthPublisherChannelConfig(ctx context.Context, argId uint32) (rowsAffected int64, err error) {

	record := &model.GrowthPublisherChannelConfig{}
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
