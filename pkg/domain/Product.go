package domain

type (
	Product struct {
		Id             uint64         `json:"id" bson:"_id"`
		Name           string         `json:"name" bson:"name"`
		Price          float64        `json:"price" bson:"price"`
		Description    string         `json:"description" bson:"description"`
		Сharacteristic Сharacteristic `json:"characteristic" bson:"characteristic"`
		User_id        string         `bson:"user_id"`
		Rating         float64        `bson:"-"`
	}
	Сharacteristic struct {
		Category string  `json:"category" bson:"category,omitempty"`
		Brand    string  `json:"brand" bson:"brand,omitempty"`
		Size     string  `json:"size" bson:"size,omitempty"`
		Color    string  `json:"color" bson:"color,omitempty"`
		Weight   float32 `json:"weight" bson:"weight,omitempty"`
	}
)
