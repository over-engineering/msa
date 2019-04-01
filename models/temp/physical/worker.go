package physical

type Worker interface {
	GetHeight() float32
	GetWeight() float32
	GetForce(ForceType) float32
	GetHealth(HealthType) float32
}
