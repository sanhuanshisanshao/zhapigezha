package scheduler

import "sync"

//Scheduler用于保存源地址，提供给downloader下载
type Scheduler struct {
	sync.Mutex
	urlChan chan string
	urls    map[string]bool
	err     error
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		urlChan: make(chan string, 1000),
		urls:    make(map[string]bool),
		err:     nil,
	}
}

func (s *Scheduler) PutUrl(url string) {
	s.Lock()
	defer s.Unlock()
	s.urlChan <- url
	s.urls[url] = true
}

func (s *Scheduler) GetUrl() chan string {
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
