package raideriogo

// ViewRaidingBossRankingsResponse defines the schema for tracking raid boss rankings.
type ViewRaidingBossRankingsResponse struct {
	BossRankings []BossRanking `json:""`
}

// Guild defines the schema for a guild.
type Guild struct {
}

// BossRanking defines the schema for a boss ranking.
type BossRanking struct {
}

/*
public class Guild
{
    public int id { get; set; }
    public string name { get; set; }
    public string faction { get; set; }
    public string logoUrl { get; set; }
    public Realm realm { get; set; }
    public Region region { get; set; }
}

public class BossRanking
{
    public int rank { get; set; }
    public int regionRank { get; set; }
    public Guild guild { get; set; }
    public List<EncountersDefeated> encountersDefeated { get; set; }
}

public class ViewRaidingBossRankingsResponse
{
    public List<BossRanking> bossRankings { get; set; }
}
*/
