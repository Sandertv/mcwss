package event

import "github.com/sandertv/mcwss/mctype"

// Many of the available screen names are listed below. Note that various screens, such as realms error
// screens, marketplace screens etc. are dynamic and are thus not listed below.
const (
	ScreenStart                    = "start_screen"
	ScreenSettingsGame             = "screen_world_controls_and_settings - game_tab"
	ScreenSettingsMultiplayer      = "screen_world_controls_and_settings - multiplayer_tab"
	ScreenSettingsKeyboardAndMouse = "screen_world_controls_and_settings - keyboard_and_mouse_tab"
	ScreenSettingsController       = "screen_world_controls_and_settings - controller_tab"
	ScreenSettingsTouch            = "screen_world_controls_and_settings - touch_tab"
	ScreenSettingsProfile          = "screen_world_controls_and_settings - profile_tab"
	ScreenSettingsVideo            = "screen_world_controls_and_settings - video_tab"
	ScreenSettingsSound            = "screen_world_controls_and_settings - sound_tab"
	ScreenSettingsGlobalResources  = "screen_world_controls_and_settings - global_texture_pack_tab"
	ScreenSettingsHowToPlay        = "screen_world_controls_and_settings - how_to_play"
	ScreenSettingsResourcePacks    = "screen_world_controls_and_settings - level_texture_pack_tab"
	ScreenSettingsBehaviourPacks   = "screen_world_controls_and_settings - addon_tab"
	ScreenSettingsLanguage         = "screen_controls_and_settings - language_tab"
	ScreenSettingsStorage          = "screen_controls_and_settings - storage_tab"
	ScreenAchievement              = "achievement_screen"
	ScreenStore                    = "store_data_driven_screen - store_home_screen"

	ScreenWorlds                    = "play_screen - worlds"
	ScreenWorldTemplates            = "world_templates_screen"
	ScreenWorldCreateGame           = "screen_world_create - game_tab"
	ScreenWorldCreateMultiplayer    = "screen_world_create - multiplayer_tab"
	ScreenWorldCreateTexturePacks   = "screen_world_create - level_texture_pack_tab"
	ScreenWorldCreateBehaviourPacks = "screen_world_create - addon_tab"

	ScreenRealmCreate     = "realms_create_screen"
	ScreenRealmJoinByCode = "friends  - join_by_code"

	ScreenWorldLoadingProgress                 = "world_loading_progress_screen - local_world_load"
	ScreenWorldLoadingJoiningMultiplayerServer = "world_loading_progress_screen - joining_multiplayer_external_server"
	ScreenWorldSavingProgress                  = "world_saving_progress_screen"

	ScreenRatingPrompt = "rating_prompt_screen"

	ScreenFriends            = "play_screen - friends"
	ScreenServers            = "play_screen - servers"
	ScreenAddExternalServer  = "add_external_server_screen_new"
	ScreenEditExternalServer = "add_external_server_screen_edit"

	ScreenInGamePlay       = "in_game_play_screen"
	ScreenChat             = "chat_screen"
	ScreenPause            = "pause_screen"
	ScreenInventory        = "inventory_screen"
	ScreenFeedback         = "feedbackScreen.title - feedbackScreen.body"
	ScreenSkinPicker       = "skin_picker_screen"
	ScreenExpandedSkinPack = "expanded_skin_pack_screen"
	ScreenInvite           = "invite_screen"
	ScreenHowToPlay        = "how_to_play_screen"
)

// ScreenChanged is sent by the client when it changes from one screen to another. This does not include just
// 'any' screen change, but only when the screen 'state' changes, for example when opening the chat screen.
type ScreenChanged struct {
	// Direction seems to always be 'forward'. It's unclear when this is not forward.
	Direction mctype.Direction
	// PreviousScreenName is the name of the screen displayed before it was changed to this new screen.
	PreviousScreenName string
	// ScreenName is the name of the new screen opened. For options this field might contain, see the
	// constants above. Not all screen names are listed above, provided some of them are dynamic.
	ScreenName string
	// ScreenVersion is the version of the screen. Currently, it appears this value is always '1'.
	ScreenVersion string
	// Seconds is a float describing how long the previous screen was open for.
	Seconds float64
	// Timestamp is a timestamp indicating the time the screen was changed relative to the time the player
	// opened the game.
	Timestamp float64 `json:"TimeStamp"`
}
