package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	var p fastjson.Parser
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		tmp, e := p.Parse(scanner.Text())
		err = e
		if err != nil {
			return
		}
		result[i] = User{
			ID:       tmp.GetInt("ID"),
			Address:  string(tmp.GetStringBytes("Address")),
			Email:    string(tmp.GetStringBytes("Email")),
			Name:     string(tmp.GetStringBytes("Name")),
			Password: string(tmp.GetStringBytes("Password")),
			Phone:    string(tmp.GetStringBytes("Phone")),
			Username: string(tmp.GetStringBytes("Username")),
		}
		i++
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		if user.Email == "" {
			continue
		}
		if matched := strings.Contains(user.Email, domain); matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] + 1
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
