package objectx

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

func TransformData(from interface{}, to interface{}) error {
	reqData, err := object.StructToHashMap(from)
	err = object.HashMapToStructure(reqData, to)

	return err
}
