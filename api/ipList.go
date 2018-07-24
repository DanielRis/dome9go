package api

func (c *Client) CreateIPList(ipList IPList) (*IPList, error) {
	req := struct {
		Name IPList `json:"name"`
	}{
		Name: ipList,
	}

	resp := struct {
		Name IPList `json:"name"`
	}{}

	_, err := c.Do("POST", "IpList", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Name, nil
}
