package models

type ViewPoint struct {
	ID        int     `json:"id" gorm:"column:id;autoincrement;primaryKey"`
	UserID    int     `json:"user_id" gorm:"column:user_id;not null"`
	CenterX   float32 `json:"center_x" gorm:"column:center_x"`
	CenterY   float32 `json:"center_y" gorm:"column:center_y"`
	Zoom      float32 `json:"zoom" gorm:"column:zoom"`
	Pitch     float32 `json:"pitch" gorm:"column:pitch"`
	Bearing   float32 `json:"bearing" gorm:"column:bearing"`
	Name      string  `json:"name" gorm:"column:name"`
	PointType string  `json:"point_type" gorm:"column:point_type"`
}

func CreateViewPoint(user_id int, center_x float32, center_y float32, zoom float32, pitch float32, bearing float32, name string, pointType string) (ViewPoint, error) {

	// 創建 ViewPoint 對象
	viewpoint := ViewPoint{
		UserID:    user_id,
		CenterX:   center_x,
		CenterY:   center_y,
		Zoom:      zoom,
		Pitch:     pitch,
		Bearing:   bearing,
		Name:      name,
		PointType: pointType,
	}

	// 將 ViewPoint 對象插入到數據庫中
	if err := DBManager.Create(&viewpoint).Error; err != nil {
		return ViewPoint{}, err
	}

	return viewpoint, nil
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
