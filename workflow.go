package main

import "fmt"

type matterIntake struct {
	MatterID       string `json:"matter_id"`
	ClientEmail    string `json:"client_email"`
	SignedDocument string `json:"signed_document"`
	Deadline       string `json:"deadline"`
}

type deliveryPlan struct {
	Subjects []string
}

func planMatterEmails(m matterIntake) deliveryPlan {
	return deliveryPlan{Subjects: []string{
		fmt.Sprintf("Matter %s intake receipt", m.MatterID),
		fmt.Sprintf("Signed document ready for matter %s", m.MatterID),
		fmt.Sprintf("Deadline follow-up for matter %s", m.MatterID),
	}}
}

func sendMatterEmails(c *emailClient, m matterIntake) ([]string, error) {
	plan := planMatterEmails(m)
	ids := make([]string, 0, len(plan.Subjects))
	bodies := []string{
		"We received your matter intake.",
		"Your signed document is ready for review.",
		"The matter deadline is " + m.Deadline + ".",
	}
	for i, subject := range plan.Subjects {
		id, err := c.send(map[string]string{"to": m.ClientEmail, "subject": subject, "body": bodies[i]}, m.MatterID+fmt.Sprintf("-%d", i))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
