package qcdn

type Flag struct {
	SiteMap string `flag:"sitemap" usage:"SiteMap 文件"`
	Purge   bool   `flag:"purge" usage:"刷新URL"`
	Push    bool   `flag:"push" usage:"预热URL"`
}
