package app

type TagService struct {
}

/*
func (ts TagService) List() ([]string, error) {
	var err error
	Db := utils.GetDB()
	tags := make([]model.Tag, 20)
	var tagNames []string

	if err = Db.Raw("SELECT id,tag_name FROM " + command.DBTag).Scan(&tags).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	for _, v := range tags {
		tagNames = append(tagNames, v.TagName)
	}

	return tagNames, err
}
*/
