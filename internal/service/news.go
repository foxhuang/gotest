package service

import ( 
	"go-service/internal/gen"
	"time"
)
 
// GetNews 取得公開且在有效期間的新聞
func (svc *Service) GetNews(limit int) ([]gen.News, error) {
	var newsList []gen.News
	now := time.Now()
	err := svc.DB.
		Where("ispublic = ? AND ? >= begindate AND (enddate >= ? OR enddate IS NULL)", 1, now, now).
		Order("important desc, begindate desc, id asc").
		Limit(limit).
		Find(&newsList).Error
	return newsList, err
}

// GetNewsCount 取得公開且在有效期間的新聞數量
func (svc *Service) GetNewsCount() (int64, error) {
	var count int64
	now := time.Now()
	err := svc.DB.
		Model(&gen.News{}).
		Where("ispublic = ? AND ? >= begindate AND (enddate >= ? OR enddate IS NULL)", 1, now, now).
		Count(&count).Error
	return count, err
}

// GetNewsWithForm 分頁取得新聞
func (svc *Service) GetNewsWithForm(form *gen.NewsForm) ([]gen.News, error) {
	var newsList []gen.News
	now := time.Now()
	page := 1
	if form.Page != nil {
		page = *form.Page
	}
	offset := (page - 1) * *form.Size
	err := svc.DB.
		Where("ispublic = ? AND ? >= begindate AND (enddate >= ? OR enddate IS NULL)", 1, now, now).
		Order("important desc, begindate desc, id asc").
		Offset(offset).
		Limit(*form.Size).
		Find(&newsList).Error
	return newsList, err
}

// GetNewsCountWithForm 取得分頁條件下的新聞數量
func (svc *Service) GetNewsCountWithForm(form *gen.NewsForm) (int64, error) {
	var count int64
	now := time.Now()
	err := svc.DB.
		Model(&gen.News{}).
		Where("ispublic = ? AND ? >= begindate AND (enddate >= ? OR enddate IS NULL)", 1, now, now).
		Count(&count).Error
	return count, err
}

// GetNewsAttachment 取得指定新聞的附件
func (svc *Service) GetNewsAttachment(newsId int) ([]gen.NewsAttachment, error) {
	var attachments []gen.NewsAttachment
	err := svc.DB.
		Where("newsid = ?", newsId).
		Order("insertDate asc").
		Find(&attachments).Error
	return attachments, err
}

// Load 依 ID 取得新聞
func (svc *Service) Load(id int) (*gen.News, error) {
	var news gen.News
	err := svc.DB.First(&news, id).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}
