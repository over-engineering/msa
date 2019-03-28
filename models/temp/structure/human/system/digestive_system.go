package system

import (
	common "temp/structure/common"
)

type DigestiveSystem struct {
	Liver     common.Organ `json:"liver"`
	Stomach   common.Organ `json:"stomach"`
	Intestine common.Organ `json:"intestine"`
}
