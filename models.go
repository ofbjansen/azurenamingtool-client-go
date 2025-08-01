package azurenamingtool

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `json:"coffee"`
	Quantity int    `json:"quantity"`
}

// Coffee -
type Coffee struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Teaser      string             `json:"teaser"`
	Collection  string             `json:"collection"`
	Origin      string             `json:"origin"`
	Color       string             `json:"color"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	Ingredient  []CoffeeIngredient `json:"ingredients"`
}

// Ingredient -
type CoffeeIngredient struct {
	ID       int    `json:"ingredient_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Ingredient -
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// ResourceTypes -
type ResourceTypes struct {
	ID                           int    `json:"id"`
	Name                         string `json:"name"`
	resource                     string `json:"resource"`
	optional                     string `json:"optional"`
	exclude                      string `json:"exclude"`
	property                     string `json:"property"`
	ShortName                    string `json:"short_name"`
	scope                        string `json:"scope"`
	lengthMin                    string `json:"length_min"`
	lengthMax                    string `json:"length_max"`
	validText                    string `json:"valid_text"`
	invalidText                  string `json:"invalid_text"`
	invalidCharacters            string `json:"invalid_characters"`
	invalidCharactersStart       string `json:"invalid_characters_start"`
	invalidCharactersEnd         string `json:"invalid_characters_end"`
	invalidCharactersConsecutive string `json:"invalid_characters_consecutive"`
	regx                         string `json:"regx"`
	staticValues                 string `json:"static_values"`
	enabled                      bool   `json:"enabled"`
	applyDelimiter               bool   `json:"apply_delimiter"`
}
