package telegram

type User struct {
	ID                        int    `json:"id"`
	IsBot                     bool   `json:"is_bot"`
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name,omitempty"`
	Username                  string `json:"username,omitempty"`
	LanguageCode              string `json:"language_code,omitempty"`
	IsPremium                 bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu     bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups             bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages   bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries     bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness      bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp             bool   `json:"has_main_web_app,omitempty"`
	HasTopicsEnabled          bool   `json:"has_topics_enabled,omitempty"`
	AllowsUsersToCreateTopics bool   `json:"allows_users_to_create_topics,omitempty"`
}

type Chat struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Title            string `json:"title,omitempty"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	IsForum          bool   `json:"is_forum,omitempty"`
	IsDirectMessages bool   `json:"is_direct_messages,omitempty"`
}

type ChatFullInfo struct {
	ID                int64    `json:"id"`
	Type              string   `json:"type"`
	Title             string   `json:"title"`
	Username          string   `json:"username"`
	ActiveUsernames   []string `json:"active_usernames"`
	Description       string   `json:"description"`
	CanSendGift       bool     `json:"can_send_gift"`
	HasVisibleHistory bool     `json:"has_visible_history"`
	CanSendPaidMedia  bool     `json:"can_send_paid_media"`
	AcceptedGiftTypes struct {
		UnlimitedGifts      bool `json:"unlimited_gifts"`
		LimitedGifts        bool `json:"limited_gifts"`
		UniqueGifts         bool `json:"unique_gifts"`
		PremiumSubscription bool `json:"premium_subscription"`
	} `json:"accepted_gift_types"`
	Photo struct {
		SmallFileID       string `json:"small_file_id"`
		SmallFileUniqueID string `json:"small_file_unique_id"`
		BigFileID         string `json:"big_file_id"`
		BigFileUniqueID   string `json:"big_file_unique_id"`
	} `json:"photo"`
	AvailableReactions []any `json:"available_reactions"`
	MaxReactionCount   int   `json:"max_reaction_count"`
	AccentColorID      int   `json:"accent_color_id"`
}

type Message struct {
	MessageID       int  `json:"message_id"`
	MessageThreadID int  `json:"message_thread_id"`
	Date            int  `json:"date"`
	Chat            Chat `json:"chat"`
}
