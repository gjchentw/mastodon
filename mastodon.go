package mastodon

import (
	"github.com/gin-gonic/gin"
	"github.com/ginmills/ginmill"
)

type IMastodon interface {

	// OAuthAuthorize for GET /oauth/authorize
	OAuthAuthorize(c *gin.Context)
	// OAuthObtainToken for POST /oauth/token
	OAuthObtainToken(c *gin.Context)
	// OAuthRevokeToken for POST /oauth/revoke
	OAuthRevokeToken(c *gin.Context)
	// OAuthUserInfo for GET /oauth/userinfo
	OAuthUserInfo(c *gin.Context)
	// OAuthServerConfiguration for GET /.well-known/oauth-authorization-server
	OAuthServerConfiguration(c *gin.Context)

	// CreateAppV1 for POST /api/v1/apps
	CreateAppV1(c *gin.Context)
	// VerifyAppCredentialsV1 for GET /api/v1/apps/verify_credentials
	VerifyAppCredentialsV1(c *gin.Context)

	// EmailResendConfirmation POST /api/v1/emails/confirmations
	EmailResendConfirmation(c *gin.Context)

	// RegisterAccountV1 for POST /api/v1/accounts
	RegisterAccountV1(c *gin.Context)
	// VerifyAccountCredentialsV1 for GET /api/v1/accounts/verify_credentials
	VerifyAccountCredentialsV1(c *gin.Context)
	// UpdateAccountCredentialsV1 for PATCH /api/v1/accounts/verify_credentials
	UpdateAccountCredentialsV1(c *gin.Context)
	// GetAccountV1 for GET /api/v1/accounts/:id
	GetAccountV1(c *gin.Context)
	// GetAccountStatusesV1 for GET /api/v1/accounts/:id/statuses
	GetAccountStatusesV1(c *gin.Context)
	// GetAccountFollowersV1 for GET /api/v1/accounts/:id/followers
	GetAccountFollowersV1(c *gin.Context)
	// GetAccountFollowingV1 for GET /api/v1/accounts/:id/following
	GetAccountFollowingV1(c *gin.Context)
	// GetAccountFeaturedTagsV1 for GET /api/v1/accounts/:id/featured_tags
	GetAccountFeaturedTagsV1(c *gin.Context)
	// GetAccountOwnedListsV1 for GET /api/v1/accounts/:id/lists
	GetAccountOwnedListsV1(c *gin.Context)
	// GetAccountIdentityProofsV1 for GET /api/v1/accounts/:id/identity_proofs
	GetAccountIdentityProofsV1(c *gin.Context)
	// FollowAccountV1 for POST /api/v1/accounts/:id/follow
	FollowAccountV1(c *gin.Context)
	// UnfollowAccountV1 for POST /api/v1/accounts/:id/unfollow
	UnfollowAccountV1(c *gin.Context)
	// BlockAccountV1 for POST /api/v1/accounts/:id/block
	BlockAccountV1(c *gin.Context)
	// UnblockAccountV1 for POST /api/v1/accounts/:id/unblock
	UnblockAccountV1(c *gin.Context)
	// MuteAccountV1 for POST /api/v1/accounts/:id/mute
	MuteAccountV1(c *gin.Context)
	// UnmuteAccountV1 for POST /api/v1/accounts/:id/unmute
	UnmuteAccountV1(c *gin.Context)
	// PinAccountV1 for POST /api/v1/accounts/:id/pin
	PinAccountV1(c *gin.Context)
	// UnpinAccountV1 for POST /api/v1/accounts/:id/unpin
	UnpinAccountV1(c *gin.Context)
	// GetAccountsRelationshipsV1 for GET /api/v1/accounts/relationships
	GetAccountsRelationshipsV1(c *gin.Context)
	// SearchAccountsV1 for GET /api/v1/accounts/search
	SearchAccountsV1(c *gin.Context)

	// GetAccountsV1 for GET /api/v1/accounts
	GetAccountsV1(c *gin.Context)
	// RemoveFollowerV1 for POST /api/v1/accounts/:id/remove_from_followers
	RemoveFollowerV1(c *gin.Context)
	// EndorseAccountV1 for POST /api/v1/accounts/:id/endorse
	EndorseAccountV1(c *gin.Context)
	// UnendorseAccountV1 for POST /api/v1/accounts/:id/unendorse
	UnendorseAccountV1(c *gin.Context)
	// SetAccountNoteV1 for POST /api/v1/accounts/:id/note
	SetAccountNoteV1(c *gin.Context)
	// FindAccountFamiliarFollowersV1 for GET /api/v1/accounts/familiar_followers
	FindAccountFamiliarFollowersV1(c *gin.Context)
	// LookupAccountIDV1 for GET /api/v1/accounts/lookup
	LookupAccountIDV1(c *gin.Context)

	// GetBookmarksV1 for GET /api/v1/bookmarks
	GetBookmarksV1(c *gin.Context)
	// GetFavouritesV1 for GET /api/v1/favourites
	GetFavouritesV1(c *gin.Context)
	// GetMutesV1 for GET /api/v1/mutes
	GetMutesV1(c *gin.Context)
	// GetBlocksV1 for GET /api/v1/blocks
	GetBlocksV1(c *gin.Context)

	// GetDomainBlocksV1 for GET /api/v1/domain_blocks
	GetDomainBlocksV1(c *gin.Context)
	// BlockDomainV1 for POST /api/v1/domain_blocks
	BlockDomainV1(c *gin.Context)
	// UnblockDomainV1 for DELETE /api/v1/domain_blocks
	UnblockDomainV1(c *gin.Context)

	// GetFiltersV1 for GET /api/v1/filters
	GetFiltersV1(c *gin.Context)
	// GetFilterV1 for GET /api/v1/filters/:id
	GetFilterV1(c *gin.Context)
	// CreateFilterV1 for POST /api/v1/filters
	CreateFilterV1(c *gin.Context)
	// UpdateFilterV1 for PUT /api/v1/filters/:id
	UpdateFilterV1(c *gin.Context)
	// RemoveFilterV1 for DELETE /api/v1/filters/:id
	RemoveFilterV1(c *gin.Context)
	// FileReportV1 for POST /api/v1/reports
	FileReportV1(c *gin.Context)
	// GetFollowRequestsV1 for GET /api/v1/follow_requests
	GetFollowRequestsV1(c *gin.Context)
	// AcceptFollowV1 for POST /api/v1/follow_requests/:id/authorize
	AcceptFollowV1(c *gin.Context)
	// RejectFollowV1 for POST /api/v1/follow_requests/:id/reject
	RejectFollowV1(c *gin.Context)
	// GetEndorsementsV1 for GET /api/v1/endorsements
	GetEndorsementsV1(c *gin.Context)
	// GetFeaturedTagsV1 for GET /api/v1/featured_tags
	GetFeaturedTagsV1(c *gin.Context)
	// GetFeaturedTagSuggestionsV1 GET /api/v1/featured_tags/suggestions
	GetFeaturedTagSuggestionsV1(c *gin.Context)
	// FeatureTagV1 for POST /api/v1/featured_tags
	FeatureTagV1(c *gin.Context)
	// UnfeatureTagV1 for DELETE /api/v1/featured_tags/:id
	UnfeatureTagV1(c *gin.Context)
	// GetFollowedTagsV1 for GET /api/v1/followed_tags
	GetFollowedTagsV1(c *gin.Context)
	// GetTagV1 for GET /api/v1/tags/:name
	GetTagV1(c *gin.Context)
	// FollowTagV1 for POST /api/v1/tags/:name/follow
	FollowTagV1(c *gin.Context)
	// UnfollowTagV1 for POST /api/v1/tags/:name/unfollow
	UnfollowTagV1(c *gin.Context)
	// SetFeatureTagV1 for POST /api/v1/tags/:name/feature
	SetFeatureTagV1(c *gin.Context)
	// SetUnfeatureTagV1 for POST /api/v1/tags/:name/unfeature
	SetUnfeatureTagV1(c *gin.Context)
	// GetPreferencesV1 for GET /api/v1/preferences
	GetPreferencesV1(c *gin.Context)
	// GetFollowSuggestionsV1 for GET /api/v1/suggestions
	GetFollowSuggestionsV1(c *gin.Context)
	// for DELETE /api/v1/suggestions/:id
	RemoveFollowSuggestionV1(c *gin.Context)
	// PubishStatusV1 for POST /api/v1/statuses
	PubishStatusV1(c *gin.Context)
	// GetSatatusV1 for GET /api/v1/statuses/:id
	GetSatatusV1(c *gin.Context)
	// DeleteStatusV1 for DELETE /api/v1/statuses/:id
	DeleteStatusV1(c *gin.Context)
	// GetStatusContextV1 for GET /api/v1/statuses/:id/context
	GetStatusContextV1(c *gin.Context)
	// GetRebloggedByV1 for GET /api/v1/statuses/:id/reblogged_by
	GetRebloggedByV1(c *gin.Context)
	// GetFavouritedByV1 for GET /api/v1/statuses/:id/favourited_by
	GetFavouritedByV1(c *gin.Context)
	// FavouriteStatusV1 for POST /api/v1/statuses/:id/favourite
	FavouriteStatusV1(c *gin.Context)
	// UnfavouriteStatusV1 for POST /api/v1/statuses/:id/unfavourite
	UnfavouriteStatusV1(c *gin.Context)
	// RebloggedStatusV1 for POST /api/v1/statuses/:id/reblog
	RebloggedStatusV1(c *gin.Context)
	// UnrebloggedStatusV1 for POST /api/v1/statuses/:id/unreblog
	UnrebloggedStatusV1(c *gin.Context)
	// BookmarkStatusV1 for POST /api/v1/statuses/:id/bookmark
	BookmarkStatusV1(c *gin.Context)
	// UnbookmarkStatusV1 for POST /api/v1/statuses/:id/unbookmark
	UnbookmarkStatusV1(c *gin.Context)
	// MuteStatusConversationV1 for POST /api/v1/statuses/:id/mute
	MuteStatusConversationV1(c *gin.Context)
	// UnmuteStatusConversationV1 for POST /api/v1/statuses/:id/unmute
	UnmuteStatusConversationV1(c *gin.Context)
	// PinSatatusV1 for POST /api/v1/statuses/:id/pin
	PinSatatusV1(c *gin.Context)
	// UnpinSatatusV1 for POST /api/v1/statuses/:id/unpin
	UnpinSatatusV1(c *gin.Context)
	// GetStatusCardV1 for GET /api/v1/statuses/:id/card")
	GetStatusCardV1(c *gin.Context)
	// CreateAttachmentV1 for POST /api/v1/media
	CreateAttachmentV1(c *gin.Context)
	// GetAttachmentV1 for GET /api/v1/media/:id
	GetAttachmentV1(c *gin.Context)
	// UpdateAttachmentV1 for PUT /api/v1/media/:id
	UpdateAttachmentV1(c *gin.Context)
	// GetPollV1 for GET /api/v1/polls/:id
	GetPollV1(c *gin.Context)
	// VotePollV1 for DELETE /api/v1/polls/:id/votes
	VotePollV1(c *gin.Context)
	// GetScheduledStatusesV1 for GET /api/v1/scheduled_statuses
	GetScheduledStatusesV1(c *gin.Context)
	// GetScheduledStatusV1 for GET /api/v1/scheduled_statuses/:id
	GetScheduledStatusV1(c *gin.Context)
	// CreateScheduledStatusV1 for PUT /api/v1/scheduled_statuses/:id
	CreateScheduledStatusV1(c *gin.Context)
	// CancellScheduledStatusV1 for DELETE /api/v1/scheduled_statuses/:id
	CancellScheduledStatusV1(c *gin.Context)
	// PublicTimelineV1 for GET /api/v1/timelines/public
	PublicTimelineV1(c *gin.Context)
	// HashtagTimelineV1 for GET /api/v1/timelines/tag/:hashtag
	HashtagTimelineV1(c *gin.Context)
	// HomeTimelineV1 for GET /api/v1/timelines/home
	HomeTimelineV1(c *gin.Context)
	// ListTimelineV1 for GET /api/v1/timelines/list/:list_id
	ListTimelineV1(c *gin.Context)
	// DirectTimelineV1 for GET /api/v1/timelines/direct
	DirectTimelineV1(c *gin.Context)
	// GetConversationsV1 for GET /api/v1/conversations
	GetConversationsV1(c *gin.Context)
	// RemoveConversationV1 for DELETE /api/v1/conversations/:id
	RemoveConversationV1(c *gin.Context)
	// MarkConversationAsReadV1 for POST /api/v1/conversations/:id/read
	MarkConversationAsReadV1(c *gin.Context)
	// GetListsV1 for GET /api/v1/lists
	GetListsV1(c *gin.Context)
	// GetListV1 for GET /api/v1/lists/:id
	GetListV1(c *gin.Context)
	// CreateListV1 for POST /api/v1/lists
	CreateListV1(c *gin.Context)
	// UpdateListV1 for PUT /api/v1/lists/:id
	UpdateListV1(c *gin.Context)
	// DeleteListV1 for DELETE /api/v1/lists/:id
	DeleteListV1(c *gin.Context)
	// GetListAccountV1 for GET /api/v1/lists/:id/accounts
	GetListAccountV1(c *gin.Context)
	// AddListAccountV1 for POST /api/v1/lists/:id/accounts
	AddListAccountV1(c *gin.Context)
	// RemoveListAccountV1 for DELETE /api/v1/lists/:id/accounts
	RemoveListAccountV1(c *gin.Context)
	// GetMarkersV1 for GET /api/v1/markers
	GetMarkersV1(c *gin.Context)
	// CreateMarkerV1 for POST /api/v1/markers
	CreateMarkerV1(c *gin.Context)
	// StreamingForHealthCheckV1 for GET /api/v1/streaming/health
	StreamingForHealthCheckV1(c *gin.Context)
	// StreamingForUserV1 for GET /api/v1/streaming/user
	StreamingForUserV1(c *gin.Context)
	// StreamingForPublicV1 for GET /api/v1/streaming/public
	StreamingForPublicV1(c *gin.Context)
	// StreamingForLocalV1 for GET /api/v1/streaming/public/local
	StreamingForLocalV1(c *gin.Context)
	// StreamingForPublicHashtagV1 for GET /api/v1/streaming/hashtag
	StreamingForPublicHashtagV1(c *gin.Context)
	// StreamingForLocalHashtagV1 for GET /api/v1/streaming/hashtag/local
	StreamingForLocalHashtagV1(c *gin.Context)
	// StreamingForListV1 for GET /api/v1/streaming/list
	StreamingForListV1(c *gin.Context)
	// StreamingForDirecttV1 for GET /api/v1/streaming/direct
	StreamingForDirectV1(c *gin.Context)
	// StreamingV1 for GET /api/v1/streaming
	StreamingV1(c *gin.Context)
	// GetNotificationsV1 for GET /api/v1/notifications
	GetNotificationsV1(c *gin.Context)
	// GetNotificationV1 for GET /api/v1/notifications/:id
	GetNotificationV1(c *gin.Context)
	// DissmissNotificationsV1 for POST /api/v1/notifications/clear
	DissmissNotificationsV1(c *gin.Context)
	// DissmissNotificationV1 for POST /api/v1/notifications/:id/dismiss
	DissmissNotificationV1(c *gin.Context)
	// DeprecatedDissmissNotificationV1 for POST /api/v1/notifications/dismiss
	DeprecatedDissmissNotificationV1(c *gin.Context)
	// SubscribeNotificationsV1 for POST /api/v1/push/subscription
	SubscribeNotificationsV1(c *gin.Context)
	// GetSubscriptionV1 for GET /api/v1/push/subscription
	GetSubscriptionV1(c *gin.Context)
	// UpdateSubscriptionV1 for PUT /api/v1/push/subscription
	UpdateSubscriptionV1(c *gin.Context)
	// UnsubscribeNotificationsV1 for DELETE /api/v1/push/subscription
	UnsubscribeNotificationsV1(c *gin.Context)
	// SearchV1 for GET /api/v1/search
	SearchV1(c *gin.Context)
	// SearchV2 for GET /api/v2/search
	SearchV2(c *gin.Context)
	// GetInstanceV1 for GET /api/v1/instance
	GetInstanceV1(c *gin.Context)
	// GetPeersV1 for GET /api/v1/instance/peers
	GetPeersV1(c *gin.Context)
	// for GET /api/v1/instance/activity
	GetInstanceActivityV1(c *gin.Context)
	// GetInstanceTrendsV1 for GET /api/v1/trends
	TrendsV1(c *gin.Context)
	// DirectoryV1 for GET /api/v1/directory
	DirectoryV1(c *gin.Context)
	// CustomEmojisV1 for GET /api/v1/custom_emojis
	CustomEmojisV1(c *gin.Context)
	// AdminAccountsV1 for GET /api/v1/admin/accounts
	AdminAccountsV1(c *gin.Context)
	// AdminAccountV1 for GET /api/v1/admin/accounts/:id
	AdminAccountV1(c *gin.Context)
	// AdminAccountActionV1 for POST /api/v1/admin/accounts/:id/action
	AdminAccountActionV1(c *gin.Context)
	// AdminAccountApproveV1 for POST /api/v1/admin/accounts/:id/approve
	AdminAccountApproveV1(c *gin.Context)
	// AdminAccountRejectV1 for POST /api/v1/admin/accounts/:id/reject
	AdminAccountRejectV1(c *gin.Context)
	// AdminAccountEnableV1 for POST /api/v1/admin/accounts/:id/enable
	AdminAccountEnableV1(c *gin.Context)
	// AdminAccountUnsilenceV1 for POST /api/v1/admin/accounts/:id/unsilence
	AdminAccountUnsilenceV1(c *gin.Context)
	// AdminAccountUnsuspendV1 for POST /api/v1/admin/accounts/:id/unsuspend
	AdminAccountUnsuspendV1(c *gin.Context)
	// AdminReportsV1 for GET /api/v1/admin/reports
	AdminReportsV1(c *gin.Context)
	// AdminReportV1 for GET /api/v1/admin/reports/:id
	AdminReportV1(c *gin.Context)
	// AdminReportAssignV1 for POST /api/v1/admin/reports/:id/assign_to_self
	AdminReportAssignV1(c *gin.Context)
	// AdminReportUnassignV1 for POST /api/v1/admin/reports/:id/unassign
	AdminReportUnassignV1(c *gin.Context)
	// AdminReportResolveV1 for POST /api/v1/admin/reports/:id/resolve
	AdminReportResolveV1(c *gin.Context)
	// AdminReportReopenV1 for POST /api/v1/admin/reports/:id/reopen
	AdminReportReopenV1(c *gin.Context)
	// GetAnnouncementsV1 for GET /api/v1/announcements
	GetAnnouncementsV1(c *gin.Context)
	// DismissAnnouncementsV1 for POST /api/v1/announcements/:id/dismiss
	DismissAnnouncementsV1(c *gin.Context)
	// ReactAnnouncementsV1 for PUT /api/v1/announcements/:id/reactions/:name
	ReactAnnouncementsV1(c *gin.Context)
	// UndoReactAnnouncementsV1 for DELETE /api/v1/announcements/:id/reactions/:name
	UndoReactAnnouncementsV1(c *gin.Context)
	// ApiProofs for GET /api/proofs
	ApiProofs(c *gin.Context)
	// ApiOEmbed for GET /api/oembed
	ApiOEmbed(c *gin.Context)
}

// mastodon features
func Features(m IMastodon) (features *ginmill.Features) {
	r := gin.New()

	oauth := r.Group("/oauth")

	oauth.GET("/authorize", m.OAuthAuthorize)
	oauth.POST("/token", m.OAuthObtainToken)
	oauth.POST("/revoke", m.OAuthRevokeToken)
	oauth.GET("/userinfo", m.OAuthUserInfo)

	r.GET("/.well-known/oauth-authorization-server", m.OAuthServerConfiguration)

	apiV1 := r.Group("/api/v1")
	apiV2 := r.Group("/api/v2")

	apiV1.POST("/apps", m.CreateAppV1)
	apiV1.GET("/apps/verify_credentials", m.VerifyAppCredentialsV1)

	apiV1.POST("/emails/confirmations", m.EmailResendConfirmation)

	apiV1.POST("/accounts", m.RegisterAccountV1)
	apiV1.GET("/accounts/verify_credentials", m.VerifyAccountCredentialsV1)
	apiV1.PATCH("/accounts/update_credentials", m.UpdateAccountCredentialsV1)
	apiV1.GET("/accounts/:id", m.GetAccountV1)
	apiV1.GET("/accounts/:id/statuses", m.GetAccountStatusesV1)
	apiV1.GET("/accounts/:id/followers", m.GetAccountFollowersV1)
	apiV1.GET("/accounts/:id/following", m.GetAccountFollowingV1)
	apiV1.GET("/accounts/:id/featured_tags", m.GetAccountFeaturedTagsV1)
	apiV1.GET("/accounts/:id/lists", m.GetAccountOwnedListsV1)
	apiV1.GET("/accounts/:id/identity_proofs", m.GetAccountIdentityProofsV1)
	apiV1.POST("/accounts/:id/follow", m.FollowAccountV1)
	apiV1.POST("/accounts/:id/unfollow", m.UnfollowAccountV1)
	apiV1.POST("/accounts/:id/block", m.BlockAccountV1)
	apiV1.POST("/accounts/:id/unblock", m.UnblockAccountV1)
	apiV1.POST("/accounts/:id/mute", m.MuteAccountV1)
	apiV1.POST("/accounts/:id/unmute", m.UnmuteAccountV1)
	apiV1.POST("/accounts/:id/pin", m.PinAccountV1)
	apiV1.POST("/accounts/:id/unpin", m.UnpinAccountV1)
	apiV1.GET("/accounts/relationships", m.GetAccountsRelationshipsV1)
	apiV1.GET("/accounts/:id/search", m.SearchAccountsV1)
	apiV1.GET("/accounts", m.GetAccountsV1)
	apiV1.POST("/accounts/:id/remove_from_followers", m.RemoveFollowerV1)
	apiV1.POST("/accounts/:id/endorse", m.EndorseAccountV1)
	apiV1.POST("/accounts/:id/unendorse", m.UnendorseAccountV1)
	apiV1.POST("/accounts/:id/note", m.SetAccountNoteV1)
	apiV1.GET("/accounts/familiar_followers", m.FindAccountFamiliarFollowersV1)
	apiV1.GET("/accounts/lookup", m.LookupAccountIDV1)

	apiV1.GET("/bookmarks", m.GetBookmarksV1)
	apiV1.GET("/favourites", m.GetFavouritesV1)
	apiV1.GET("/mutes", m.GetMutesV1)
	apiV1.GET("/blocks", m.GetBlocksV1)

	apiV1.GET("/domain_blocks", m.GetDomainBlocksV1)
	apiV1.POST("/domain_blocks", m.BlockDomainV1)
	apiV1.DELETE("/domain_blocks", m.UnblockDomainV1)

	apiV1.GET("/filters", m.GetFiltersV1)
	apiV1.GET("/filters/:id", m.GetFilterV1)
	apiV1.POST("/filters", m.CreateFilterV1)
	apiV1.PUT("/filters/:id", m.UpdateFilterV1)
	apiV1.DELETE("/filters/:id", m.RemoveFilterV1)

	apiV1.POST("/reports", m.FileReportV1)

	apiV1.GET("/follow_requests", m.GetFollowRequestsV1)
	apiV1.POST("/follow_requests/:id/authorize", m.AcceptFollowV1)
	apiV1.POST("/follow_requests/:id/reject", m.RejectFollowV1)

	apiV1.GET("/endorsements", m.GetEndorsementsV1)

	apiV1.GET("/featured_tags", m.GetFeaturedTagsV1)
	apiV1.GET("/featured_tags/suggestions", m.GetFeaturedTagSuggestionsV1)
	apiV1.POST("/featured_tags", m.FeatureTagV1)
	apiV1.DELETE("/featured_tags/:id", m.UnfeatureTagV1)
	apiV1.GET("/followed_tags", m.GetFollowedTagsV1)

	apiV1.GET("/tags/:name", m.GetTagV1)
	apiV1.POST("/tags/:name/follow", m.FollowTagV1)
	apiV1.POST("/tags/:name/unfollow", m.UnfollowTagV1)
	apiV1.POST("/tags/:name/feature", m.SetFeatureTagV1)
	apiV1.POST("/tags/:name/unfeature", m.SetUnfeatureTagV1)

	apiV1.GET("/preferences", m.GetPreferencesV1)

	apiV1.GET("/suggestions", m.GetFollowSuggestionsV1)
	apiV1.DELETE("/suggestions/:id", m.RemoveFollowSuggestionV1)

	apiV1.POST("/statuses", m.PubishStatusV1)
	apiV1.GET("/statuses/:id", m.GetSatatusV1)
	apiV1.DELETE("/statuses/:id", m.DeleteStatusV1)
	apiV1.GET("/statuses/:id/context", m.GetStatusContextV1)
	apiV1.GET("/statuses/:id/reblogged_by", m.GetRebloggedByV1)
	apiV1.GET("/statuses/:id/favourited_by", m.GetFavouritedByV1)
	apiV1.POST("/statuses/:id/favourite", m.FavouriteStatusV1)
	apiV1.POST("/statuses/:id/unfavourite", m.UnfavouriteStatusV1)
	apiV1.POST("/statuses/:id/reblog", m.RebloggedStatusV1)
	apiV1.POST("/statuses/:id/unreblog", m.UnrebloggedStatusV1)
	apiV1.POST("/statuses/:id/bookmark", m.BookmarkStatusV1)
	apiV1.POST("/statuses/:id/unbookmark", m.UnbookmarkStatusV1)
	apiV1.POST("/statuses/:id/mute", m.MuteStatusConversationV1)
	apiV1.POST("/statuses/:id/unmute", m.UnmuteStatusConversationV1)
	apiV1.POST("/statuses/:id/pin", m.PinSatatusV1)
	apiV1.POST("/statuses/:id/unpin", m.UnpinSatatusV1)
	apiV1.GET("/statuses/:id/card", m.GetStatusCardV1)

	apiV1.POST("/media", m.CreateAttachmentV1)
	apiV1.GET("/media/:id", m.GetAttachmentV1)
	apiV1.PUT("/media/:id", m.UpdateAttachmentV1)

	apiV1.GET("/polls/:id", m.GetPollV1)
	apiV1.DELETE("/polls/:id/votes", m.VotePollV1)

	apiV1.GET("/scheduled_statuses", m.GetScheduledStatusesV1)
	apiV1.GET("/scheduled_statuses/:id", m.GetScheduledStatusV1)
	apiV1.PUT("/scheduled_statuses/:id", m.CreateScheduledStatusV1)
	apiV1.DELETE("/scheduled_statuses/:id", m.CancellScheduledStatusV1)

	apiV1.GET("/timelines/public", m.PublicTimelineV1)
	apiV1.GET("/timelines/tag/:hashtag", m.HashtagTimelineV1)
	apiV1.GET("/timelines/home", m.HomeTimelineV1)
	apiV1.GET("/timelines/list/:list_id", m.ListTimelineV1)
	apiV1.GET("/timelines/direct", m.DirectTimelineV1)

	apiV1.GET("/conversations", m.GetConversationsV1)
	apiV1.DELETE("/conversations/:id", m.RemoveConversationV1)
	apiV1.POST("/conversations/:id/read", m.MarkConversationAsReadV1)

	apiV1.GET("/lists", m.GetListsV1)
	apiV1.GET("/lists/:id", m.GetListV1)
	apiV1.POST("/lists", m.CreateListV1)
	apiV1.PUT("/lists/:id", m.UpdateListV1)
	apiV1.DELETE("/lists/:id", m.DeleteListV1)
	apiV1.GET("/lists/:id/accounts", m.GetListAccountV1)
	apiV1.POST("/lists/:id/accounts", m.AddListAccountV1)
	apiV1.DELETE("/lists/:id/accounts", m.RemoveListAccountV1)

	apiV1.GET("/markers", m.GetMarkersV1)
	apiV1.POST("/markers", m.CreateMarkerV1)

	apiV1.GET("/streaming/health", m.StreamingForHealthCheckV1)
	apiV1.GET("/streaming/user", m.StreamingForUserV1)
	apiV1.GET("/streaming/public", m.StreamingForPublicV1)
	apiV1.GET("/streaming/public/local", m.StreamingForLocalV1)
	apiV1.GET("/streaming/hashtag", m.StreamingForPublicHashtagV1)
	apiV1.GET("/streaming/hashtag/local", m.StreamingForLocalHashtagV1)
	apiV1.GET("/streaming/list", m.StreamingForListV1)
	apiV1.GET("/streaming/direct", m.StreamingForDirectV1)

	apiV1.GET("/streaming", m.StreamingV1)

	apiV1.GET("/notifications", m.GetNotificationsV1)
	apiV1.GET("/notifications/:id", m.GetNotificationV1)
	apiV1.POST("/notifications/clear", m.DissmissNotificationsV1)
	apiV1.POST("/notifications/:id/dismiss", m.DissmissNotificationV1)
	apiV1.POST("/notifications/dismiss", m.DeprecatedDissmissNotificationV1)

	apiV1.POST("/push/subscription", m.SubscribeNotificationsV1)
	apiV1.GET("/push/subscription", m.GetSubscriptionV1)
	apiV1.PUT("/push/subscription", m.UpdateSubscriptionV1)
	apiV1.DELETE("/push/subscription", m.UnsubscribeNotificationsV1)

	apiV1.GET("/search", m.SearchV1)
	apiV2.GET("/search", m.SearchV2)

	apiV1.GET("/instance", m.GetInstanceV1)
	apiV1.GET("/instance/peers", m.GetPeersV1)
	apiV1.GET("/instance/activity", m.GetInstanceActivityV1)

	apiV1.GET("/trends", m.TrendsV1)
	apiV1.GET("/directory", m.DirectoryV1)
	apiV1.GET("/custom_emojis", m.CustomEmojisV1)

	apiV1.GET("/admin/accounts", m.AdminAccountsV1)
	apiV1.GET("/admin/accounts/:id", m.AdminAccountV1)
	apiV1.POST("/admin/accounts/:id/action", m.AdminAccountActionV1)
	apiV1.POST("/admin/accounts/:id/approve", m.AdminAccountApproveV1)
	apiV1.POST("/admin/accounts/:id/reject", m.AdminAccountRejectV1)
	apiV1.POST("/admin/accounts/:id/enable", m.AdminAccountEnableV1)
	apiV1.POST("/admin/accounts/:id/unsilence", m.AdminAccountUnsilenceV1)
	apiV1.POST("/admin/accounts/:id/unsuspend", m.AdminAccountUnsuspendV1)

	apiV1.GET("/admin/reports", m.AdminReportsV1)
	apiV1.GET("/admin/reports/:id", m.AdminReportV1)
	apiV1.POST("/admin/reports/:id/assign_to_self", m.AdminReportAssignV1)
	apiV1.POST("/admin/reports/:id/unassign", m.AdminReportUnassignV1)
	apiV1.POST("/admin/reports/:id/resolve", m.AdminReportResolveV1)
	apiV1.POST("/admin/reports/:id/reopen", m.AdminReportReopenV1)

	apiV1.GET("/announcements", m.GetAnnouncementsV1)
	apiV1.POST("/announcements/:id/dismiss", m.DismissAnnouncementsV1)
	apiV1.PUT("/announcements/:id/reactions/:name", m.ReactAnnouncementsV1)
	apiV1.DELETE("/announcements/:id/reactions/:name", m.UndoReactAnnouncementsV1)

	r.GET("/api/proofs", m.ApiProofs)
	r.GET("/api/oembed", m.ApiOEmbed)

	features = ginmill.NewFeatures(r.Routes())

	return features
}
