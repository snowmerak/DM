package manager

type Manager interface {
	CreateGroup(namespace string, groupId string) error
	DeleteGroup(namespace string, groupId string) error
	JoinGroup(namespace string, groupId string, userId string) error
	LeaveGroup(namespace string, groupId string, userId string) error
	GetGroupMembers(namespace string, groupId string) ([]string, error)
}
