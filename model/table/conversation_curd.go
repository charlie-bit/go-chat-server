package relation

import (
	"context"

	"gorm.io/gorm"
)

type ConversationModelInter interface {
	Create(ctx context.Context, conversations []*ConversationModel) (err error)
	Delete(ctx context.Context, groupIDs []string) (err error)
	UpdateByMap(ctx context.Context, userIDs []string, conversationID string, args map[string]interface{}) (
		rows int64, err error,
	)
	Update(ctx context.Context, conversation *ConversationModel) (err error)
	Find(ctx context.Context, ownerUserID string, conversationIDs []string) (
		conversations []*ConversationModel, err error,
	)
	FindUserID(ctx context.Context, userIDs []string, conversationIDs []string) ([]string, error)
	FindUserIDAllConversationID(ctx context.Context, userID string) ([]string, error)
	Take(ctx context.Context, userID, conversationID string) (conversation *ConversationModel, err error)
	FindConversationID(ctx context.Context, userID string, conversationIDs []string) (
		existConversationID []string, err error,
	)
	FindUserIDAllConversations(ctx context.Context, userID string) (conversations []*ConversationModel, err error)
	GetAllConversationIDs(ctx context.Context) ([]string, error)
	GetUserAllHasReadSeqs(ctx context.Context, ownerUserID string) (hashReadSeqs map[string]int64, err error)
	GetConversationsByConversationID(ctx context.Context, conversationIDs []string) ([]*ConversationModel, error)
	GetConversationIDsNeedDestruct(ctx context.Context) ([]*ConversationModel, error)
	NewTx(tx any) ConversationModelInter
}

type ConversationGorm struct {
	DB     *gorm.DB
	TModel any
}

func NewConversationGorm(db *gorm.DB) ConversationModelInter {
	return &ConversationGorm{db, &ConversationModel{}}
}

func (c *ConversationGorm) NewTx(tx any) ConversationModelInter {
	return &ConversationGorm{tx.(*gorm.DB), &ConversationModel{}}
}

func (c *ConversationGorm) Create(ctx context.Context, conversations []*ConversationModel) (err error) {
	return c.DB.Create(&conversations).Error
}

func (c *ConversationGorm) Delete(ctx context.Context, groupIDs []string) (err error) {
	return c.DB.Where("group_id in (?)", groupIDs).Delete(&ConversationModel{}).Error
}

func (c *ConversationGorm) UpdateByMap(
	ctx context.Context,
	userIDList []string,
	conversationID string,
	args map[string]interface{},
) (rows int64, err error) {
	result := c.DB.Where("owner_user_id IN (?) and  conversation_id=?", userIDList, conversationID).Updates(args)
	return result.RowsAffected, nil
}

func (c *ConversationGorm) Update(ctx context.Context, conversation *ConversationModel) (err error) {
	return c.DB.Where(
		"owner_user_id = ? and conversation_id = ?", conversation.OwnerUserID, conversation.ConversationID,
	).Updates(conversation).Error
}

func (c *ConversationGorm) Find(
	ctx context.Context,
	ownerUserID string,
	conversationIDs []string,
) (conversations []*ConversationModel, err error) {
	err = c.DB.
		Where("owner_user_id=? and conversation_id IN (?)", ownerUserID, conversationIDs).
		Find(&conversations).
		Error
	return conversations, err
}

func (c *ConversationGorm) Take(
	ctx context.Context,
	userID, conversationID string,
) (conversation *ConversationModel, err error) {
	cc := &ConversationModel{}
	return cc, c.DB.Where("conversation_id = ? And owner_user_id = ?", conversationID, userID).Take(cc).Error
}

func (c *ConversationGorm) FindUserID(
	ctx context.Context,
	userIDs []string,
	conversationIDs []string,
) (existUserID []string, err error) {
	return existUserID, c.DB.
		Where(" owner_user_id IN (?) and conversation_id in (?)", userIDs, conversationIDs).
		Pluck("owner_user_id", &existUserID).
		Error
}

func (c *ConversationGorm) FindConversationID(
	ctx context.Context,
	userID string,
	conversationIDList []string,
) (existConversationID []string, err error) {
	return existConversationID, c.DB.
		Where(" conversation_id IN (?) and owner_user_id=?", conversationIDList, userID).
		Pluck("conversation_id", &existConversationID).
		Error
}

func (c *ConversationGorm) FindUserIDAllConversationID(
	ctx context.Context,
	userID string,
) (conversationIDList []string, err error) {
	return conversationIDList, c.DB.Where("owner_user_id=?", userID).Pluck("conversation_id", &conversationIDList).Error
}

func (c *ConversationGorm) FindUserIDAllConversations(
	ctx context.Context,
	userID string,
) (conversations []*ConversationModel, err error) {
	return conversations, c.DB.Where(
		"owner_user_id=?", userID,
	).Find(&conversations).Error
}

func (c *ConversationGorm) GetUserRecvMsgOpt(
	ctx context.Context,
	ownerUserID, conversationID string,
) (opt int, err error) {
	var conversation ConversationModel
	return int(
			conversation.RecvMsgOpt,
		), c.DB.
			Where("conversation_id = ? And owner_user_id = ?", conversationID, ownerUserID).
			Select("recv_msg_opt").
			Find(&conversation).
			Error
}

func (c *ConversationGorm) GetAllConversationIDs(ctx context.Context) (conversationIDs []string, err error) {
	return conversationIDs, c.DB.Distinct("conversation_id").Pluck("conversation_id", &conversationIDs).Error
}

func (c *ConversationGorm) GetUserAllHasReadSeqs(
	ctx context.Context,
	ownerUserID string,
) (hasReadSeqs map[string]int64, err error) {
	return nil, nil
}

func (c *ConversationGorm) GetConversationsByConversationID(
	ctx context.Context,
	conversationIDs []string,
) (conversations []*ConversationModel, err error) {
	return conversations, c.DB.Where("conversation_id IN (?)", conversationIDs).Find(&conversations).Error
}

func (c *ConversationGorm) GetConversationIDsNeedDestruct(
	ctx context.Context,
) (conversations []*ConversationModel, err error) {
	return conversations, c.DB.
		Where("is_msg_destruct = 1 && msg_destruct_time != 0 && (UNIX_TIMESTAMP(NOW()) > (msg_destruct_time + UNIX_TIMESTAMP(latest_msg_destruct_time)) || latest_msg_destruct_time is NULL)").
		Find(&conversations).
		Error
}
