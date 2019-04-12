package human

import (
	physical "temp/physical"
	body "temp/structure/human/body"
	system "temp/structure/human/system"
)

type Human struct {
	Bust      body.Bust
	UpperBody body.Body
	LeftArm   body.Arm
	RightArm  body.Arm
	LeftLeg   body.Leg
	RightLeg  body.Leg

	OrganSystem system.OrganSystem

	//Temporary property for 1st testing
	Height  float32               `json:"height"`
	Weight  float32               `json:"weight"`
	Joints  map[JointType]Joint   `json:"joints"`
	Muscles map[MuscleType]Muscle `json:"muscles"`
	Fat     float32               `json:"fat"`
}

// Method implementation for Worker interface
func (h *Human) GetForce(fType physical.ForceType) float32 {

}
