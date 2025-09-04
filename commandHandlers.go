package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gfrei/gator-cli/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login error: no username")
	}

	user := cmd.args[0]

	return setUser(s, user)
}

func setUser(s *state, username string) error {
	_, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %q not registered", username)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("logged in as: %q\n", username)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register error: no username")
	}

	fullName := (strings.Join(cmd.args, " "))

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      fullName,
	})

	if err != nil {
		return err
	}

	fmt.Printf("registered user %q\n", user.Name)

	return setUser(s, user.Name)
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed error: add name and url")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("registered feed %q\n", feed.Name)

	return followFeed(s, feed.ID, user.ID)
}

func handlerGetAllFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("(%v) %v: %v\n", user.Name, feed.Name, feed.Url)
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	count, err := s.db.CountUsers(context.Background())
	if err != nil {
		return err
	}

	err = s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("deleted %v users\n", count)

	return nil
}

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("feedfollow error: select a feed url")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	return followFeed(s, feed.ID, user.ID)
}
func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("unfollow error: select a feed url")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("user %q unfollows feed %q\n", user.Name, feed.Name)

	return nil
}

func followFeed(s *state, feedId, userId uuid.UUID) error {
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
		FeedID:    feedId,
	})
	if err != nil {
		return err
	}

	fmt.Printf("user %q now follows %q\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Printf("user %q is following\n", s.config.CurrentUserName)

	for _, follow := range follows {
		fmt.Printf("\t- %q\n", follow.FeedName)
	}

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name != s.config.CurrentUserName {
			fmt.Printf("* %v\n", user.Name)
		} else {
			fmt.Printf("* %v (current)\n", user.Name)
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) == 0 {
	// 	return fmt.Errorf("agg error: add an url")
	// }

	// feed, err := fetchFeed(context.Background(), cmd.args[0])
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
