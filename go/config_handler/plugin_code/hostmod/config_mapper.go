package main

import (
	"fmt"
	"strings"
	"github.com/kerrgavin/srb2kart-go/config"
)

func mapHostmod(hostmod *hostmod) string {
	var configHelper config.ConfigHelper
	configHelper.Comment("hostmod configs")
	configHelper.ConfigBool("hm_automate", hostmod.Automate)
	configHelper.ConfigInt("hm_encore", hostmod.Encore)
	mapSpecbomb(&hostmod.Specbomb, &configHelper)
	mapNameFilter(&hostmod.NameFilter, &configHelper)
	mapBattle(&hostmod.Battle, &configHelper)
	mapMotd(&hostmod.Motd, &configHelper)
	mapRestat(&hostmod.Restat, &configHelper)
	mapShout(&hostmod.Shout, &configHelper)
	mapSchedule(&hostmod.Schedule, &configHelper)
	mapVote(&hostmod.Vote, &configHelper)
	mapCustomCheck(&hostmod.CustomCheck, &configHelper)
	mapScoreboard(&hostmod.Scoreboard, &configHelper)
	mapVeto(&hostmod.Veto, &configHelper)
	return configHelper.String()
}

func mapSpecbomb(specbomb *specbomb, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_specbomb", specbomb.Enabled)
	configHelper.ConfigBool("hm_specbomb_antisoftlock", specbomb.AntiSoftlock)
}

func mapNameFilter(nameFilter *nameFilter, configHelper *config.ConfigHelper) {
	configHelper.ConfigString("hm_namefilter_mode", nameFilter.Mode)
	for _, name := range nameFilter.Names {
		configHelper.ConfigString("hm_namefilter", &name)
	}
}

func mapBattle(battle *battle, configHelper *config.ConfigHelper) {
	configHelper.ConfigInt("hm_timelimit", battle.Timelimit)
	configHelper.ConfigBool("hm_bail", battle.Bail)
}

func mapMotd(motd *motd, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_motd", motd.Enabled)
	configHelper.ConfigBool("hm_motd_nag", motd.Nag)
	configHelper.ConfigString("hm_motd_bg", motd.Background)
	configHelper.ConfigString("hm_motd_name", motd.Name)
	configHelper.ConfigString("hm_motd_contact", motd.Contact)
	configHelper.ConfigString("hm_motd_tagline", motd.Tagline)
}

func mapRestat(restat *restat, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_restat", restat.Enabled)
	configHelper.ConfigBool("hm_restat_notify", restat.Notify)
}

func mapShout(shout *shout, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_autoshout", shout.Autoshout)
	configHelper.ConfigString("hm_shout_name", shout.Name)
	configHelper.ConfigString("hm_shout_color", shout.Color)
}

func mapSchedule(schedule *schedule, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_schedule", schedule.Enabled)
	for _, job := range schedule.Jobs {
		var stringBuilder strings.Builder
		stringBuilder.WriteString(fmt.Sprintf("%d", *job.Timer))
		for index, command := range job.Commands {
			if index == 0 {
				stringBuilder.WriteString(fmt.Sprintf(" %s", command))
			} else {
				stringBuilder.WriteString(fmt.Sprintf(" [%s]", command))
			}
		}
		formatedCommand := stringBuilder.String()
		configHelper.ConfigString("hm_schedule_add", &formatedCommand)
	}
}

func mapVote(vote *vote, configHelper *config.ConfigHelper) {
	configHelper.ConfigInt("hm_vote_timer", vote.Timer)
	configHelper.ConfigInt("hm_vote_autopass", vote.Autopass)
	configHelper.ConfigBool("hm_vote_allowidc", vote.Allowidc)
	for _, command := range vote.Whitelist {
		configHelper.ConfigString("hm_votable", &command)
	}
}

func mapCustomCheck(customcheck *customcheck, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_customcheck", customcheck.Enabled)
	configHelper.ConfigBool("hm_customcheck_1", customcheck.CheckOne)
	configHelper.ConfigBool("hm_customcheck_2", customcheck.CheckTwo)
	configHelper.ConfigBool("hm_customcheck_3", customcheck.CheckThree)
}

func mapScoreboard(scoreboard *scoreboard, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_scoreboard", scoreboard.Enabled)
	configHelper.ConfigBool("hm_scoreboard_humor", scoreboard.Humor)
	for _, line := range scoreboard.Lines {
		configHelper.ConfigString("hm_scoreboard_addline", &line)
	}
	for _, modLine := range scoreboard.ModLines {
		configHelper.ConfigString("hm_scoreboard_addmod", &modLine)
	}
}

func mapVeto(veto *veto, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("hm_veto", veto.Enabled)
	configHelper.ConfigInt("hm_veto_threshold", veto.Threshold)
	configHelper.ConfigInt("hm_veto_hellclosed", veto.Hellclosed)
	configHelper.ConfigInt("hm_veto_hellopen", veto.Hellopen)
}
