# Pod Deletion Notification Controller

This is a very basic simple controller to send slack messages to a designated  workspace channel when a pod dies

## ENV REQUIRED

- OAUTH_TOKEN (From Slack)
- CHANNEL_ID (Channel to send messages)

## How to get Slack Token

- From your workspace -> settings and administration
- Manage apps
- Click on Build (top right corner)
- Create the app


## Images to demostrate 

- Delete or Restart a pod
![Pod Deletion](https://github.com/DiptoChakrabarty/Pod-Delete-Notify-Controller/blob/main/images/pod_restartpng.png)

- Get a log message from the app
![Terminal Log](https://github.com/DiptoChakrabarty/Pod-Delete-Notify-Controller/blob/main/images/delete_msg.png)

- Slack notification in channel
![Slack Notification](https://github.com/DiptoChakrabarty/Pod-Delete-Notify-Controller/blob/main/images/poddeleteslack.png)