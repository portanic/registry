package models

type Plugin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
	RepoURL     string `json:"repo_url"`
	Version     string `json:"version"`
}

type PluginRegistry struct {
	Plugins map[string]Plugin
}

var Registry = PluginRegistry{
	Plugins: make(map[string]Plugin),
}
