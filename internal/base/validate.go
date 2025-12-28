package base

import "strings"

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, "Empty request!")
		return res
	}

	if strings.TrimSpace(req.UserID) == "" {
		res = append(res, "UserID is required")
	}

	if strings.TrimSpace(req.Title) == "" {
		res = append(res, "Title is required")
	}

	if strings.TrimSpace(req.Description) == "" {
		res = append(res, "Description is required")
	}

	return res
}
