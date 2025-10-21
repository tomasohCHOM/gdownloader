package options

const (
	DOWNLOAD = "Download files"
	PATH     = "Manage paths"
	EXIT     = "Exit"
)

const (
	ADD_PATH    = "Add path"
	REMOVE_PATH = "Remove path"
	LIST_PATHS  = "List paths"
)

var ROOT_CMD_OPTIONS = []string{DOWNLOAD, PATH, EXIT}
var PATH_CMD_OPTIONS = []string{ADD_PATH, REMOVE_PATH, LIST_PATHS, EXIT}

const DOWNlOAD_QUERY_MORE_PROMPT = "Search next page"
