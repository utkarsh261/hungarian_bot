# Contributing to the bot 🔥✨

Contributions are always welcome, no matter how large or small!🙂

The Guide assumes you have Go installed. If not, head [here](https://golang.org/doc/install).

### Setting up the bot

1. Ping [Botfather](https://telegram.me/botfather) on Telegram and make your instance of OSDC-Bot 🤖 bot by selecting `/newbot` from the options provided.
2. Copy the `TELEGRAM_TOKEN` provided by Botfather.
3. Make sure you have followed all the above steps and are in the `telegram-bot` directory.
4. If you have installed golang, run `go build .`
5. Wait ⏳
6. Run `export TELEGRAM_TOKEN=<TELEGRAM_TOKEN>`
7. Now, run `./telegram-bot`. The bot would be running at the username provided by you on telegram. 🚀
8. If you would like to make some changes and contribute to the bot, follow the steps below.

## Making Pull-Requests (Contributions)

Having setup the bot and tested its working, if you want to contribute to it, follow the steps below :

1. Make a new branch of the project using the `git checkout` command :
```
$ git checkout -b "Name-of-the-branch"
```
2. Make changes according to the issue. Test the working of the changes.
3. Add the changes to staging area using the `git add` command :
```
$ git add .
```
4. Commit the changes made using the `git commit` commad :
```
$ git commit -m "Commit-message"
```
5. Push the chages to your branch on Github using the `git push` command :
```
$ git push -u origin "Name-of-the-branch-from-step-1"
```
6. Then, go to your forked repository and make a Pull Request 🎉. Refer [this](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) for more details.


