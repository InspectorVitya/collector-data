package db

type InfoDevice struct {
	ID               uint32 `bson:"id" json:"id"`
	NameOS           string `bson:"name_os" json:"nameOS"`
	VersionOS        string `bson:"version_os" json:"versionOS"`
	NameBrowser      string `bson:"name_browser" json:"nameBrowser"`
	VersionBrowser   string `bson:"version_browser" json:"versionBrowser"`
	IP               string `bson:"ip" json:"IP"`
	BrandPhone       string `bson:"brand_phone" json:"brandPhone"`
	ModelPhone       string `bson:"model_phone" json:"modelPhone"`
	ScreenResolution string `bson:"screen_resolution" json:"screenResolution"`
}

type Top struct {
	Name  string `bson:"_id" json:"name"`
	Count uint32 `bson:"count" json:"count"`
}
