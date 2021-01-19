package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tbot.BotAPI

//ButtonLinks sends the parsed arguements as an inline buttons to the specified chat ID
func ButtonLinks(ID int64, ButtonText string, ButtonURL string, MessageText string) {
	var button = tbot.NewInlineKeyboardMarkup(
		tbot.NewInlineKeyboardRow(
			tbot.NewInlineKeyboardButtonURL(ButtonText, ButtonURL),
		),
	)
	msg := tbot.NewMessage(ID, MessageText)
	msg.ReplyMarkup = button
	bot.Send(msg)
}

//extracts the details of the user who sent the message to check whether the user is creator/admin. Returns true in this case else false.
func check_for_mod(ID int64, userid int) bool {
	response, _ := bot.GetChatMember(tbot.ChatConfigWithUser{
		ChatID: ID,
		UserID: userid,
	})
	if response.IsCreator() || response.IsAdministrator() {
		return true
	}
	return false
}

func google(ID int64, message *tbot.Message) {
	base := "http://letmegooglethat.com/?q="
	if message.ReplyToMessage != nil {
		if message.ReplyToMessage.Text != "" {
			query := message.ReplyToMessage.Text
			query = strings.Replace(query, " ", "+", -1)
			log.Printf("Query: " + query)
			url := base + query
			log.Printf(url)
			ButtonLinks(ID, "Click Here", string(url), "Find your answer below ;)")
		} else {
			bot.Send(tbot.NewMessage(ID, "I don't know what to Google🙁. Reply to a text message so that I can Google it."))
		}
	} else {
		bot.Send(tbot.NewMessage(ID, "I don't know what to Google🙁. Reply to a text message so that I can Google it."))
	}
}

// extract member details
func member_details(msg *tbot.Message) string {
	log.Printf("[%s]", msg.From.UserName)
	return msg.From.UserName
}

func help(ID int64) {
	msg := ` Use one of the following commands to :
	/github - Get a link to KPH's Github page.
	/telegram - Get an invite link for KPH's Telegram Group.
	/facebook - Get the link of KPH's facebook page.
	/discord - Get a link to KPH's Discord server and come hang out!
	/resources - A curated list of Competitive Programming resources.
	/meetups - Details of meetups organised by the Hub and related resources.
	/workshops - Details of meetups organised by the Hub and related resources.
	/contests - Details of previous contests organised by the Hub.
	/paste - Create a pastebin for contents of a text
	/let_me_google_that - Let me google that for you.

	* Admin only commands: 
	/offtopic - Mark a text as offtopic.
	/spam - Mark a text as spam
	
	To contribute to this bot : https://github.com/Knuth-Programming-Hub/
	`
	bot.Send(tbot.NewMessage(ID, msg))
}

func offtopic(ID int64) {
	response := "Hi, please continue to off-topic group for further discussion on this topic."
	ButtonLinks(ID, "Join Off-Topic Group", "https://t.me/joinchat/JY7tkxjqeHx1UY0K14fc0Q", response)
}

func main() {
	var err error
	bot, err = tbot.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		fmt.Println("error in auth")
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tbot.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 5000)
	updates.Clear()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		ID := update.Message.Chat.ID

		log.Print(update.Message.Photo)
		if update.Message.Photo != nil {
			response := "Hi @" + member_details(update.Message) + ", if the picture(s) you just sent is of code, please use https://ideone.com/ or https://pastebin.com/ and send the link here. Else please ignore, sorry hehe😅"

			ButtonLinks(ID, "Ideone.com", "https://ideone.com/", response)
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand() {
			switch update.Message.Command() {

			case "start":
				ButtonLinks(ID, "Click Here", "https://t.me/joinchat/LGo0IhZoPRjRjBJHJPf3OA", "Hello , I'm Hungarian Bot, named after a famous algorithm, Use /help to know more. To join Knuth Programming Hub group:")
			case "help":
				help(ID)
			case "resources":
				ButtonLinks(ID, "Resources", "https://github.com/Knuth-Programming-Hub/CP-Resources", "A curated list of Competitive Programming resources.")
			case "workshops":
				ButtonLinks(ID, "Workshops", "https://github.com/Knuth-Programming-Hub/Workshops", "Repo for keeping record of Workshops organised by the Hub and managing related resources.")
			case "meetups":
				ButtonLinks(ID, "Meetups", "https://github.com/Knuth-Programming-Hub/Meetups", "Checkout out details of our previous meetups.")
			case "contests":
				ButtonLinks(ID, "Knuth Contests", "https://github.com/Knuth-Programming-Hub/Knuth-Contests", "Contests orgaranised by Knuth Programming Hub.")
			case "github":
				ButtonLinks(ID, "Github", "http://github.com/Knuth-Programming-Hub/", "Checkout our Github.")
			case "discord":
				ButtonLinks(ID, "Discord", "https://discord.gg/j5WdgpbzWP", "Come hang out on discord!")
			case "facebook":
				ButtonLinks(ID, "Click here", "https://www.facebook.com/groups/jiit.knuth/", "KPH's Facebook Page.")
			case "telegram":
				_, err := bot.Send(tbot.NewMessage(ID, "https://t.me/joinchat/LGo0IhZoPRjRjBJHJPf3OA"))
				log.Print(err)
			case "paste":
				paste(ID, update.Message)
			case "offtopic":
				check := check_for_mod(ID, update.Message.From.ID)
				if check == true {
					offtopic(ID)
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, this looks like an admin only command💢."))
				}
			case "addevent":
				check := check_for_mod(ID, update.Message.From.ID)
				if check == true {
					bot.Send(tbot.NewMessage(ID, "This is still a beta feature."))
					addevent(ID, update.Message.Text)
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, this looks like an admin only command."))
				}
			case "listevents":
				bot.Send(tbot.NewMessage(ID, "This is still a beta feature."))
				listevents(ID)
			case "spam":
				check := check_for_mod(ID, update.Message.From.ID)
				if check == true {
					if update.Message.ReplyToMessage != nil {
						res := "Hi @" + member_details(update.Message.ReplyToMessage) + ", there are a lot of people in this group and it can get pretty annoying for them with these kind of messages, please don't spam again. Thanks🙂."
						msg := tbot.NewMessage(ID, res)

						bot.Send(msg)
					} else {
						bot.Send(tbot.NewMessage(ID, "Reply to a text to mark it as spam! 💢."))
					}
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, this looks like an admin only command💢."))
				}
			case "let_me_google_that":
				google(ID, update.Message)
			default:
				{
					msg, err := bot.Send(tbot.NewMessage(ID, "I don't know this command🙁"))

					log.Print(err)
					timer1 := time.NewTimer(5 * time.Second)
					<-timer1.C

					bot.DeleteMessage(tbot.NewDeleteMessage(ID, msg.MessageID))
				}
			}
		} else {
			analyse(ID, update.Message)
		}
	}
}
