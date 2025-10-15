package mastodon

import (
	"testing"

	"github.com/gin-gonic/gin"
)

// mock IMastodon 實作，所有 handler 皆為空

type mockMastodon struct{}

func (m *mockMastodon) OAuthAuthorize(c *gin.Context)                   {}
func (m *mockMastodon) OAuthObtainToken(c *gin.Context)                 {}
func (m *mockMastodon) OAuthRevokeToken(c *gin.Context)                 {}
func (m *mockMastodon) OAuthUserInfo(c *gin.Context)                    {}
func (m *mockMastodon) OAuthServerConfiguration(c *gin.Context)         {}
func (m *mockMastodon) CreateAppV1(c *gin.Context)                      {}
func (m *mockMastodon) VerifyAppCredentialsV1(c *gin.Context)           {}
func (m *mockMastodon) EmailResendConfirmation(c *gin.Context)          {}
func (m *mockMastodon) RegisterAccountV1(c *gin.Context)                {}
func (m *mockMastodon) VerifyAccountCredentialsV1(c *gin.Context)       {}
func (m *mockMastodon) UpdateAccountCredentialsV1(c *gin.Context)       {}
func (m *mockMastodon) GetAccountV1(c *gin.Context)                     {}
func (m *mockMastodon) GetAccountStatusesV1(c *gin.Context)             {}
func (m *mockMastodon) GetAccountFollowersV1(c *gin.Context)            {}
func (m *mockMastodon) GetAccountFollowingV1(c *gin.Context)            {}
func (m *mockMastodon) GetAccountFeaturedTagsV1(c *gin.Context)         {}
func (m *mockMastodon) GetAccountOwnedListsV1(c *gin.Context)           {}
func (m *mockMastodon) GetAccountIdentityProofsV1(c *gin.Context)       {}
func (m *mockMastodon) FollowAccountV1(c *gin.Context)                  {}
func (m *mockMastodon) UnfollowAccountV1(c *gin.Context)                {}
func (m *mockMastodon) BlockAccountV1(c *gin.Context)                   {}
func (m *mockMastodon) UnblockAccountV1(c *gin.Context)                 {}
func (m *mockMastodon) MuteAccountV1(c *gin.Context)                    {}
func (m *mockMastodon) UnmuteAccountV1(c *gin.Context)                  {}
func (m *mockMastodon) PinAccountV1(c *gin.Context)                     {}
func (m *mockMastodon) UnpinAccountV1(c *gin.Context)                   {}
func (m *mockMastodon) NoteAccountV1(c *gin.Context)                    {}
func (m *mockMastodon) GetAccountsRelationshipsV1(c *gin.Context)       {}
func (m *mockMastodon) SearchAccountsV1(c *gin.Context)                 {}
func (m *mockMastodon) GetAccountsV1(c *gin.Context)                    {}
func (m *mockMastodon) GetAccountListsV1(c *gin.Context)                {}
func (m *mockMastodon) RemoveFollowerV1(c *gin.Context)                 {}
func (m *mockMastodon) EndorseAccountV1(c *gin.Context)                 {}
func (m *mockMastodon) UnendorseAccountV1(c *gin.Context)               {}
func (m *mockMastodon) SetAccountNoteV1(c *gin.Context)                 {}
func (m *mockMastodon) CheckAccountRelationshipsV1(c *gin.Context)      {}
func (m *mockMastodon) FindAccountFamiliarFollowersV1(c *gin.Context)   {}
func (m *mockMastodon) LookupAccountIDV1(c *gin.Context)                {}
func (m *mockMastodon) GetBookmarksV1(c *gin.Context)                   {}
func (m *mockMastodon) GetFavouritesV1(c *gin.Context)                  {}
func (m *mockMastodon) GetMutesV1(c *gin.Context)                       {}
func (m *mockMastodon) GetBlocksV1(c *gin.Context)                      {}
func (m *mockMastodon) GetDomainBlocksV1(c *gin.Context)                {}
func (m *mockMastodon) BlockDomainV1(c *gin.Context)                    {}
func (m *mockMastodon) UnblockDomainV1(c *gin.Context)                  {}
func (m *mockMastodon) GetFiltersV1(c *gin.Context)                     {}
func (m *mockMastodon) GetFilterV1(c *gin.Context)                      {}
func (m *mockMastodon) CreateFilterV1(c *gin.Context)                   {}
func (m *mockMastodon) UpdateFilterV1(c *gin.Context)                   {}
func (m *mockMastodon) RemoveFilterV1(c *gin.Context)                   {}
func (m *mockMastodon) FileReportV1(c *gin.Context)                     {}
func (m *mockMastodon) GetFollowRequestsV1(c *gin.Context)              {}
func (m *mockMastodon) AcceptFollowV1(c *gin.Context)                   {}
func (m *mockMastodon) RejectFollowV1(c *gin.Context)                   {}
func (m *mockMastodon) GetEndorsementsV1(c *gin.Context)                {}
func (m *mockMastodon) GetFeaturedTagsV1(c *gin.Context)                {}
func (m *mockMastodon) GetFeaturedTagSuggestionsV1(c *gin.Context)      {}
func (m *mockMastodon) FeatureTagV1(c *gin.Context)                     {}
func (m *mockMastodon) UnfeatureTagV1(c *gin.Context)                   {}
func (m *mockMastodon) GetTagSuggestionsV1(c *gin.Context)              {}
func (m *mockMastodon) GetFollowedTagsV1(c *gin.Context)                {}
func (m *mockMastodon) GetTagV1(c *gin.Context)                         {}
func (m *mockMastodon) FollowTagV1(c *gin.Context)                      {}
func (m *mockMastodon) UnfollowTagV1(c *gin.Context)                    {}
func (m *mockMastodon) SetFeatureTagV1(c *gin.Context)                  {}
func (m *mockMastodon) SetUnfeatureTagV1(c *gin.Context)                {}
func (m *mockMastodon) GetPreferencesV1(c *gin.Context)                 {}
func (m *mockMastodon) GetFollowSuggestionsV1(c *gin.Context)           {}
func (m *mockMastodon) RemoveFollowSuggestionV1(c *gin.Context)         {}
func (m *mockMastodon) PubishStatusV1(c *gin.Context)                   {}
func (m *mockMastodon) GetSatatusV1(c *gin.Context)                     {}
func (m *mockMastodon) DeleteStatusV1(c *gin.Context)                   {}
func (m *mockMastodon) GetStatusContextV1(c *gin.Context)               {}
func (m *mockMastodon) GetRebloggedByV1(c *gin.Context)                 {}
func (m *mockMastodon) GetFavouritedByV1(c *gin.Context)                {}
func (m *mockMastodon) FavouriteStatusV1(c *gin.Context)                {}
func (m *mockMastodon) UnfavouriteStatusV1(c *gin.Context)              {}
func (m *mockMastodon) RebloggedStatusV1(c *gin.Context)                {}
func (m *mockMastodon) UnrebloggedStatusV1(c *gin.Context)              {}
func (m *mockMastodon) BookmarkStatusV1(c *gin.Context)                 {}
func (m *mockMastodon) UnbookmarkStatusV1(c *gin.Context)               {}
func (m *mockMastodon) MuteStatusConversationV1(c *gin.Context)         {}
func (m *mockMastodon) UnmuteStatusConversationV1(c *gin.Context)       {}
func (m *mockMastodon) PinSatatusV1(c *gin.Context)                     {}
func (m *mockMastodon) UnpinSatatusV1(c *gin.Context)                   {}
func (m *mockMastodon) GetStatusCardV1(c *gin.Context)                  {}
func (m *mockMastodon) CreateAttachmentV1(c *gin.Context)               {}
func (m *mockMastodon) GetAttachmentV1(c *gin.Context)                  {}
func (m *mockMastodon) UpdateAttachmentV1(c *gin.Context)               {}
func (m *mockMastodon) GetPollV1(c *gin.Context)                        {}
func (m *mockMastodon) VotePollV1(c *gin.Context)                       {}
func (m *mockMastodon) GetScheduledStatusesV1(c *gin.Context)           {}
func (m *mockMastodon) GetScheduledStatusV1(c *gin.Context)             {}
func (m *mockMastodon) CreateScheduledStatusV1(c *gin.Context)          {}
func (m *mockMastodon) CancellScheduledStatusV1(c *gin.Context)         {}
func (m *mockMastodon) PublicTimelineV1(c *gin.Context)                 {}
func (m *mockMastodon) HashtagTimelineV1(c *gin.Context)                {}
func (m *mockMastodon) HomeTimelineV1(c *gin.Context)                   {}
func (m *mockMastodon) ListTimelineV1(c *gin.Context)                   {}
func (m *mockMastodon) DirectTimelineV1(c *gin.Context)                 {}
func (m *mockMastodon) GetConversationsV1(c *gin.Context)               {}
func (m *mockMastodon) RemoveConversationV1(c *gin.Context)             {}
func (m *mockMastodon) MarkConversationAsReadV1(c *gin.Context)         {}
func (m *mockMastodon) GetListsV1(c *gin.Context)                       {}
func (m *mockMastodon) GetListV1(c *gin.Context)                        {}
func (m *mockMastodon) CreateListV1(c *gin.Context)                     {}
func (m *mockMastodon) UpdateListV1(c *gin.Context)                     {}
func (m *mockMastodon) DeleteListV1(c *gin.Context)                     {}
func (m *mockMastodon) GetListAccountV1(c *gin.Context)                 {}
func (m *mockMastodon) AddListAccountV1(c *gin.Context)                 {}
func (m *mockMastodon) RemoveListAccountV1(c *gin.Context)              {}
func (m *mockMastodon) GetMarkersV1(c *gin.Context)                     {}
func (m *mockMastodon) CreateMarkerV1(c *gin.Context)                   {}
func (m *mockMastodon) StreamingForHealthCheckV1(c *gin.Context)        {}
func (m *mockMastodon) StreamingForUserV1(c *gin.Context)               {}
func (m *mockMastodon) StreamingForPublicV1(c *gin.Context)             {}
func (m *mockMastodon) StreamingForLocalV1(c *gin.Context)              {}
func (m *mockMastodon) StreamingForPublicHashtagV1(c *gin.Context)      {}
func (m *mockMastodon) StreamingForLocalHashtagV1(c *gin.Context)       {}
func (m *mockMastodon) StreamingForListV1(c *gin.Context)               {}
func (m *mockMastodon) StreamingForDirectV1(c *gin.Context)             {}
func (m *mockMastodon) StreamingV1(c *gin.Context)                      {}
func (m *mockMastodon) GetNotificationsV1(c *gin.Context)               {}
func (m *mockMastodon) GetNotificationV1(c *gin.Context)                {}
func (m *mockMastodon) DissmissNotificationsV1(c *gin.Context)          {}
func (m *mockMastodon) DissmissNotificationV1(c *gin.Context)           {}
func (m *mockMastodon) DeprecatedDissmissNotificationV1(c *gin.Context) {}
func (m *mockMastodon) SubscribeNotificationsV1(c *gin.Context)         {}
func (m *mockMastodon) GetSubscriptionV1(c *gin.Context)                {}
func (m *mockMastodon) UpdateSubscriptionV1(c *gin.Context)             {}
func (m *mockMastodon) UnsubscribeNotificationsV1(c *gin.Context)       {}
func (m *mockMastodon) SearchV1(c *gin.Context)                         {}
func (m *mockMastodon) SearchV2(c *gin.Context)                         {}
func (m *mockMastodon) GetInstanceV1(c *gin.Context)                    {}
func (m *mockMastodon) GetPeersV1(c *gin.Context)                       {}
func (m *mockMastodon) GetInstanceActivityV1(c *gin.Context)            {}
func (m *mockMastodon) TrendsV1(c *gin.Context)                         {}
func (m *mockMastodon) DirectoryV1(c *gin.Context)                      {}
func (m *mockMastodon) CustomEmojisV1(c *gin.Context)                   {}
func (m *mockMastodon) AdminAccountsV1(c *gin.Context)                  {}
func (m *mockMastodon) AdminAccountV1(c *gin.Context)                   {}
func (m *mockMastodon) AdminAccountActionV1(c *gin.Context)             {}
func (m *mockMastodon) AdminAccountApproveV1(c *gin.Context)            {}
func (m *mockMastodon) AdminAccountRejectV1(c *gin.Context)             {}
func (m *mockMastodon) AdminAccountEnableV1(c *gin.Context)             {}
func (m *mockMastodon) AdminAccountUnsilenceV1(c *gin.Context)          {}
func (m *mockMastodon) AdminAccountUnsuspendV1(c *gin.Context)          {}
func (m *mockMastodon) AdminReportsV1(c *gin.Context)                   {}
func (m *mockMastodon) AdminReportV1(c *gin.Context)                    {}
func (m *mockMastodon) AdminReportAssignV1(c *gin.Context)              {}
func (m *mockMastodon) AdminReportUnassignV1(c *gin.Context)            {}
func (m *mockMastodon) AdminReportResolveV1(c *gin.Context)             {}
func (m *mockMastodon) AdminReportReopenV1(c *gin.Context)              {}
func (m *mockMastodon) GetAnnouncementsV1(c *gin.Context)               {}
func (m *mockMastodon) DismissAnnouncementsV1(c *gin.Context)           {}
func (m *mockMastodon) ReactAnnouncementsV1(c *gin.Context)             {}
func (m *mockMastodon) UndoReactAnnouncementsV1(c *gin.Context)         {}
func (m *mockMastodon) ApiProofs(c *gin.Context)                        {}
func (m *mockMastodon) ApiOEmbed(c *gin.Context)                        {}

func TestFeatures_Routes(t *testing.T) {
	m := &mockMastodon{}
	features := Features(m)
	if features == nil {
		t.Fatal("Features should not be nil")
	}

	want := []string{
		"/oauth/authorize",
		"/api/v1/accounts",
		"/api/v1/accounts/:id/follow",
		"/api/v1/featured_tags/suggestions",
		"/api/v1/tags/:name/follow",
	}
	got := map[string]bool{}

	for _, r := range features.GetRoutes() {
		got[r.Path] = true
	}
	for _, path := range want {
		if !got[path] {
			t.Errorf("route %s not found in Features", path)
		}
	}
}
