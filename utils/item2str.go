package utils

import (
	"DsaGo/dsago/geom/Point"
	"DsaGo/dsago/geom/Point3D"
	"DsaGo/dsago/geom/Vector3D"
	"strconv"
)

//func item2str(item interface{}) string {
//	if reflect.TypeOf(item).Kind().String() == "struct" {
//		return item.ToString()
//	}
//}

func Point2str(point interface{}) string {
	return point.(*Point.Point).Point2Str()
}

func Point3D2str(point *Point3D.Point3D) string {
	return point.Point2Str()
}

func Vector2str(vector Vector3D.Vector3D) string {
	return vector.Vector2Str()
}

func int2str(value *int) string {
	return strconv.Itoa(*value)
}
