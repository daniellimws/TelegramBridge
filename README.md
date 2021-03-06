# Telegram Bridge
  This is a GitHub Action plugin to notify about the changes in GitHub to Telegram Chat via Bot. This could highly help in incresing the development pace.

  Made some changes.

## What all can this do ?
  If you're using the provided [main.yaml](https://github.com/GokulDas027/TelegramBridge/blob/master/main.yaml) (recommended)
  
  Will Notify about :
  * Pushes
  * Pull Requests (opened, closed, edited, ready_for_review, review_requested, reopened)
  * Pull Request Review Comments (created)
  * Issues (opened, edited, pinned, closed, reopened, assigned, labeled)
  * Issue Comments (created)
  
  > the limits are not the boundary, but set for reducing the message bombarding, and it can be easily changed at the main.yaml file at your repository. This could help [Action Triggers Docs](https://developer.github.com/webhooks/).
  
## How to Use this ?
  This will be available as a plugin [here](https://github.com/marketplace/actions/telegrambridge).
  
  Or can be used through this [main.yaml](https://github.com/GokulDas027/TelegramBridge/blob/master/main.yaml) file
    > move it to <repository>/.github/workflows/main.yml
  
  Connect it to a [Telegram bot](https://core.telegram.org/bots).
  * Token is the token of the Bot given by BotFather
  * Id is the Channel id or your user id, depends on how you need the message to be delivered.
  
  ### [Check here for step by step instructions](https://github.com/GokulDas027/TelegramBridge/blob/master/instruction.md)

## How does the message looks like ?

![Sample Message](https://drive.google.com/uc?export=view&id=1TsZZi9yS6qC_oQVdoOrRLbL8Zq6WeMqE)

## What's next ?

This is a tailored implementation. This can be changed or re-tailored at [main.go](https://github.com/GokulDas027/TelegramBridge/blob/master/main.go). Curresponting chages must be made at the action's main.yaml too

Thanks to [Yanzay](https://github.com/yanzay) and [Athul](github.com/athul).
