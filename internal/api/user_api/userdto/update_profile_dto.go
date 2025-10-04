package userdto

import "social/internal/domain/user"

type UpdateProfileDTO struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty" validate:"omitempty,email"`
	Mobile *string `json:"mobile,omitempty"`
	Bio    *string `json:"bio,omitempty"`
}

func (dto *UpdateProfileDTO) ToDomain() *user.Profile {
	profile := &user.Profile{}

	if dto.Name != nil {
		profile.Name = *dto.Name
	}
	if dto.Email != nil {
		profile.Email = *dto.Email
	}
	if dto.Mobile != nil {
		profile.Mobile = *dto.Mobile
	}

	if dto.Bio != nil {
		profile.Bio = dto.Bio
	}

	return profile
}
