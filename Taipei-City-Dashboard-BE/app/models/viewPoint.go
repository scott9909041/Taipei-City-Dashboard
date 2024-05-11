package models

type ViewPoint struct {
	ID      int     `json:"id" gorm:"column:id;autoincrement;primaryKey"`
	UserID  int     `json:"user_id" gorm:"column:user_id;not null"`
	Center  string  `json:"center" gorm:"column:center;type:float[]"`
	Zoom    float32 `json:"zoom" gorm:"column:zoom"`
	Pitch   float32 `json:"pitch" gorm:"column:pitch"`
	Bearing float32 `json:"bearing" gorm:"column:bearing"`
}

func CreateViewPoint(user_id int, center string, zoom float32, pitch float32, bearing float32) error {
	// 創建 ViewPoint 對象
	viewpoint := ViewPoint{
		UserID:  user_id,
		Center:  center,
		Zoom:    zoom,
		Pitch:   pitch,
		Bearing: bearing,
	}

	// 將 ViewPoint 對象插入到數據庫中
	if err := DBManager.Create(&viewpoint).Error; err != nil {
		return err
	}

	return nil
}

func GetViewPointByUserID(user_id int) ([]ViewPoint, error) {
	var viewpoint []ViewPoint

	err := DBManager.Where("user_id = ?", user_id).Find(&viewpoint).Error
	if err != nil {
		return []ViewPoint{}, err
	}

	return viewpoint, nil
}

func DeleteViewPoint(id int) error {
	err := DBManager.Delete(&ViewPoint{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
