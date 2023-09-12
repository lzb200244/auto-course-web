package mock

type Route struct {
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Redirect  string  `json:"redirect"`
	Meta      Meta    `json:"meta"`
	Children  []Route `json:"children"`
	Component string  `json:"component"`
}

type Meta struct {
	Title       string `json:"title"`
	KeepAlive   bool   `json:"keepAlive"`
	RequireAuth bool   `json:"requireAuth"`
}
