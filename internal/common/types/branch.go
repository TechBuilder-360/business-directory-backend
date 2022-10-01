package types

type CreateBranch struct {
	BranchName string `json:"branch_name" validate:"required"`
	//Contact    model.Contact `json:"contact" `
	//Address    model.Address `json:"address"`
}

type Branch struct {
	BranchName string `json:"branch_name"`
	IsHQ       bool   `json:"is_hq"`
	//Address
	//Location
	//Contact
}
