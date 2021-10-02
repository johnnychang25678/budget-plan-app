package models

type SpaceMember struct {
	Id       int
	SpaceId  int
	MemberId int
	IsOwner  bool
}
