package raideriogo

type EncountersDefeated struct {
}

type RaidRanking struct {
}

type ViewRaidingRaidRankingsResponse struct {
}

/*
public class EncountersDefeated
{
    public string slug { get; set; }
    public DateTime lastDefeated { get; set; }
    public DateTime firstDefeated { get; set; }
}

public class RaidRanking
{
    public int rank { get; set; }
    public int regionRank { get; set; }
    public Guild guild { get; set; }
    public List<EncountersDefeated> encountersDefeated { get; set; }
}

public class RootObject
{
    public List<RaidRanking> raidRankings { get; set; }
}
*/
