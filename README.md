# 📰 WriteUp App

## 👀 ʀᴇQᴜɪʀᴇᴍᴇɴᴛꜱ :
```yaml
- name : Python
    - type    : Programming Language
    - version : >= 3.10

- name : Go
    - type    : Programming Language
    - Version : >= 1.19

- name : MySQL
    - type : Database

- name : Git
    - type    : Tool
    - version : >= 2.39

- name : Linux
    - type : Distro
```

## 🦾 ꜰᴇᴀᴛᴜʀᴇꜱ:
```yaml
- name : Social
    - Description : This App Will Send New Available WriteUps To [ Discord , Telegram ]

- name : resource
    - Description : WriteUps Will Send From Diffrent Tags Of medium.com

- name : System Log
    - Description : System Log Available in log/log.log

- name : Automated Configuration
    - Description : Automation Scripts To Config Your MySQL & Check [Files, Directories, Packages], Written in Python
```

## 🏁 ɪɴꜱᴛᴀʟʟᴀᴛɪᴏɴ:
```yaml
- Step One :
    - Description : Clone Repository
    - Command     : git clone https://github.com/JesusKian/WriteUp.git

- Step Two :
    - Description : Go To Project's Directory
    - Command     : cd WriteUp

- Step Three :
    - Description : Fill The Variables in config.env file
    - Variables :
        - TELEGRAM_API : Enter Your Telegram's Bot's API
        - DISCORD_WEBHOOK : Enter Your Discord's Webhook's URL
        - CHANNEL_NAME : Enter Your Channel's ID with @
        - MYSQL_USERNAME : Enter Your MySQL Username (default=root)
        - MYSQL_PASSWORD : Enter Your MySQL Password (default=system password)
        - Distro : Enter Number in range [1 , 4]
            - 1 : Debian
            - 2 : Arch
            - 3 : Fedora
            - 4 : Another

- Step Four :
    - Description : You Must Run Config Files
    - Commands : 
        - python3 -m pip install -r requirements.txt
        - python3 run.py

- Step Five :
    - Description : go run ./main.go
    - Command     : ./main
```


## ❓ ʜᴏᴡ ᴛᴏ ᴄʀᴇᴀᴛᴇ ᴛᴇʟᴇɢʀᴀᴍ ʙᴏᴛ
```yaml
- Step One :
    - Description : Go To BotFather : https://t.me/BotFather

- Step Two :
    - Description : Start The Bot
    - Command     : /start

- Step Three :
    - Description : Create New Bot
    - Command     : /newbot

- Step Four :
    - Description : Enter The Name For Your Bot
    - Example     : YoungBoy

- Step Five :
    - Description : Enter The Username For Your Bot
    - Example     : YoungBotBot

- Step Six :
    - Description : Now Copy Your Bot Token
    - Example     : 1234567890:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

## ⚙️ Config .env File Example
```yaml
TELEGRAM_API=1234567890:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
DISCORD_WEBHOOK=https://discord.com/api/webhooks/.../...
CHANNEL_NAME=https://t.me/WriteUpChannel
MYSQL_USERNAME=JesusKian
MYSQL_PASSWORD=ExampleP@33W0rd
DISTRO=2
```

## ⚡️ Automate WriteUp Sender
```yaml
- name : Crontab
    - Description : You Can Automate Process to Run main.go File
    - Example : With Below Code, You Can Run main.go Every 1 Hours
    - Command : 0 */1 * * * /usr/bin/go /PATH/TO/WriteUp/main.go
    - Resource : https://geekflare.com/crontab-linux-with-real-time-examples-and-tools/
```

## 🔴 Important
```yaml
- Create a Mysql User And Grant Privileges For it =)
```

## 📹 Watch Video Below
[![asciicast](https://asciinema.org/a/jDtA4QHs0s4zKJDXOdUK7dVdu.svg)](https://asciinema.org/a/jDtA4QHs0s4zKJDXOdUK7dVdu)
