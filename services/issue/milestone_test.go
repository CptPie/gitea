// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package issue

import (
	"testing"

	"code.gitea.io/gitea/models"
	issues_model "code.gitea.io/gitea/models/issues"
	"code.gitea.io/gitea/models/unittest"
	user_model "code.gitea.io/gitea/models/user"

	"github.com/stretchr/testify/assert"
)

func TestChangeMilestoneAssign(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	issue := unittest.AssertExistsAndLoadBean(t, &models.Issue{RepoID: 1}).(*models.Issue)
	doer := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	assert.NotNil(t, issue)
	assert.NotNil(t, doer)

	oldMilestoneID := issue.MilestoneID
	issue.MilestoneID = 2
	assert.NoError(t, ChangeMilestoneAssign(issue, doer, oldMilestoneID))
	unittest.AssertExistsAndLoadBean(t, &models.Comment{
		IssueID:        issue.ID,
		Type:           models.CommentTypeMilestone,
		MilestoneID:    issue.MilestoneID,
		OldMilestoneID: oldMilestoneID,
	})
	unittest.CheckConsistencyFor(t, &issues_model.Milestone{}, &models.Issue{})
}
