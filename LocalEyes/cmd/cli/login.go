package cli

import (
	"fmt"
	"localEyes/constants"
	"localEyes/internal/services"
	"localEyes/utils"
)

func login(userService *services.UserService, questionService *services.QuestionService, postService *services.PostService) {
	fmt.Println(constants.Blue + "==============================")
	fmt.Println("LOGIN")
	fmt.Println("=============================" + constants.Reset)
	username := utils.PromptInput("Enter your username:")
	password := utils.PromptPassword(constants.Cyan + "Enter your password:" + constants.Reset)
	user, err := userService.Login(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(constants.Green + "\nUser logged in successfully" + constants.Reset)

	for {
		fmt.Println(constants.Blue + "\n1.View my Profile")
		fmt.Println("2.Create Food post")
		fmt.Println("3.Create Travel post")
		fmt.Println("4.Create Shopping post")
		fmt.Println("5.Create Other post")
		fmt.Println("6.Update Post")
		fmt.Println("7 View Posts")
		fmt.Println("8.Filter Posts")
		fmt.Println("9.Open Post")
		fmt.Println("10.Like Post")
		fmt.Println("11.Delete Post")
		fmt.Println("12.Deactivate account")
		fmt.Println("13.Return" + constants.Reset)
		choice := utils.GetChoice()
		switch choice {
		case 1:
			fmt.Println(constants.Magenta + "\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("Welcome ", user.Username)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" + constants.Reset)
			fmt.Println("City:", user.City)
			fmt.Println("Type of user:", user.Tag)
			fmt.Println("Living in City Till:", user.DwellingAge)
		case 2:
			title := utils.PromptInput("Enter post title:")
			content := utils.PromptInput("Enter post content:")
			err := postService.CreatePost(user.UId, title, content, "food")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.Green+"Post created:", title)
			}

		case 3:
			title := utils.PromptInput("Enter post title:")
			content := utils.PromptInput("Enter post content:")
			err := postService.CreatePost(user.UId, title, content, "travel")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.Green+"Post created:", title)
			}
		case 4:
			title := utils.PromptInput("Enter post title:")
			content := utils.PromptInput("Enter post content:")
			err := postService.CreatePost(user.UId, title, content, "shopping")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.Green+"Post created:", title)
			}
		case 5:
			title := utils.PromptInput("Enter post title:")
			content := utils.PromptInput("Enter post content:")
			err := postService.CreatePost(user.UId, title, content, "other")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.Green+"Post created:", title)
			}
		case 6:
			myPosts, err := postService.GiveMyPosts(user.UId)
			if err != nil {
				fmt.Println(constants.Red + "Error loading posts:" + err.Error() + constants.Reset)
			} else {
				displayPosts(myPosts)
			}
			PId, err := utils.PromptID("Enter post id to update:")
			if err != nil {
				fmt.Println(constants.Red + err.Error() + constants.Reset)
			}
			title := utils.PromptInput("Enter new post title:")
			content := utils.PromptInput("Enter new post content:")
			err = postService.UpdateMyPost(PId, user.UId, title, content)
			if err != nil {
				fmt.Println(constants.Red + "Error updating post:" + err.Error() + constants.Reset)
			} else {
				fmt.Println(constants.Green+"Post updated:", title)
			}
		case 7:
			posts, err := postService.GiveAllPosts()
			if err != nil {
				fmt.Println(constants.Red + "Error loading posts:" + err.Error() + constants.Reset)
			} else {
				displayPosts(posts)
			}

		case 8:
			var filterType string
			for {
				filterType = utils.PromptInput("Enter filter [food/travel/shopping/other]:")
				if utils.ValidateFilter(filterType) {
					break
				} else {
					fmt.Println("Invalid filter type:", filterType)
				}
			}
			posts, err := postService.GiveFilteredPosts(filterType)
			if err != nil {
				fmt.Println(constants.Red + "Error loading posts:" + err.Error() + constants.Reset)
			} else {
				displayPosts(posts)
			}

		case 9:
			PId, err := utils.PromptID("Enter post id to open:")
			if err != nil {
				fmt.Println(constants.Red + err.Error() + constants.Reset)
				break
			}
			openPost(questionService, postService, PId, user.UId)

		case 10:
			PId, err := utils.PromptID("Enter post id to like:")
			err = postService.Like(PId)
			if err != nil {
				fmt.Println(constants.Red + "Error liking post:" + err.Error() + constants.Reset)
			} else {
				fmt.Println(constants.Green + "Post Liked" + constants.Reset)
			}

		case 11:
			myPosts, err := postService.GiveMyPosts(user.UId)
			if err != nil {
				fmt.Println(constants.Red + "Error loading posts:" + err.Error() + constants.Reset)
			} else {
				displayPosts(myPosts)
			}
			PId, err := utils.PromptID("Enter post id to delete:")
			if err != nil {
				fmt.Println(constants.Red + "Error taking postId input:" + err.Error() + constants.Reset)
			}
			err = postService.DeleteMyPost(user.UId, PId)
			err = questionService.DeleteQuesByPId(PId)
			if err != nil {
				fmt.Println(constants.Red + "Error deleting question:" + err.Error() + constants.Reset)
			} else {
				fmt.Println(constants.Green + "Post deleted successfully" + constants.Reset)
			}
		case 12:
			err := userService.DeActivate(user.UId)
			if err != nil {
				fmt.Println(constants.Red + "Error Deactivating user:" + err.Error() + constants.Reset)
			} else {
				fmt.Println(constants.Green + "User Deactivated successfully" + constants.Reset)
				return
			}
		case 13:
			return
		default:
			fmt.Println(constants.Red + "Invalid Choice,Try Again" + constants.Reset)
		}
	}
}
