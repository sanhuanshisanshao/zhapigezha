package scheduler

import (
	"sync"
	"zhapigezha/models"
)

//Scheduler用于保存源地址，提供给downloader下载
type Scheduler struct {
	sync.Mutex
	urlChan chan models.SourceInfo
	urls    map[string]models.SourceType
	err     error
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		urlChan: make(chan models.SourceInfo, 1000),
		urls:    make(map[string]models.SourceType),
		err:     nil,
	}
}

func (s *Scheduler) PutUrl(url string, urlType models.SourceType) {
	s.Lock()
	defer s.Unlock()
	s.urlChan <- models.SourceInfo{Url: url, SourceType: urlType}
	s.urls[url] = urlType
}

func (s *Scheduler) GetUrl() chan models.SourceInfo {
	return s.urlChan
}

//移除urls map中的key
func (s *Scheduler) RemoveKey(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.urls, key)
}

//遍历map中的值
func (s *Scheduler) RangeMap() []string {
	s.Lock()
	defer s.Unlock()
	list := make([]string, 0)
	for k, _ := range s.urls {
		list = append(list, k)
	}
	return list
}
