package app

type TypeService struct {
}

/*
func (ts TypeService) List() ([]model.Type, error) {
	Db := utils.GetDB()
	types := make([]model.Type, 4)
	if err := Db.Raw("SELECT id,type_name,amount FROM " + command.DBType + " ORDER BY amount DESC").Scan(&types).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	return types, nil
}

func (ts TypeService) typeList(id int) ([]vo.IndexVO, error) {

	Db := utils.GetDB()

	var blogs []vo.IndexVO
	if err := Db.Raw(`SELECT b.id,b.title,b.content,b.update_time,t.type_name,u.avatar,u.username
						FROM t_blog b JOIN t_user u ON u.id = b.user_id JOIN t_type t ON b.type_id = t.id AND b.type_id = ?`, id).Scan(&blogs).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	for k, v := range blogs {
		if err := Db.Raw(`select t.id,t.tag_name from t_tag t JOIN
					(select DISTINCT(bt.tag_id) from t_blog_tag bt JOIN t_blog b ON bt.blog_id = ?) as tag
					ON t.id = tag.tag_id`, v.Id).Scan(&blogs[k].TagNames).Error; err != nil {
			return nil, errors.New("查询失败")
		}
	}

	return blogs, nil
}
*/
