# GPT4All-Server-Mode-wrapper-golang

## This project is a wrapper implemented in golang language of [gpt4all](https://github.com/nomic-ai/gpt4all/tree/main) API

## **Warning, this project is only for personal test development and use, and does not guarantee any degree of stability and compatibility**

### For the actual API call content, you can refer to [API](https://docs.gpt4all.io/gpt4all_chat.html#server-mode)

**Since the server address binding in the [source code line 83](https://github.com/nomic-ai/gpt4all/blob/main/gpt4all-chat/server.cpp) is hardcoded for *LocalHost*, at least for now, this wrapper will only work on devices with the same ip as the server**
