package cloud

type CloudPod struct {
	Id         string
	User       string
	ProjectId  string
	LastCommit string
}

type PodInfo struct {
	Error     string
	AccessURL string
}

type CloudService interface {
	SetCloudInfo(*CloudPod, *PodInfo) error
}
