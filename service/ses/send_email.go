package ses

import (
	"errors"
	"regexp"
)

func (c *Client) SendEmail(data SendEmailInput) (map[string]interface{}, error) {
	// TODO: implement actual logic to send email

	c.mu.Lock()
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(data.Source) {
		c.TotalFaildCount++
		c.mu.Unlock()
		return nil, errors.New("Invalid Sender Email")
	}
	for _, email := range data.Destination.CcAddresses {
		if !re.MatchString(email) {
			c.TotalFaildCount++
			c.mu.Unlock()
			return nil, errors.New("Invalid CcAddr: " + email)
		}
	}
	for _, email := range data.Destination.ToAddresses {
		if !re.MatchString(email) {
			c.TotalFaildCount++
			c.mu.Unlock()
			return nil, errors.New("Invalid ToAddr: " + email)
		}
	}
	c.TotalSuccessCount += uint64(len(data.Destination.ToAddresses))
	c.mu.Unlock()
	return map[string]interface{}{
		"status": "ok",
		"count":  len(data.Destination.ToAddresses),
	}, nil
}
