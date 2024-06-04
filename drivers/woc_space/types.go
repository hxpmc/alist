package woc_space

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Space struct {
	Guid string  `json:"guid"`
	Name string  `json:"name"`
	Size float64 `json:"size"`
}

type SpaceResp struct {
	Resp,
	Data struct {
		Space []Space `json:"space"`
	} `json:"data"`
}
