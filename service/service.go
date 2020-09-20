package service

type OperationInfo struct {
	Run            bool
	Saving         bool
	RealtimeLength int
	Bots           []string
}

type Service struct {
	ServiceID  string
	EngineType EngineType
	Bots       []string
	OI         OperationInfo `json:"operationInfo"`

	heartbeat chan bool
}
type realtimeDecision struct {
	Timestamp  int64      `json:"timestamp"`
	EngineType EngineType `json:"engineType"`
	ServiceID  string     `json:"serviceId"`
	Decision   decision   `json:"decision"`
}
type decision struct {
	Detection   int `faker:"boundary_start=0, boundary_end=100" json:"detection"`
	Statistics  int `faker:"boundary_start=0, boundary_end=100" json:"statistics"`
	Median      int `faker:"boundary_start=0, boundary_end=100" json:"median"`
	Sensitivity int `faker:"boundary_start=0, boundary_end=100" json:"sensitivity"`
}

// NewService represent to create a new service
func NewService(s *Service, realtimeLength int) *Service {
	s.heartbeat = make(chan bool, 1)
	s.OI.RealtimeLength = realtimeLength
	return s
}
