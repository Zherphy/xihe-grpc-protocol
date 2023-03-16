package cloud

type CloudPod struct {
	Id string
}

type PodInfo struct {
	Error     string
	AccessURL string
}

type CloudService interface {
	SetPodInfo(*CloudPod, *PodInfo) error
}
