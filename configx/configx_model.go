package configx

type FieldCommentConfig map[string]string

type CommentedConfig struct {
	Data     interface{}        `json:"data" binding:"required" yaml:"-"`
	Comments FieldCommentConfig `json:"comments" yaml:"-"`
}
