package pokeapi

type Stats struct {
   BaseStat    int    `json:"base_stat"`
   Effort      int     `json:"effort"`
   Stat        struct{
	 Name      string   `json:"name"`
	 URL       string    `json:"url"`
   }`json:"stat"`
}

type Types struct{
	Slot    int     `json:"slot"`
	Type    struct{
	 Name      string   `json:"name"`
	 URL       string    `json:"url"`
   }`json:"type"`
}

type Pokemon struct {
	Id      int          `json:"id"`
	Name    string       `json:"name"`
	BaseExperience  int  `json:"base_experience"`
	Order           int  `json:"order"`
	Weight         int    `json:"weight"`
	Height        int     `json:"height"`
	Stats         []Stats   `json:"stats"`
    Types         []Types     `json:"types"`

}
