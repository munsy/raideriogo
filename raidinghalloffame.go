package raideriogo

/*
//ViewRaidingHallOfFameResponse defines the schema for tracking the hall of fame.
type ViewRaidingHallOfFameResponse struct {
}
*/

// ViewRaidingHallOfFameResponse defines the schema for tracking the hall of fame.
type ViewRaidingHallOfFameResponse struct {
	HallOfFame struct {
		BossKills []struct {
			Boss        string `json:"boss"`
			BossSummary struct {
				EncounterID int    `json:"encounterId"`
				Name        string `json:"name"`
				Slug        string `json:"slug"`
				Ordinal     int    `json:"ordinal"`
				WingID      int    `json:"wingId"`
				IconURL     string `json:"iconUrl"`
			} `json:"bossSummary"`
			DefeatedBy struct {
				TotalCount int `json:"totalCount"`
				Guilds     []struct {
					Guild struct {
						ID      int    `json:"id"`
						Name    string `json:"name"`
						Faction string `json:"faction"`
						LogoURL string `json:"logoUrl"`
						Realm   struct {
							ID          int    `json:"id"`
							Name        string `json:"name"`
							Slug        string `json:"slug"`
							AltSlug     string `json:"altSlug"`
							Locale      string `json:"locale"`
							IsConnected bool   `json:"isConnected"`
						} `json:"realm"`
						Region struct {
							Name      string `json:"name"`
							Slug      string `json:"slug"`
							ShortName string `json:"short_name"`
						} `json:"region"`
					} `json:"guild"`
					DefeatedAt string `json:"defeatedAt"`
				} `json:"guilds"`
			} `json:"defeatedBy"`
		} `json:"bossKills"`
		WinningGuilds []struct {
			Rank  int `json:"rank"`
			Guild struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Faction string `json:"faction"`
				LogoURL string `json:"logoUrl"`
				Realm   struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Slug        string `json:"slug"`
					AltSlug     string `json:"altSlug"`
					Locale      string `json:"locale"`
					IsConnected bool   `json:"isConnected"`
				} `json:"realm"`
				Region struct {
					Name      string `json:"name"`
					Slug      string `json:"slug"`
					ShortName string `json:"short_name"`
				} `json:"region"`
			} `json:"guild"`
			EncountersDefeated []interface{} `json:"encountersDefeated"`
		} `json:"winningGuilds"`
	} `json:"hallOfFame"`
}

/*
type BossSummary struct {
}

type GuildDefeats struct {
}

type Guild struct {
}

type DefeatedBy struct {
}

type BossKill struct {
}

type Realm2 struct {
}

type Region2 struct {
}

type Guild3 struct {
}

type WinningGuild struct {
}

type HallOfFame struct {
}
*/

/*
public class BossSummary
{
    public int encounterId { get; set; }
    public string name { get; set; }
    public string slug { get; set; }
    public int ordinal { get; set; }
    public int wingId { get; set; }
    public string iconUrl { get; set; }
}

public class Guild
{
    public int id { get; set; }
    public string name { get; set; }
    public string faction { get; set; }
    public string logoUrl { get; set; }
    public Realm realm { get; set; }
    public Region region { get; set; }
}

public class GuildDefeats
{
    public Guild guild { get; set; }
    public string defeatedAt { get; set; }
}

public class DefeatedBy
{
    public int totalCount { get; set; }
    public List<Guild> guilds { get; set; }
}

public class BossKill
{
    public string boss { get; set; }
    public BossSummary bossSummary { get; set; }
    public DefeatedBy defeatedBy { get; set; }
}

public class Realm2
{
    public int id { get; set; }
    public string name { get; set; }
    public string slug { get; set; }
    public string altSlug { get; set; }
    public string locale { get; set; }
    public bool isConnected { get; set; }
}

public class Region2
{
    public string name { get; set; }
    public string slug { get; set; }
    public string short_name { get; set; }
}

public class Guild3
{
    public int id { get; set; }
    public string name { get; set; }
    public string faction { get; set; }
    public string logoUrl { get; set; }
    public Realm2 realm { get; set; }
    public Region2 region { get; set; }
}

public class WinningGuild
{
    public int rank { get; set; }
    public Guild3 guild { get; set; }
    public List<object> encountersDefeated { get; set; }
}

public class HallOfFame
{
    public List<BossKill> bossKills { get; set; }
    public List<WinningGuild> winningGuilds { get; set; }
}

public class ViewRaidingHallOfFameResponse
{
    public HallOfFame hallOfFame { get; set; }
}
*/
