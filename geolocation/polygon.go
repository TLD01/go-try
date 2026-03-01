package geolocation

type Polygon struct {
	Coordinates []Point `json:"coordinates" bson:"coordinates"`
}


func (p *Polygon) String() string {

	if p.Coordinates == nil {
		return "[]"
	}

	result := "["
	for i, point := range p.Coordinates {
		result += point.String()
		if i < len(p.Coordinates)-1 {
			result += ","
		}
	}
	result += "]"
	return result
}