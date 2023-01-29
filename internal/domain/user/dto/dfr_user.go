package dto

type DtfUser struct {
	Message string `json:"message"`
	Result  struct {
		Id          int    `json:"id"`
		Url         string `json:"url"`
		Created     int    `json:"created"`
		CreatedRFC  string `json:"createdRFC"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Avatar      struct {
			Type string `json:"type"`
			Data struct {
				Uuid   string `json:"uuid"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Size   int    `json:"size"`
				Type   string `json:"type"`
				Color  string `json:"color"`
				Hash   string `json:"hash"`
			} `json:"data"`
		} `json:"avatar"`
		AvatarUrl      string      `json:"avatar_url"`
		Type           int         `json:"type"`
		Cover          interface{} `json:"cover"`
		Karma          int         `json:"karma"`
		AdvancedAccess struct {
			IsNeedsAdvancedAccess bool   `json:"is_needs_advanced_access"`
			Hash                  string `json:"hash"`
		} `json:"advanced_access"`
		IsAvailableForMessenger bool `json:"isAvailableForMessenger"`
		IsVerified              bool `json:"isVerified"`
		IsMuted                 bool `json:"is_muted"`
		IsVerified1             bool `json:"is_verified"`
		Counters                struct {
			Entries    int `json:"entries"`
			Comments   int `json:"comments"`
			Favourites int `json:"favourites"`
		} `json:"counters"`
		IsSubscribedToNewPosts bool `json:"is_subscribed_to_new_posts"`
		SubscribersCount       int  `json:"subscribers_count"`
		SubscribersAvatars     []struct {
			AvatarUrl string `json:"avatar_url"`
			Name      string `json:"name"`
		} `json:"subscribers_avatars"`
		IsSubscribed bool `json:"is_subscribed"`
		IsPlus       bool `json:"is_plus"`
	} `json:"result"`
}
