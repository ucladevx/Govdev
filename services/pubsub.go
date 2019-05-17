package services

type PubSubStore interface {
	Publish(queue string, value string) error
	Subscribe(queueid string, msg chan string) error
}

// PubSubService wraps a PubSubStore, implementing Publish() and Subscribe()
// methods
type PubSubService struct {
	store PubSubStore
}

func NewPubSubService(store PubSubStore) *PubSubService {
	return &PubSubService{
		store: store,
	}
}

func (pbs *PubSubService) Publish(queueid, value string) error {
	return pbs.store.Publish(queueid, value)
}

func (pbs *PubSubService) Subscribe(queueid string, msg chan string) error {
	return pbs.store.Subscribe(queueid, msg)
}
