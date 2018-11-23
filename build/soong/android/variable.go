package android
type Product_variables struct {
	target_board struct {
		Cflags []string
        }
	use_samsung_color struct {
		Cflags []string
	}
}

type ProductVariables struct {
	target_board *string `json:",omitempty"`
	use_samsung_color  *bool `json:",omitempty"`


}
