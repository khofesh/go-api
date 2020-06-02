package menumodel

import "go.mongodb.org/mongo-driver/bson/primitive"

// Menu : Menu shown to a user
type Menu struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string               `bson:"name" json:"name"`
	Caption    string               `bson:"caption" json:"caption"`
	RoleCode   string               `bson:"rolecode" json:"rolecode"`
	RoleName   string               `bson:"rolename" json:"rolename"`
	Info       string               `bson:"info" json:"info"`
	canAdd     bool                 `bson:"canAdd" json:"canAdd"`
	canEdit    bool                 `bson:"canEdit" json:"canEdit"`
	canView    bool                 `bson:"canView" json:"canView"`
	canDelete  bool                 `bson:"canDelete" json:"canDelete"`
	EntryDate  string               `bson:"entryDate" json:"entryDate"`
	UpdateDate string               `bson:"updateDate" json:"updateDate"`
	UserID     []primitive.ObjectID `bson:"user_id" json:"user_id"`
}
