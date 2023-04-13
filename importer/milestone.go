// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/stevejefferson/trac2gitea/accessor/gitea"
	"github.com/stevejefferson/trac2gitea/accessor/trac"
	"github.com/stevejefferson/trac2gitea/log"
)

// ImportMilestones imports Trac milestones as Gitea milestones.
func (importer *Importer) ImportMilestones() error {
	err := importer.tracAccessor.GetMilestones(func(tracMilestone *trac.Milestone) error {
		if tracMilestone.Name == "" {
			log.Debug("skipping unnamed Trac milestone...")
			return nil
		}

		giteaMilestone := gitea.Milestone{
			Name:        tracMilestone.Name,
			Description: tracMilestone.Description,
			Closed:      tracMilestone.Completed != 0,
			DueTime:     tracMilestone.Due,
			ClosedTime:  tracMilestone.Completed,
			Created:     tracMilestone.Completed,
		}

		if tracMilestone.Due == 0 {
			giteaMilestone.DueTime = 253402300799 // 31st Dec 9999
		}

		milestoneID, err := importer.giteaAccessor.AddMilestone(&giteaMilestone)
		if err != nil {
			return err
		}

		log.Debug("added milestone (id %d) %s", milestoneID, tracMilestone.Name)
		return nil
	})
	if err != nil {
		return err
	}

	return importer.giteaAccessor.UpdateRepoMilestoneCounts()
}
