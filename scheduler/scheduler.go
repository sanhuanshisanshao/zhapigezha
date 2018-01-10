package scheduler

//Scheduler用于保存源地址，提供给downloader下载
type Scheduler struct {
	urls chan string
	err  error
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		urls: make(chan string, 1000),
		err:  nil,
	}
}

func (s *Scheduler) PutUrl(url string) {
	s.urls <- url
}

func (s *Scheduler) GetUrl() chan string {
	return s.urls
}
