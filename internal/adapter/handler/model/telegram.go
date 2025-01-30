package model

type TelegramWebhook struct {
	Message       Message       `json:"message"`
	EditedMessage EditedMessage `json:"edited_message"`
	MyChatMember  MyChatMember  `json:"my_chat_member,omitempty"`
	UpdateID      int           `json:"update_id"`
}

type MyChatMember struct {
	Chat          Chat          `json:"chat"`
	Date          int           `json:"date"`
	From          TelegramFrom  `json:"from"`
	NewChatMember NewChatMember `json:"new_chat_member"`
	OldChatMember OldChatMember `json:"old_chat_member"`
}

type Chat struct {
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
	ID                          int    `json:"id"`
	Title                       string `json:"title"`
	Type                        string `json:"type"`
}

type NewChatMember struct {
	CanBeEdited         bool         `json:"can_be_edited"`
	CanChangeInfo       bool         `json:"can_change_info"`
	CanDeleteMessages   bool         `json:"can_delete_messages"`
	CanInviteUsers      bool         `json:"can_invite_users"`
	CanManageChat       bool         `json:"can_manage_chat"`
	CanManageVideoChats bool         `json:"can_manage_video_chats"`
	CanManageVoiceChats bool         `json:"can_manage_voice_chats"`
	CanPinMessages      bool         `json:"can_pin_messages"`
	CanPromoteMembers   bool         `json:"can_promote_members"`
	CanRestrictMembers  bool         `json:"can_restrict_members"`
	IsAnonymous         bool         `json:"is_anonymous"`
	Status              string       `json:"status"`
	User                TelegramUser `json:"user"`
}

type OldChatMember struct {
	Status string       `json:"status"`
	User   TelegramUser `json:"user"`
}

type EditedMessage struct {
	Chat           TelegramUser      `json:"chat"`
	From           TelegramFrom      `json:"from"`
	Date           int               `json:"date"`
	MessageID      int               `json:"message_id"`
	Text           string            `json:"text,omitempty"`
	Caption        string            `json:"caption,omitempty"`
	ReplyToMessage ReplyToMessage    `json:"reply_to_message,omitempty"`
	Sticker        TelegramSticker   `json:"sticker,omitempty"`
	Photo          []TelegramPhoto   `json:"photo,omitempty"`
	Document       TelegramDocument  `json:"document,omitempty"`
	Video          TelegramVideo     `json:"video,omitempty"`
	Contact        TelegramContact   `json:"contact,omitempty"`
	Location       TelegramLocation  `json:"location,omitempty"`
	Poll           TelegramPoll      `json:"poll,omitempty"`
	Animation      TelegramAnimation `json:"animation,omitempty"`
	PinnedMessage  PinnedMessage     `json:"pinned_message"`
	// message for chat group ↧↧↧↧↧↧↧
	NewChatTitle        string                      `json:"new_chat_title,omitempty"`
	NewChatParticipant  TelegramNewChatParticipant  `json:"new_chat_participant,omitempty"`
	LeftChatParticipant TelegramLeftChatParticipant `json:"left_chat_participant,omitempty"`
	NewChatPhoto        []TelegramPhoto             `json:"new_chat_photo,omitempty"`
	// message for chat group ↥↥↥↥↥↥↥
}

type Message struct {
	Chat           TelegramUser      `json:"chat"`
	From           TelegramFrom      `json:"from"`
	Date           int               `json:"date"`
	MessageID      int               `json:"message_id"`
	Text           string            `json:"text,omitempty"`
	Caption        string            `json:"caption,omitempty"`
	ReplyToMessage ReplyToMessage    `json:"reply_to_message,omitempty"`
	Sticker        TelegramSticker   `json:"sticker,omitempty"`
	Photo          []TelegramPhoto   `json:"photo,omitempty"`
	Document       TelegramDocument  `json:"document,omitempty"`
	Video          TelegramVideo     `json:"video,omitempty"`
	Contact        TelegramContact   `json:"contact,omitempty"`
	Location       TelegramLocation  `json:"location,omitempty"`
	Poll           TelegramPoll      `json:"poll,omitempty"`
	Animation      TelegramAnimation `json:"animation,omitempty"`
	PinnedMessage  PinnedMessage     `json:"pinned_message,omitempty"`
	// message for chat group ↧↧↧↧↧↧↧
	NewChatTitle        string                      `json:"new_chat_title,omitempty"`
	NewChatParticipant  TelegramNewChatParticipant  `json:"new_chat_participant,omitempty"`
	LeftChatParticipant TelegramLeftChatParticipant `json:"left_chat_participant,omitempty"`
	NewChatPhoto        []TelegramPhoto             `json:"new_chat_photo,omitempty"`
}

type PinnedMessage struct {
	Chat      TelegramUser `json:"chat"`
	Date      int          `json:"date"`
	From      TelegramFrom `json:"from"`
	MessageID int          `json:"message_id"`
}

type ChatInformation struct {
	Ok     bool                  `json:"ok"`
	Result ResultChatInformation `json:"result"`
}

type ResultChatInformation struct {
	ID                          int         `json:"id"`
	Title                       string      `json:"title"`
	Type                        string      `json:"type"`
	InviteLink                  string      `json:"invite_link"`
	Permissions                 Permissions `json:"permissions"`
	AllMembersAreAdministrators bool        `json:"all_members_are_administrators"`
	Photo                       Photo       `json:"photo"`
}

type Permissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

type Photo struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type TelegramUser struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Username  string `json:"username"`
	Title     string `json:"title"`
}

type TelegramFrom struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	Username     string `json:"username"`
}

type TelegramPhoto struct {
	FileID       string `json:"file_id"`
	FileSize     int    `json:"file_size"`
	FileUniqueID string `json:"file_unique_id"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type TelegramSticker struct {
	Emoji        string `json:"emoji"`
	FileID       string `json:"file_id"`
	FileSize     int    `json:"file_size"`
	FileUniqueID string `json:"file_unique_id"`
	Height       int    `json:"height"`
	IsAnimated   bool   `json:"is_animated"`
	IsVideo      bool   `json:"is_video"`
	SetName      string `json:"set_name"`
	Thumb        Thumb  `json:"thumb"`
	Type         string `json:"type"`
	Width        int    `json:"width"`
}

type Thumb struct {
	FileID       string `json:"file_id"`
	FileSize     int    `json:"file_size"`
	FileUniqueID string `json:"file_unique_id"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type TelegramContact struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	UserID      int64  `json:"user_id"`
}

type TelegramLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TelegramVideo struct {
	Duration     int       `json:"duration"`
	FileID       string    `json:"file_id"`
	FileName     string    `json:"file_name"`
	FileSize     int       `json:"file_size"`
	FileUniqueID string    `json:"file_unique_id"`
	Height       int       `json:"height"`
	MimeType     string    `json:"mime_type"`
	Thumb        Thumb     `json:"thumb"`
	Thumbnail    Thumbnail `json:"thumbnail"`
	Width        int       `json:"width"`
}

type Thumbnail struct {
	FileID       string `json:"file_id"`
	FileSize     int    `json:"file_size"`
	FileUniqueID string `json:"file_unique_id"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type TelegramAnimation struct {
	Duration     int       `json:"duration"`
	FileID       string    `json:"file_id"`
	FileName     string    `json:"file_name"`
	FileSize     int       `json:"file_size"`
	FileUniqueID string    `json:"file_unique_id"`
	Height       int       `json:"height"`
	MimeType     string    `json:"mime_type"`
	Thumb        Thumb     `json:"thumb"`
	Thumbnail    Thumbnail `json:"thumbnail"`
	Width        int       `json:"width"`
}

type TelegramDocument struct {
	FileID       string    `json:"file_id"`
	FileName     string    `json:"file_name"`
	FileSize     int       `json:"file_size"`
	FileUniqueID string    `json:"file_unique_id"`
	MimeType     string    `json:"mime_type"`
	Thumb        Thumb     `json:"thumb,omitempty"`
	Thumbnail    Thumbnail `json:"thumbnail,omitempty"`
}

type TelegramPoll struct {
	AllowsMultipleAnswers bool      `json:"allows_multiple_answers"`
	ID                    string    `json:"id"`
	IsAnonymous           bool      `json:"is_anonymous"`
	IsClosed              bool      `json:"is_closed"`
	Options               []Options `json:"options"`
	Question              string    `json:"question"`
	TotalVoterCount       int       `json:"total_voter_count"`
	Type                  string    `json:"type"`
}

type Options struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type TelegramNewChatParticipant struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	LastName  string `json:"last_name"`
}

type TelegramLeftChatParticipant struct {
	FirstName    string `json:"first_name"`
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	LastName     string `json:"last_name"`
}

type SetWebhookTelegram struct {
	Token string `json:"token" binding:"required"`
	Url   string `json:"url" binding:"required"`
}

type ProfileTelegram struct {
	Ok     bool          `json:"ok"`
	Result ResultProfile `json:"result"`
}

type ResultProfile struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	Username                string `json:"username"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type ImagesProfileTelegram struct {
	Ok     bool                `json:"ok"`
	Result ResultImagesProfile `json:"result"`
}

type ResultImagesProfile struct {
	TotalCount int               `json:"total_count"`
	Photos     [][]TelegramPhoto `json:"photos"`
}

type FilePathTelegram struct {
	Ok     bool           `json:"ok"`
	Result ResultFilePath `json:"result"`
}

type ResultFilePath struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

type MessageTelegram struct {
	Type             string `json:"type"`
	Text             string `json:"text"`
	ChatID           string `json:"chat_id"`
	StickerID        string `json:"stickerId"`
	PackageID        string `json:"packageId"`
	ReplyToMessageID int    `json:"reply_to_message_id"`
}

type ReplyToMessage struct {
	Chat      Chat         `json:"chat"`
	Date      int          `json:"date"`
	From      TelegramFrom `json:"from"`
	MessageID int          `json:"message_id"`
	Text      string       `json:"text"`
}

type ErrorMessageTelegram struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type ResponseMessageForTelegram struct {
	Ok     bool              `json:"ok"`
	Result ResultRespMessage `json:"result"`
}

type ResultRespMessage struct {
	MessageID      int            `json:"message_id"`
	From           TelegramFrom   `json:"from"`
	Chat           TelegramUser   `json:"chat"`
	Date           int64          `json:"date"`
	Text           string         `json:"text"`
	ReplyToMessage ReplyToMessage `json:"reply_to_message"`
}
