package backend

type CreateVolumeRequest struct {
	Name           string
	RequestedBytes int64
}

type CreateVolumeResponse struct {
	Path     string
	Host     string
	Capacity int64
}
