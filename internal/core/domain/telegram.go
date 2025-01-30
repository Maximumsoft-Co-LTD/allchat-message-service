package domain

type TelegramWebhook struct {
	Message       Message      `json:"message" bson:"message"`
	EditedMessage Message      `json:"edited_message" bson:"edited_message"`
	MyChatMember  MyChatMember `json:"my_chat_member,omitempty" bson:"my_chat_member,omitempty"`
	UpdateID      int          `json:"update_id" bson:"update_id"`
}

type MyChatMember struct {
	Chat          Chat          `json:"chat" bson:"chat"`
	Date          int           `json:"date" bson:"date"`
	From          TelegramFrom  `json:"from" bson:"from"`
	NewChatMember NewChatMember `json:"new_chat_member" bson:"new_chat_member"`
	OldChatMember OldChatMember `json:"old_chat_member" bson:"old_chat_member"`
}

type Chat struct {
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators" bson:"all_members_are_administrators"`
	ID                          int    `json:"id" bson:"id"`
	Title                       string `json:"title" bson:"title"`
	Type                        string `json:"type" bson:"type"`
}

type NewChatMember struct {
	CanBeEdited         bool         `json:"can_be_edited" bson:"can_be_edited"`
	CanChangeInfo       bool         `json:"can_change_info" bson:"can_change_info"`
	CanDeleteMessages   bool         `json:"can_delete_messages" bson:"can_delete_messages"`
	CanInviteUsers      bool         `json:"can_invite_users" bson:"can_invite_users"`
	CanManageChat       bool         `json:"can_manage_chat" bson:"can_manage_chat"`
	CanManageVideoChats bool         `json:"can_manage_video_chats" bson:"can_manage_video_chats"`
	CanManageVoiceChats bool         `json:"can_manage_voice_chats" bson:"can_manage_voice_chats"`
	CanPinMessages      bool         `json:"can_pin_messages" bson:"can_pin_messages"`
	CanPromoteMembers   bool         `json:"can_promote_members" bson:"can_promote_members"`
	CanRestrictMembers  bool         `json:"can_restrict_members" bson:"can_restrict_members"`
	IsAnonymous         bool         `json:"is_anonymous" bson:"is_anonymous"`
	Status              string       `json:"status" bson:"status"`
	User                TelegramUser `json:"user" bson:"user"`
}

type OldChatMember struct {
	Status string       `json:"status" bson:"status"`
	User   TelegramUser `json:"user" bson:"user"`
}

type EditedMessage struct {
	Chat                TelegramUser                `json:"chat" bson:"chat"`
	From                TelegramFrom                `json:"from" bson:"from"`
	Date                int                         `json:"date" bson:"date"`
	MessageID           int                         `json:"message_id" bson:"message_id"`
	Text                string                      `json:"text,omitempty" bson:"text,omitempty"`
	Caption             string                      `json:"caption,omitempty" bson:"caption,omitempty"`
	ReplyToMessage      ReplyToMessage              `json:"reply_to_message,omitempty" bson:"reply_to_message,omitempty"`
	Sticker             TelegramSticker             `json:"sticker,omitempty" bson:"sticker,omitempty"`
	Photo               []TelegramPhoto             `json:"photo,omitempty" bson:"photo,omitempty"`
	Document            TelegramDocument            `json:"document,omitempty" bson:"document,omitempty"`
	Video               TelegramVideo               `json:"video,omitempty" bson:"video,omitempty"`
	Contact             TelegramContact             `json:"contact,omitempty" bson:"contact,omitempty"`
	Location            TelegramLocation            `json:"location,omitempty" bson:"location,omitempty"`
	Poll                TelegramPoll                `json:"poll,omitempty" bson:"poll,omitempty"`
	Animation           TelegramAnimation           `json:"animation,omitempty" bson:"animation,omitempty"`
	PinnedMessage       PinnedMessage               `json:"pinned_message" bson:"pinned_message"`
	NewChatTitle        string                      `json:"new_chat_title,omitempty" bson:"new_chat_title,omitempty"`
	NewChatParticipant  TelegramNewChatParticipant  `json:"new_chat_participant,omitempty" bson:"new_chat_participant,omitempty"`
	LeftChatParticipant TelegramLeftChatParticipant `json:"left_chat_participant,omitempty" bson:"left_chat_participant,omitempty"`
	NewChatPhoto        []TelegramPhoto             `json:"new_chat_photo,omitempty" bson:"new_chat_photo,omitempty"`
}

type Message struct {
	Chat                TelegramUser                `json:"chat" bson:"chat"`
	From                TelegramFrom                `json:"from" bson:"from"`
	Date                int                         `json:"date" bson:"date"`
	MessageID           int                         `json:"message_id" bson:"message_id"`
	Text                string                      `json:"text,omitempty" bson:"text,omitempty"`
	Caption             string                      `json:"caption,omitempty" bson:"caption,omitempty"`
	ReplyToMessage      ReplyToMessage              `json:"reply_to_message,omitempty" bson:"reply_to_message,omitempty"`
	Sticker             TelegramSticker             `json:"sticker,omitempty" bson:"sticker,omitempty"`
	Photo               []TelegramPhoto             `json:"photo,omitempty" bson:"photo,omitempty"`
	Document            TelegramDocument            `json:"document,omitempty" bson:"document,omitempty"`
	Video               TelegramVideo               `json:"video,omitempty" bson:"video,omitempty"`
	Contact             TelegramContact             `json:"contact,omitempty" bson:"contact,omitempty"`
	Location            TelegramLocation            `json:"location,omitempty" bson:"location,omitempty"`
	Poll                TelegramPoll                `json:"poll,omitempty" bson:"poll,omitempty"`
	Animation           TelegramAnimation           `json:"animation,omitempty" bson:"animation,omitempty"`
	PinnedMessage       PinnedMessage               `json:"pinned_message,omitempty" bson:"pinned_message,omitempty"`
	NewChatTitle        string                      `json:"new_chat_title,omitempty" bson:"new_chat_title,omitempty"`
	NewChatParticipant  TelegramNewChatParticipant  `json:"new_chat_participant,omitempty" bson:"new_chat_participant,omitempty"`
	LeftChatParticipant TelegramLeftChatParticipant `json:"left_chat_participant,omitempty" bson:"left_chat_participant,omitempty"`
	NewChatPhoto        []TelegramPhoto             `json:"new_chat_photo,omitempty" bson:"new_chat_photo,omitempty"`
}

type PinnedMessage struct {
	Chat      TelegramUser `json:"chat" bson:"chat"`
	Date      int          `json:"date" bson:"date"`
	From      TelegramFrom `json:"from" bson:"from"`
	MessageID int          `json:"message_id" bson:"message_id"`
}

type ChatInformation struct {
	Ok     bool                  `json:"ok" bson:"ok"`
	Result ResultChatInformation `json:"result" bson:"result"`
}

type ResultChatInformation struct {
	ID                          int         `json:"id" bson:"id"`
	Title                       string      `json:"title" bson:"title"`
	Type                        string      `json:"type" bson:"type"`
	InviteLink                  string      `json:"invite_link" bson:"invite_link"`
	Permissions                 Permissions `json:"permissions" bson:"permissions"`
	AllMembersAreAdministrators bool        `json:"all_members_are_administrators" bson:"all_members_are_administrators"`
	Photo                       Photo       `json:"photo" bson:"photo"`
}

type Permissions struct {
	CanSendMessages       bool `json:"can_send_messages" bson:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages" bson:"can_send_media_messages"`
	CanSendAudios         bool `json:"can_send_audios" bson:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents" bson:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos" bson:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos" bson:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes" bson:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes" bson:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls" bson:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages" bson:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews" bson:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info" bson:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users" bson:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages" bson:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics" bson:"can_manage_topics"`
}

type Photo struct {
	SmallFileID       string `json:"small_file_id" bson:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id" bson:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id" bson:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id" bson:"big_file_unique_id"`
}

type TelegramUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	ID        int64  `json:"id" bson:"id"`
	Type      string `json:"type" bson:"type"`
	Username  string `json:"username" bson:"username"`
	Title     string `json:"title" bson:"title"`
}

type TelegramFrom struct {
	ID           int64  `json:"id" bson:"id"`
	FirstName    string `json:"first_name" bson:"first_name"`
	IsBot        bool   `json:"is_bot" bson:"is_bot"`
	LanguageCode string `json:"language_code" bson:"language_code"`
	Username     string `json:"username" bson:"username"`
}

type TelegramPhoto struct {
	FileID       string `json:"file_id" bson:"file_id"`
	FileSize     int    `json:"file_size" bson:"file_size"`
	FileUniqueID string `json:"file_unique_id" bson:"file_unique_id"`
	Height       int    `json:"height" bson:"height"`
	Width        int    `json:"width" bson:"width"`
}

type TelegramSticker struct {
	Emoji        string `json:"emoji" bson:"emoji"`
	FileID       string `json:"file_id" bson:"file_id"`
	FileSize     int    `json:"file_size" bson:"file_size"`
	FileUniqueID string `json:"file_unique_id" bson:"file_unique_id"`
	Height       int    `json:"height" bson:"height"`
	IsAnimated   bool   `json:"is_animated" bson:"is_animated"`
	IsVideo      bool   `json:"is_video" bson:"is_video"`
	SetName      string `json:"set_name" bson:"set_name"`
	Thumb        Thumb  `json:"thumb" bson:"thumb"`
	Type         string `json:"type" bson:"type"`
	Width        int    `json:"width" bson:"width"`
}

type Thumb struct {
	FileID       string `json:"file_id" bson:"file_id"`
	FileSize     int    `json:"file_size" bson:"file_size"`
	FileUniqueID string `json:"file_unique_id" bson:"file_unique_id"`
	Height       int    `json:"height" bson:"height"`
	Width        int    `json:"width" bson:"width"`
}

type TelegramContact struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	UserID      int64  `json:"user_id" bson:"user_id"`
}

type TelegramLocation struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type TelegramVideo struct {
	Duration     int       `json:"duration" bson:"duration"`
	FileID       string    `json:"file_id" bson:"file_id"`
	FileName     string    `json:"file_name" bson:"file_name"`
	FileSize     int       `json:"file_size" bson:"file_size"`
	FileUniqueID string    `json:"file_unique_id" bson:"file_unique_id"`
	Height       int       `json:"height" bson:"height"`
	MimeType     string    `json:"mime_type" bson:"mime_type"`
	Thumb        Thumb     `json:"thumb" bson:"thumb"`
	Thumbnail    Thumbnail `json:"thumbnail" bson:"thumbnail"`
	Width        int       `json:"width" bson:"width"`
}

type Thumbnail struct {
	FileID       string `json:"file_id" bson:"file_id"`
	FileSize     int    `json:"file_size" bson:"file_size"`
	FileUniqueID string `json:"file_unique_id" bson:"file_unique_id"`
	Height       int    `json:"height" bson:"height"`
	Width        int    `json:"width" bson:"width"`
}

type TelegramAnimation struct {
	Duration     int       `json:"duration" bson:"duration"`
	FileID       string    `json:"file_id" bson:"file_id"`
	FileName     string    `json:"file_name" bson:"file_name"`
	FileSize     int       `json:"file_size" bson:"file_size"`
	FileUniqueID string    `json:"file_unique_id" bson:"file_unique_id"`
	Height       int       `json:"height" bson:"height"`
	MimeType     string    `json:"mime_type" bson:"mime_type"`
	Thumb        Thumb     `json:"thumb" bson:"thumb"`
	Thumbnail    Thumbnail `json:"thumbnail" bson:"thumbnail"`
	Width        int       `json:"width" bson:"width"`
}

type TelegramDocument struct {
	FileID       string    `json:"file_id" bson:"file_id"`
	FileName     string    `json:"file_name" bson:"file_name"`
	FileSize     int       `json:"file_size" bson:"file_size"`
	FileUniqueID string    `json:"file_unique_id" bson:"file_unique_id"`
	MimeType     string    `json:"mime_type" bson:"mime_type"`
	Thumb        Thumb     `json:"thumb,omitempty" bson:"thumb,omitempty"`
	Thumbnail    Thumbnail `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
}

type TelegramPoll struct {
	AllowsMultipleAnswers bool      `json:"allows_multiple_answers" bson:"allows_multiple_answers"`
	ID                    string    `json:"id" bson:"id"`
	IsAnonymous           bool      `json:"is_anonymous" bson:"is_anonymous"`
	IsClosed              bool      `json:"is_closed" bson:"is_closed"`
	Options               []Options `json:"options" bson:"options"`
	Question              string    `json:"question" bson:"question"`
	TotalVoterCount       int       `json:"total_voter_count" bson:"total_voter_count"`
	Type                  string    `json:"type" bson:"type"`
}

type Options struct {
	Text       string `json:"text" bson:"text"`
	VoterCount int    `json:"voter_count" bson:"voter_count"`
}

type TelegramNewChatParticipant struct {
	FirstName string `json:"first_name" bson:"first_name"`
	ID        int64  `json:"id" bson:"id"`
	IsBot     bool   `json:"is_bot" bson:"is_bot"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type TelegramLeftChatParticipant struct {
	FirstName    string `json:"first_name" bson:"first_name"`
	ID           int64  `json:"id" bson:"id"`
	IsBot        bool   `json:"is_bot" bson:"is_bot"`
	LanguageCode string `json:"language_code" bson:"language_code"`
	LastName     string `json:"last_name" bson:"last_name"`
}

type SetWebhookTelegram struct {
	Token string `json:"token" binding:"required" bson:"token"`
	Url   string `json:"url" binding:"required" bson:"url"`
}

type ProfileTelegram struct {
	Ok     bool          `json:"ok" bson:"ok"`
	Result ResultProfile `json:"result" bson:"result"`
}

type ResultProfile struct {
	ID                      int64  `json:"id" bson:"id"`
	IsBot                   bool   `json:"is_bot" bson:"is_bot"`
	FirstName               string `json:"first_name" bson:"first_name"`
	Username                string `json:"username" bson:"username"`
	CanJoinGroups           bool   `json:"can_join_groups" bson:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages" bson:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries" bson:"supports_inline_queries"`
}

type ImagesProfileTelegram struct {
	Ok     bool                `json:"ok" bson:"ok"`
	Result ResultImagesProfile `json:"result" bson:"result"`
}

type ResultImagesProfile struct {
	TotalCount int               `json:"total_count" bson:"total_count"`
	Photos     [][]TelegramPhoto `json:"photos" bson:"photos"`
}

type FilePathTelegram struct {
	Ok     bool           `json:"ok" bson:"ok"`
	Result ResultFilePath `json:"result" bson:"result"`
}

type ResultFilePath struct {
	FileID       string `json:"file_id" bson:"file_id"`
	FileUniqueID string `json:"file_unique_id" bson:"file_unique_id"`
	FileSize     int    `json:"file_size" bson:"file_size"`
	FilePath     string `json:"file_path" bson:"file_path"`
}

type MessageTelegram struct {
	Type             string `json:"type" bson:"type"`
	Text             string `json:"text" bson:"text"`
	ChatID           string `json:"chat_id" bson:"chat_id"`
	StickerID        string `json:"stickerId" bson:"stickerId"`
	PackageID        string `json:"packageId" bson:"packageId"`
	ReplyToMessageID int    `json:"reply_to_message_id" bson:"reply_to_message_id"`
}

type ReplyToMessage struct {
	Chat      Chat         `json:"chat" bson:"chat"`
	Date      int          `json:"date" bson:"date"`
	From      TelegramFrom `json:"from" bson:"from"`
	MessageID int          `json:"message_id" bson:"message_id"`
	Text      string       `json:"text" bson:"text"`
}

type ErrorMessageTelegram struct {
	Ok          bool   `json:"ok" bson:"ok"`
	ErrorCode   int    `json:"error_code" bson:"error_code"`
	Description string `json:"description" bson:"description"`
}

type ResponseMessageForTelegram struct {
	Ok     bool              `json:"ok" bson:"ok"`
	Result ResultRespMessage `json:"result" bson:"result"`
}

type ResultRespMessage struct {
	MessageID      int            `json:"message_id" bson:"message_id"`
	From           TelegramFrom   `json:"from" bson:"from"`
	Chat           TelegramUser   `json:"chat" bson:"chat"`
	Date           int64          `json:"date" bson:"date"`
	Text           string         `json:"text" bson:"text"`
	ReplyToMessage ReplyToMessage `json:"reply_to_message" bson:"reply_to_message"`
}
