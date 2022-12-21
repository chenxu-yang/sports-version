package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"wxcloudrun-golang/internal/app/court"
	"wxcloudrun-golang/internal/app/event"
	"wxcloudrun-golang/internal/app/user"
	"wxcloudrun-golang/internal/app/video"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Service struct {
	UserService  *user.Service
	CourtService *court.Service
	EventService *event.Service
	VideoService *video.Service
}

func NewService() *Service {
	return &Service{
		UserService:  user.NewService(),
		CourtService: court.NewService(),
		EventService: event.NewService(),
	}
}

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func (s *Service) WeChatLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := s.UserService.WXLogin(code)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	// 保存登录态
	session := sessions.Default(c)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)
	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	mySession := user.GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	c.String(200, mySession)
}

func (s *Service) StartEvent(c *gin.Context) {
	userOpenIDString := c.Request.Header["X-WX-OPENID"]
	openID, _ := strconv.Atoi(userOpenIDString[0])
	body, _ := ioutil.ReadAll(c.Request.Body)
	event := &model.Event{}
	err := json.Unmarshal(body, event)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	newEvent, err := s.EventService.CreateEvent(int32(openID), event.CourtID, event.StartTime, event.EndTime)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(newEvent, err))
}

// 主页面相关

// 获取推荐视频
func (s *Service) GetRecommendVideos(c *gin.Context) {
	limit := c.Query("limit")
	limitInt, _ := strconv.Atoi(limit)
	videos, err := s.VideoService.GetByDescRank(int32(limitInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(videos, err))
}

// 收藏视频
func (s *Service) CollectVideo(c *gin.Context) {
	userOpenIDString := c.Request.Header["X-WX-OPENID"]
	openID, _ := strconv.Atoi(userOpenIDString[0])
	videoID := c.Query("videoID")
	videoIDInt, _ := strconv.Atoi(videoID)
	collectRecord, err := s.VideoService.CollectVideo(&model.Collect{OpenId: int32(openID), VideoId: int32(videoIDInt)})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(collectRecord, err))
}

// 获取场地, TODO(按位置排序)
func (s *Service) GetCounts(c *gin.Context) {
	limit := c.Query("limit")
	limitInt, _ := strconv.Atoi(limit)
	counts, err := s.CourtService.GetCourtsWithLimit(int32(limitInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(counts, err))
}

func (s *Service) GetCountInfo(c *gin.Context) {
	countID := c.Query("countID")
	countIDInt, _ := strconv.Atoi(countID)
	countInfo, err := s.CourtService.GetCountInfo(int32(countIDInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(countInfo, err))
}

// 获取用户所属事件的视频
func (s *Service) GetEventVideos(c *gin.Context) {
	userOpenIDString := c.Request.Header["X-WX-OPENID"]
	openID, _ := strconv.Atoi(userOpenIDString[0])
	events, err := s.EventService.GetEventsByUser(int32(openID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	url := []string{}
	for _, event := range events {
		url = append(url, fmt.Sprintf("%d%s%s", event.CourtID, event.StartTime, event.StartTime))
	}
	c.JSON(200, resp.ToStruct(url, err))
}

// 获取用户收藏的视频
func (s *Service) GetCollectVideos(c *gin.Context) {
	userOpenIDString := c.Request.Header["X-WX-OPENID"]
	openID, _ := strconv.Atoi(userOpenIDString[0])
	collects, err := s.VideoService.GetCollectByUser(int32(openID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	url := []string{}
	for _, collect := range collects {
		url = s.VideoService.GetVideoUrl(collect.VideoId)
	}
	c.JSON(200, resp.ToStruct(url, err))
}
