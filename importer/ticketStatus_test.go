// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package importer_test

import "testing"

func TestImportTicketClose(t *testing.T) {
	setUpTickets(t)
	defer tearDown(t)

	// first thing to expect is retrieval of ticket from Trac
	expectTracTicketRetrievals(t, openTicket)

	// expect all actions for creating Gitea issue from Trac ticket
	expectAllTicketActions(t, openTicket)

	// expect trac to return us no attachments
	expectTracAttachmentRetrievals(t, openTicket)

	// expect trac to return us one close ticket change
	expectTracChangeRetrievals(t, openTicket, closeTicketChange)

	// expect all actions for creating Gitea comments from Trac ticket status changes
	expectAllTicketStatusActions(t, openTicket, closeTicketChange)

	// expect issue update time to be updated
	expectIssueUpdateTimeSetToLatestOf(t, openTicket, closeTicketChange)

	// expect issue comment count to be updated
	expectIssueCommentCountUpdate(t, openTicket)

	// expect all issue counts to be updated
	expectIssueCountUpdates(t)

	// expect to convert ticket description to markdown
	expectDescriptionMarkdownConversion(t, openTicket)

	// expect to update Gitea issue description
	expectIssueDescriptionUpdates(t, openTicket.issueID, openTicket.descriptionMarkdown)

	dataImporter.ImportTickets(userMap, componentMap, priorityMap, resolutionMap, severityMap, typeMap, versionMap, revisionMap)
}

func TestImportTicketReopen(t *testing.T) {
	setUpTickets(t)
	defer tearDown(t)

	// first thing to expect is retrieval of ticket from Trac
	expectTracTicketRetrievals(t, openTicket)

	// expect all actions for creating Gitea issue from Trac ticket
	expectAllTicketActions(t, openTicket)

	// expect trac to return us no attachments
	expectTracAttachmentRetrievals(t, openTicket)

	// expect trac to return us one reopen ticket change
	expectTracChangeRetrievals(t, openTicket, reopenTicketChange)

	// expect all actions for creating Gitea comments from Trac ticket status changes
	expectAllTicketStatusActions(t, openTicket, reopenTicketChange)

	// expect issue update time to be updated
	expectIssueUpdateTimeSetToLatestOf(t, openTicket, reopenTicketChange)

	// expect issue comment count to be updated
	expectIssueCommentCountUpdate(t, openTicket)

	// expect all issue counts to be updated
	expectIssueCountUpdates(t)

	// expect to convert ticket description to markdown
	expectDescriptionMarkdownConversion(t, openTicket)

	// expect to update Gitea issue description
	expectIssueDescriptionUpdates(t, openTicket.issueID, openTicket.descriptionMarkdown)

	dataImporter.ImportTickets(userMap, componentMap, priorityMap, resolutionMap, severityMap, typeMap, versionMap, revisionMap)
}
