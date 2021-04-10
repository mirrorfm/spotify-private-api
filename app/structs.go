package app

type OpsMov struct {
	FromIndex int `json:"fromIndex"`
	Length    int `json:"length"`
	ToIndex   int `json:"toIndex"`
}

type DeltaOps struct {
	Kind string `json:"kind"`
	Mov  OpsMov `json:"mov"`
}

type ContentsItemAttributes struct {
	Public    bool   `json:"public"`
	Timestamp string `json:"Timestamp"`
}

type ContentsItem struct {
	Uri        string                 `json:"uri"`
	Attributes ContentsItemAttributes `json:"attributes"`
}

type MetaItemAttributes struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

type MetaItems struct {
	Length        int    `json:"length"`
	OwnerUsername string `json:"ownerUsername"`
	Timestamp     string `json:"timestamp"`
	Revision      string `json:"revision"`
}

type RootListResponseContents struct {
	Items     []ContentsItem `json:"items"`
	MetaItems []MetaItems    `json:"metaItems"`
	Pos       int            `json:"pos"`
	Truncated bool           `json:"truncated"`
}

type RootListResponse struct {
	Length    int                      `json:"length"`
	Revision  string                   `json:"revision"`
	Timestamp string                   `json:"timestamp"`
	Contents  RootListResponseContents `json:"contents"`
}

type ChangeDelta struct {
	Ops []DeltaOps `json:"ops"`
}

type ChangesPayload struct {
	BaseRevision string        `json:"baseRevision"`
	Deltas       []ChangeDelta `json:"deltas"`
}
