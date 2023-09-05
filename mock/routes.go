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

var RoutesMock = []Route{
	{
		Name:     "test",
		Path:     "/test",
		Redirect: "/test/test1",
		Meta: Meta{
			Title:       "测试",
			KeepAlive:   false,
			RequireAuth: false,
		},
		Children: []Route{
			{
				Path:      "test1",
				Component: "/test/test1",
				Meta: Meta{
					Title:       "测试1",
					KeepAlive:   false,
					RequireAuth: false,
				},
			},
			{
				Path:      "test2",
				Component: "/test/test2",
				Meta: Meta{
					Title:       "测试2",
					KeepAlive:   false,
					RequireAuth: false,
				},
			},
		},
	},
	{
		Path:     "/user",
		Name:     "user",
		Redirect: "/user/my",
		Meta: Meta{
			Title:       "个人中心",
			KeepAlive:   true,
			RequireAuth: true,
		},
		Children: []Route{
			{
				Path:      "my",
				Name:      "my",
				Component: "/user/my",
				Meta: Meta{
					Title:       "我的主页",
					KeepAlive:   true,
					RequireAuth: true,
				},
			},
			{
				Path:      "course",
				Name:      "course",
				Component: "/user/course",
				Meta: Meta{
					Title:       "我的课程",
					KeepAlive:   true,
					RequireAuth: true,
				},
			},
		},
	},
}
