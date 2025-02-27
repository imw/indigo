// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.actor.defs

import (
	"encoding/json"
	"fmt"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/lex/util"
)

// ActorDefs_AdultContentPref is a "adultContentPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_AdultContentPref
type ActorDefs_AdultContentPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#adultContentPref" cborgen:"$type,const=app.bsky.actor.defs#adultContentPref"`
	Enabled       bool   `json:"enabled" cborgen:"enabled"`
}

// ActorDefs_ContentLabelPref is a "contentLabelPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_ContentLabelPref
type ActorDefs_ContentLabelPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#contentLabelPref" cborgen:"$type,const=app.bsky.actor.defs#contentLabelPref"`
	Label         string `json:"label" cborgen:"label"`
	// labelerDid: Which labeler does this preference apply to? If undefined, applies globally.
	LabelerDid *string `json:"labelerDid,omitempty" cborgen:"labelerDid,omitempty"`
	Visibility string  `json:"visibility" cborgen:"visibility"`
}

// ActorDefs_FeedViewPref is a "feedViewPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_FeedViewPref
type ActorDefs_FeedViewPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#feedViewPref" cborgen:"$type,const=app.bsky.actor.defs#feedViewPref"`
	// feed: The URI of the feed, or an identifier which describes the feed.
	Feed string `json:"feed" cborgen:"feed"`
	// hideQuotePosts: Hide quote posts in the feed.
	HideQuotePosts *bool `json:"hideQuotePosts,omitempty" cborgen:"hideQuotePosts,omitempty"`
	// hideReplies: Hide replies in the feed.
	HideReplies *bool `json:"hideReplies,omitempty" cborgen:"hideReplies,omitempty"`
	// hideRepliesByLikeCount: Hide replies in the feed if they do not have this number of likes.
	HideRepliesByLikeCount *int64 `json:"hideRepliesByLikeCount,omitempty" cborgen:"hideRepliesByLikeCount,omitempty"`
	// hideRepliesByUnfollowed: Hide replies in the feed if they are not by followed users.
	HideRepliesByUnfollowed *bool `json:"hideRepliesByUnfollowed,omitempty" cborgen:"hideRepliesByUnfollowed,omitempty"`
	// hideReposts: Hide reposts in the feed.
	HideReposts *bool `json:"hideReposts,omitempty" cborgen:"hideReposts,omitempty"`
}

// ActorDefs_HiddenPostsPref is a "hiddenPostsPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_HiddenPostsPref
type ActorDefs_HiddenPostsPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#hiddenPostsPref" cborgen:"$type,const=app.bsky.actor.defs#hiddenPostsPref"`
	// items: A list of URIs of posts the account owner has hidden.
	Items []string `json:"items" cborgen:"items"`
}

// ActorDefs_InterestsPref is a "interestsPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_InterestsPref
type ActorDefs_InterestsPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#interestsPref" cborgen:"$type,const=app.bsky.actor.defs#interestsPref"`
	// tags: A list of tags which describe the account owner's interests gathered during onboarding.
	Tags []string `json:"tags" cborgen:"tags"`
}

// ActorDefs_LabelerPrefItem is a "labelerPrefItem" in the app.bsky.actor.defs schema.
type ActorDefs_LabelerPrefItem struct {
	Did string `json:"did" cborgen:"did"`
}

// ActorDefs_LabelersPref is a "labelersPref" in the app.bsky.actor.defs schema.
type ActorDefs_LabelersPref struct {
	Labelers []*ActorDefs_LabelerPrefItem `json:"labelers" cborgen:"labelers"`
}

// ActorDefs_MutedWord is a "mutedWord" in the app.bsky.actor.defs schema.
//
// A word that the account owner has muted.
type ActorDefs_MutedWord struct {
	// targets: The intended targets of the muted word.
	Targets []*string `json:"targets" cborgen:"targets"`
	// value: The muted word itself.
	Value string `json:"value" cborgen:"value"`
}

// ActorDefs_MutedWordsPref is a "mutedWordsPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_MutedWordsPref
type ActorDefs_MutedWordsPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#mutedWordsPref" cborgen:"$type,const=app.bsky.actor.defs#mutedWordsPref"`
	// items: A list of words the account owner has muted.
	Items []*ActorDefs_MutedWord `json:"items" cborgen:"items"`
}

// ActorDefs_PersonalDetailsPref is a "personalDetailsPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_PersonalDetailsPref
type ActorDefs_PersonalDetailsPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#personalDetailsPref" cborgen:"$type,const=app.bsky.actor.defs#personalDetailsPref"`
	// birthDate: The birth date of account owner.
	BirthDate *string `json:"birthDate,omitempty" cborgen:"birthDate,omitempty"`
}

type ActorDefs_Preferences_Elem struct {
	ActorDefs_AdultContentPref    *ActorDefs_AdultContentPref
	ActorDefs_ContentLabelPref    *ActorDefs_ContentLabelPref
	ActorDefs_SavedFeedsPref      *ActorDefs_SavedFeedsPref
	ActorDefs_PersonalDetailsPref *ActorDefs_PersonalDetailsPref
	ActorDefs_FeedViewPref        *ActorDefs_FeedViewPref
	ActorDefs_ThreadViewPref      *ActorDefs_ThreadViewPref
	ActorDefs_InterestsPref       *ActorDefs_InterestsPref
	ActorDefs_MutedWordsPref      *ActorDefs_MutedWordsPref
	ActorDefs_HiddenPostsPref     *ActorDefs_HiddenPostsPref
}

func (t *ActorDefs_Preferences_Elem) MarshalJSON() ([]byte, error) {
	if t.ActorDefs_AdultContentPref != nil {
		t.ActorDefs_AdultContentPref.LexiconTypeID = "app.bsky.actor.defs#adultContentPref"
		return json.Marshal(t.ActorDefs_AdultContentPref)
	}
	if t.ActorDefs_ContentLabelPref != nil {
		t.ActorDefs_ContentLabelPref.LexiconTypeID = "app.bsky.actor.defs#contentLabelPref"
		return json.Marshal(t.ActorDefs_ContentLabelPref)
	}
	if t.ActorDefs_SavedFeedsPref != nil {
		t.ActorDefs_SavedFeedsPref.LexiconTypeID = "app.bsky.actor.defs#savedFeedsPref"
		return json.Marshal(t.ActorDefs_SavedFeedsPref)
	}
	if t.ActorDefs_PersonalDetailsPref != nil {
		t.ActorDefs_PersonalDetailsPref.LexiconTypeID = "app.bsky.actor.defs#personalDetailsPref"
		return json.Marshal(t.ActorDefs_PersonalDetailsPref)
	}
	if t.ActorDefs_FeedViewPref != nil {
		t.ActorDefs_FeedViewPref.LexiconTypeID = "app.bsky.actor.defs#feedViewPref"
		return json.Marshal(t.ActorDefs_FeedViewPref)
	}
	if t.ActorDefs_ThreadViewPref != nil {
		t.ActorDefs_ThreadViewPref.LexiconTypeID = "app.bsky.actor.defs#threadViewPref"
		return json.Marshal(t.ActorDefs_ThreadViewPref)
	}
	if t.ActorDefs_InterestsPref != nil {
		t.ActorDefs_InterestsPref.LexiconTypeID = "app.bsky.actor.defs#interestsPref"
		return json.Marshal(t.ActorDefs_InterestsPref)
	}
	if t.ActorDefs_MutedWordsPref != nil {
		t.ActorDefs_MutedWordsPref.LexiconTypeID = "app.bsky.actor.defs#mutedWordsPref"
		return json.Marshal(t.ActorDefs_MutedWordsPref)
	}
	if t.ActorDefs_HiddenPostsPref != nil {
		t.ActorDefs_HiddenPostsPref.LexiconTypeID = "app.bsky.actor.defs#hiddenPostsPref"
		return json.Marshal(t.ActorDefs_HiddenPostsPref)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ActorDefs_Preferences_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.actor.defs#adultContentPref":
		t.ActorDefs_AdultContentPref = new(ActorDefs_AdultContentPref)
		return json.Unmarshal(b, t.ActorDefs_AdultContentPref)
	case "app.bsky.actor.defs#contentLabelPref":
		t.ActorDefs_ContentLabelPref = new(ActorDefs_ContentLabelPref)
		return json.Unmarshal(b, t.ActorDefs_ContentLabelPref)
	case "app.bsky.actor.defs#savedFeedsPref":
		t.ActorDefs_SavedFeedsPref = new(ActorDefs_SavedFeedsPref)
		return json.Unmarshal(b, t.ActorDefs_SavedFeedsPref)
	case "app.bsky.actor.defs#personalDetailsPref":
		t.ActorDefs_PersonalDetailsPref = new(ActorDefs_PersonalDetailsPref)
		return json.Unmarshal(b, t.ActorDefs_PersonalDetailsPref)
	case "app.bsky.actor.defs#feedViewPref":
		t.ActorDefs_FeedViewPref = new(ActorDefs_FeedViewPref)
		return json.Unmarshal(b, t.ActorDefs_FeedViewPref)
	case "app.bsky.actor.defs#threadViewPref":
		t.ActorDefs_ThreadViewPref = new(ActorDefs_ThreadViewPref)
		return json.Unmarshal(b, t.ActorDefs_ThreadViewPref)
	case "app.bsky.actor.defs#interestsPref":
		t.ActorDefs_InterestsPref = new(ActorDefs_InterestsPref)
		return json.Unmarshal(b, t.ActorDefs_InterestsPref)
	case "app.bsky.actor.defs#mutedWordsPref":
		t.ActorDefs_MutedWordsPref = new(ActorDefs_MutedWordsPref)
		return json.Unmarshal(b, t.ActorDefs_MutedWordsPref)
	case "app.bsky.actor.defs#hiddenPostsPref":
		t.ActorDefs_HiddenPostsPref = new(ActorDefs_HiddenPostsPref)
		return json.Unmarshal(b, t.ActorDefs_HiddenPostsPref)

	default:
		return nil
	}
}

// ActorDefs_ProfileAssociated is a "profileAssociated" in the app.bsky.actor.defs schema.
type ActorDefs_ProfileAssociated struct {
	Feedgens *int64 `json:"feedgens,omitempty" cborgen:"feedgens,omitempty"`
	Labeler  *bool  `json:"labeler,omitempty" cborgen:"labeler,omitempty"`
	Lists    *int64 `json:"lists,omitempty" cborgen:"lists,omitempty"`
}

// ActorDefs_ProfileView is a "profileView" in the app.bsky.actor.defs schema.
type ActorDefs_ProfileView struct {
	Avatar      *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Description *string                            `json:"description,omitempty" cborgen:"description,omitempty"`
	Did         string                             `json:"did" cborgen:"did"`
	DisplayName *string                            `json:"displayName,omitempty" cborgen:"displayName,omitempty"`
	Handle      string                             `json:"handle" cborgen:"handle"`
	IndexedAt   *string                            `json:"indexedAt,omitempty" cborgen:"indexedAt,omitempty"`
	Labels      []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Viewer      *ActorDefs_ViewerState             `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// ActorDefs_ProfileViewBasic is a "profileViewBasic" in the app.bsky.actor.defs schema.
type ActorDefs_ProfileViewBasic struct {
	Avatar      *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Did         string                             `json:"did" cborgen:"did"`
	DisplayName *string                            `json:"displayName,omitempty" cborgen:"displayName,omitempty"`
	Handle      string                             `json:"handle" cborgen:"handle"`
	Labels      []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	Viewer      *ActorDefs_ViewerState             `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// ActorDefs_ProfileViewDetailed is a "profileViewDetailed" in the app.bsky.actor.defs schema.
type ActorDefs_ProfileViewDetailed struct {
	Associated     *ActorDefs_ProfileAssociated       `json:"associated,omitempty" cborgen:"associated,omitempty"`
	Avatar         *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Banner         *string                            `json:"banner,omitempty" cborgen:"banner,omitempty"`
	Description    *string                            `json:"description,omitempty" cborgen:"description,omitempty"`
	Did            string                             `json:"did" cborgen:"did"`
	DisplayName    *string                            `json:"displayName,omitempty" cborgen:"displayName,omitempty"`
	FollowersCount *int64                             `json:"followersCount,omitempty" cborgen:"followersCount,omitempty"`
	FollowsCount   *int64                             `json:"followsCount,omitempty" cborgen:"followsCount,omitempty"`
	Handle         string                             `json:"handle" cborgen:"handle"`
	IndexedAt      *string                            `json:"indexedAt,omitempty" cborgen:"indexedAt,omitempty"`
	Labels         []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	PostsCount     *int64                             `json:"postsCount,omitempty" cborgen:"postsCount,omitempty"`
	Viewer         *ActorDefs_ViewerState             `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// ActorDefs_SavedFeedsPref is a "savedFeedsPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_SavedFeedsPref
type ActorDefs_SavedFeedsPref struct {
	LexiconTypeID string   `json:"$type,const=app.bsky.actor.defs#savedFeedsPref" cborgen:"$type,const=app.bsky.actor.defs#savedFeedsPref"`
	Pinned        []string `json:"pinned" cborgen:"pinned"`
	Saved         []string `json:"saved" cborgen:"saved"`
	TimelineIndex *int64   `json:"timelineIndex,omitempty" cborgen:"timelineIndex,omitempty"`
}

// ActorDefs_ThreadViewPref is a "threadViewPref" in the app.bsky.actor.defs schema.
//
// RECORDTYPE: ActorDefs_ThreadViewPref
type ActorDefs_ThreadViewPref struct {
	LexiconTypeID string `json:"$type,const=app.bsky.actor.defs#threadViewPref" cborgen:"$type,const=app.bsky.actor.defs#threadViewPref"`
	// prioritizeFollowedUsers: Show followed users at the top of all replies.
	PrioritizeFollowedUsers *bool `json:"prioritizeFollowedUsers,omitempty" cborgen:"prioritizeFollowedUsers,omitempty"`
	// sort: Sorting mode for threads.
	Sort *string `json:"sort,omitempty" cborgen:"sort,omitempty"`
}

// ActorDefs_ViewerState is a "viewerState" in the app.bsky.actor.defs schema.
//
// Metadata about the requesting account's relationship with the subject account. Only has meaningful content for authed requests.
type ActorDefs_ViewerState struct {
	BlockedBy      *bool                    `json:"blockedBy,omitempty" cborgen:"blockedBy,omitempty"`
	Blocking       *string                  `json:"blocking,omitempty" cborgen:"blocking,omitempty"`
	BlockingByList *GraphDefs_ListViewBasic `json:"blockingByList,omitempty" cborgen:"blockingByList,omitempty"`
	FollowedBy     *string                  `json:"followedBy,omitempty" cborgen:"followedBy,omitempty"`
	Following      *string                  `json:"following,omitempty" cborgen:"following,omitempty"`
	Muted          *bool                    `json:"muted,omitempty" cborgen:"muted,omitempty"`
	MutedByList    *GraphDefs_ListViewBasic `json:"mutedByList,omitempty" cborgen:"mutedByList,omitempty"`
}
