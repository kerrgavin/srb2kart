package main

type hostmod struct {
	Automate *bool `json:"automate"`
	Encore *int `json:"encore"`
	Specbomb specbomb `json:"specbomb"`
	NameFilter nameFilter `json:"namefilter"`
	Battle battle `json:"battle"`
	Motd motd `json:"motd"`
	Restat restat `json:"restat"`
	Shout shout `json:"shout"`
	Schedule schedule `json:"schedule"`
	Vote vote `json:"vote"`
	CustomCheck customcheck `json:"customcheck"`
	Scoreboard scoreboard `json:"scoreboard"`
	Veto veto `json:"veto"`
}

type specbomb struct {
	Enabled *bool `json:"enabled"`
	AntiSoftlock *bool `json:"antisoftlock"`
}

type nameFilter struct {
	Mode *string `json:"mode"`
	Names []string `json:"names"`
}

type battle struct {
	Timelimit *int `json:"timelimit"`
	Bail *bool `json:"bail"`
}

type motd struct {
	Enabled *bool `json:"enabled"`
	Nag *bool `json:"nag"`
	Background *string `json:"background"`
	Name *string `json:"name"`
	Contact *string `json:"contact"`
	Tagline *string `json:"tagline"`
}

type restat struct {
	Enabled *bool `json:"enabled"`
	Notify *bool `json:"notify"`
}

type shout struct {
	Autoshout *bool `json:"autoshout"`
	Name *string `json:"name"`
	Color *string `json:"color"`
}

type schedule struct {
	Enabled *bool `json:"enabled"`
	Jobs []job `json:"jobs"`
}

type job struct {
	Timer *int `json:"timer"`
	Commands []string `json:"commands"`
}

type vote struct {
	Whitelist []string `json:"whitelist"`
	Timer *int `json:"timer"`
	Autopass *int `json:"autopass"`
	Allowidc *bool `json:"allowidc"`
}

type customcheck struct {
	Enabled *bool `json:"enabled"`
	CheckOne *bool `json:"checkone"`
	CheckTwo *bool `json:"checktwo"`
	CheckThree *bool `json:"checkthree"`
}

type scoreboard struct {
	Enabled *bool `json:"enabled"`
	Humor *bool `json:"humor"`
	Lines []string `json:"lines"`
	ModLines []string `json:"modlines"`
}

type veto struct {
	Enabled *bool `json:"enabled"`
	Threshold *int `json:"threshold"`
	Hellclosed *int `json:"hellclosed"`
	Hellopen *int `json:"hellopen"`
}
