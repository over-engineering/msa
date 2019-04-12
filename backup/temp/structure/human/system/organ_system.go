package system

type OrganSystem struct {
	Respiratory RespiratorySystem
	Circulatory CirculatorySystem
	Digestive   DigestiveSystem
	Excretory   ExcretorySystem
}
