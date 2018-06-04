package raideriogo

// ViewRaidingHallOfFameResponse defines the schema for tracking the hall of fame.
type ViewRaidingHallOfFameResponse struct {
}

/*
type BossSummary struct {
}

type Guild2 struct {
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

public class Guild2
{
    public Guild2 guild { get; set; }
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
