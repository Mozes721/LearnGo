package test_examples

import (
	"fmt"
	"testing"
)

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

type GetPetNamesStub struct {
	Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}

func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
	switch userID {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userID)
	}
}

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowball"}},
		{"case3", "3", nil},
	}
	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}

