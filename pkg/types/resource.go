package types

// Resources represents a resource type for a task
type Resources struct {
	CPUs  float64 `json:"cpus"`
	Disk  float64 `json:"disk"`
	Mem   float64 `json:"mem"`
	Ports string  `json:"ports"`
}
