package fame

// FanInfo shows the data about the character's influence.
// We may consider 'set'-typed struct.
// And also, we may use the fan info of the character's team.
type FanInfo struct {
	Fan  int `json:"fan"`
	Anti int `json:"anti"`
}
