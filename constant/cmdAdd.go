package constant

type CmdAddFlag struct {
	Org      string
	Email    string
	Name     string
	Phrase   string
	Platform string
}

var CmdAddFlags = CmdAddFlag{
	Org:      "org",
	Email:    "email",
	Name:     "name",
	Phrase:   "phrase",
	Platform: "platform",
}
