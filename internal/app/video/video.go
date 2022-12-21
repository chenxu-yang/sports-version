package video

import "wxcloudrun-golang/internal/pkg/model"

type Service struct {
	videoDao   *model.Video
	CollectDao *model.Collect
}

func NewService() *Service {
	return &Service{
		videoDao:   &model.Video{},
		CollectDao: &model.Collect{},
	}
}

func (s *Service) GetVideoUrl(videoID int32) (*model.Video, error) {
	video, err := s.videoDao.Get(&model.Video{ID: videoID})
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (s *Service) GetByDescRank(limit int32) ([]model.Video, error) {
	results, err := s.videoDao.GetByDescRank(limit)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Service) CollectVideo(collect *model.Collect) (*model.Collect, error) {
	// 查询是否已经收藏过
	collect, err := s.CollectDao.Get(collect)
	if err != nil {
		return nil, err
	}
	if collect.ID > 0 {
		return collect, nil
	}

	// 创建收藏
	collect, err = s.CollectDao.Create(collect)
	if err != nil {
		return nil, err
	}
	return collect, nil
}

func (s *Service) GetCollectByUser(userOpenID int32) ([]model.Collect, error) {
	collects, err := s.CollectDao.Gets(&model.Collect{OpenId: userOpenID})
	if err != nil {
		return nil, err
	}
	return collects, nil
}
