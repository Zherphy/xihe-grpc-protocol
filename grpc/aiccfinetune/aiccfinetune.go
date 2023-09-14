package aiccfinetune

type AICCFinetuneIndex struct {
	Id    string
	User  string
	Model string
}

type AICCFinetuneInfo struct {
	Duration      int
	Status        string
	LogPath       string
	OutputZipPath string
}

type AICCFinetuneService interface {
	SetAICCFinetuneInfo(*AICCFinetuneIndex, *AICCFinetuneInfo) error
}
