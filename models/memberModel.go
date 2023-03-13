package models

// TO DO
// Trainer Member

type Member struct {
	MemberID      uint   `gorm:"primaryKey" json:"member_id"`
	MemberName    string `gorm:"index" json:"member_name"`
	MemberNIK     uint   `json:"member_nik"`
	MemberContact string `json:"member_contact"`
	MemberEmail   string `gorm:"uniqueIndex;size:256" json:"member_email"`
	MemberAddress string `json:"member_address"`
	MemberGender  string `json:"member_gender"`
	MemberWight   uint16 `json:"member_wight"`
	MemberPackage string `json:"member_package"`
}
