package repo

import (
	"context"
	"time"
)

func FindAllTransactionByDate(ctx context.Context, date time.Time, limit int, page int) ([]Transactions, error) {
	var transactions []Transactions
	result := DB.WithContext(ctx).Where("date between ? and ?", date, date.AddDate(0, 0, 1).Add(time.Millisecond*-1)).Order("id desc").Limit(limit).Offset((page - 1) * limit).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func SaveTransaction(ctx context.Context, transaction *Transactions) error {
	result := DB.WithContext(ctx).Save(&transaction)
	return result.Error
}

func DeleteTransaction(ctx context.Context, id uint) error {
	result := DB.WithContext(ctx).Delete(&Transactions{}, id)
	return result.Error

}

func FindTransactionByID(ctx context.Context, id uint) (*Transactions, error) {
	var transactions Transactions
	result := DB.WithContext(ctx).First(&transactions, id)
	if result.Error != nil {
		return nil, result.Error

	}
	return &transactions, nil
}
