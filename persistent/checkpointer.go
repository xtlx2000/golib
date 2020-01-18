package persistent

type Checkpointer interface {
	Load(filePath string) error
	Save(filePath string) error
}
