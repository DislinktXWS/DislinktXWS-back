package api

import (
	pb "module/common/proto/user_service"
	"module/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:          user.Id.Hex(),
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Email:       user.Email,
		Phone:       user.Phone,
		Biography:   user.Biograpy,
		IsPublic:    user.IsPublic,
	}
	return userPb
}

func mapNewUser(userPb *pb.User) *domain.User {
	user := &domain.User{
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Username:    userPb.Username,
		DateOfBirth: userPb.DateOfBirth,
		Gender:      userPb.Gender,
		Email:       userPb.Email,
		Phone:       userPb.Phone,
	}
	return user
}

func mapEditUser(userPb *pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &domain.User{
		Id:          id,
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Username:    userPb.Username,
		DateOfBirth: userPb.DateOfBirth,
		Gender:      userPb.Gender,
		Email:       userPb.Email,
		Phone:       userPb.Phone,
		Biograpy:    userPb.Biography,
		IsPublic:    userPb.IsPublic,
	}
	return user
}

func mapAddEducation(educationPb *pb.Education) *domain.Education {
	education := &domain.Education{
		School:       educationPb.School,
		Degree:       educationPb.Degree,
		FieldOfStudy: educationPb.FieldOfStudy,
		StartDate:    educationPb.StartDate,
		EndDate:      educationPb.EndDate,
		Grade:        educationPb.Grade,
		Description:  educationPb.Description,
	}
	return education
}

func mapAddExperience(experiencePb *pb.Experience) *domain.Experience {
	experience := &domain.Experience{
		Title:       experiencePb.Title,
		CompanyName: experiencePb.CompanyName,
		StartDate:   experiencePb.StartDate,
		EndDate:     experiencePb.EndDate,
		Industry:    experiencePb.Industry,
		Description: experiencePb.Description,
	}
	return experience
}

func mapAddSkill(skillPb *pb.Skill) *domain.Skill {
	skill := &domain.Skill{
		Name:        skillPb.Name,
		Proficiency: mapProficiency(skillPb.Proficiency),
	}
	return skill
}

func mapProficiency(proficiency pb.Skill_SkillProficiency) domain.SkillProficiency {
	switch proficiency {
	case pb.Skill_novice:
		return domain.Novice
	case pb.Skill_advancedBeginner:
		return domain.AdvancedBeginner
	case pb.Skill_proficient:
		return domain.Proficient
	case pb.Skill_expert:
		return domain.Expert
	case pb.Skill_master:
		return domain.Master
	}
	return domain.Novice
}

func mapSkill(skill *domain.Skill) *pb.Skill {
	skillPb := &pb.Skill{
		Name:        skill.Name,
		Proficiency: mapProficiencyToPb(skill.Proficiency),
	}
	return skillPb
}

func mapExperience(experience *domain.Experience) *pb.Experience {
	experiencePb := &pb.Experience{
		Title:       experience.Title,
		CompanyName: experience.CompanyName,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
		Industry:    experience.Industry,
		Description: experience.Description,
	}
	return experiencePb
}

func mapEducation(education *domain.Education) *pb.Education {
	educationPb := &pb.Education{
		School:       education.School,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		StartDate:    education.StartDate,
		EndDate:      education.EndDate,
		Grade:        education.Grade,
		Description:  education.Description,
	}
	return educationPb
}

func mapProficiencyToPb(proficiency domain.SkillProficiency) pb.Skill_SkillProficiency {
	switch proficiency {
	case domain.Novice:
		return pb.Skill_novice
	case domain.AdvancedBeginner:
		return pb.Skill_advancedBeginner
	case domain.Proficient:
		return pb.Skill_proficient
	case domain.Expert:
		return pb.Skill_expert
	case domain.Master:
		return pb.Skill_master
	}
	return pb.Skill_novice
}
