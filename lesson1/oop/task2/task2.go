package main

import "fmt"

// Базовый интерфейс
type User interface {
	GetUsername() string
	HasPermission(permission string) bool
	GetRole() string
}

type BasicUser struct {
	username    string
	permissions map[string]bool
}

func NewBasicUser(username string) *BasicUser {
	user := BasicUser{
		username: username,
		permissions: map[string]bool{
			"read": true,
		},
	}
	return &user
}

func (bu BasicUser) GetUsername() string {
	return bu.username
}

func (bu BasicUser) HasPermission(permission string) bool {
	if _, ok := bu.permissions[permission]; ok {
		return true
	}
	return false
}

func (bu BasicUser) GetRole() string {
	return "basic user"
}

///////////////////////////////////////////////////

type Moderator struct {
	BasicUser
}

//можно не дублировать инициализацию базового пользователя
/*
{
base := NewBasicUser(username)

base.permissions["edit"]: true,
//тут добалвяем новые методы

moderator := Moderator {
		BasicUser: *base,
	}
}

*/
func NewModerator(username string) *Moderator {
	moderator := Moderator{
		BasicUser: BasicUser{
			username: username,
			permissions: map[string]bool{
				"read":     true,
				"edit":     true,
				"ban_user": true,
			},
		},
	}
	return &moderator
}

func (mod Moderator) GetRole() string {
	return "moderator user"
}

////////////////////////////////////////
type Admin struct {
	Moderator
}

func NewAdmin(username string) *Admin {
	base := NewModerator(username)

	base.permissions["delete"] = true
	base.permissions["manage_roles"] = true

	admin := Admin{
		Moderator: *base,
	}
	return &admin
}

func (a Admin) GetRole() string {
	return "admin user"
}

///////////////////////////////////////
func main() {
	// --- Шаг 1: создаём по одному каждого типа ---
	basic := NewBasicUser("вася")
	mod := NewModerator("петя")
	admin := NewAdmin("маша")

	// --- Шаг 2: складываем в срез ИНТЕРФЕЙСОВ ---
	// сюда влезают три разных типа, потому что каждый удовлетворяет User
	users := []User{basic, mod, admin}

	// --- Шаг 3: один цикл — разное поведение (полиморфизм) ---
	for _, u := range users {
		fmt.Printf("Роль: %-12s имя: %s\n", u.GetRole(), u.GetUsername())
		fmt.Printf("  read?        %v\n", u.HasPermission("read"))
		fmt.Printf("  edit?        %v\n", u.HasPermission("edit"))
		fmt.Printf("  ban_user?    %v\n", u.HasPermission("ban_user"))
		fmt.Printf("  delete?      %v\n", u.HasPermission("delete"))
		fmt.Printf("  manage_roles? %v\n", u.HasPermission("manage_roles"))
		fmt.Println()
	}
}
