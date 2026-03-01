
package geolocation

type BoundingBox struct {
	SW Point `json:"sw" bson:"sw"`
	NE Point `json:"ne" bson:"ne"`
}