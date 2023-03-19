# ChatGPT go client

Command-line utility for having fun with ChatGPT. 

## Installation

```
go build
```

or 

```
go install
```

## Token

Go to https://platform.openai.com/account/api-keys to get API keys for your account. This may require you to set up billing to pay for your usage.

Then, make that key an environmental variable in your shell.

```
Example:
    export CHATGPT_TOKEN=<your-token>
```

## Running

When running, just type as you normally would with any chat. Press enter or return to send the chat, and ChatGPT will send a response back. Type `exit` to exit the program.

It does carry on a conversation, but there are limits to how long the conversation can be. See `How it works` below.

```
$ chatgpt-go                     
Human> Tell me a joke.
ChatGPT> Why did the chicken cross the playground? To get to the other slide!
Human> Tell me the same joke, but with a flamingo.
ChatGPT> Why did the flamingo cross the playground? To get to the other pink slide!
Human> exit
```


## How it works

This uses the /v1/completions API to interact with the `text-davinci-003` language model. Normally, you can't carry on a conversation with just this API endpoint because each request is a unique and new request, there's no way to indicate a conversation ID (at least with the current public API).

To get around this, this program prepends this text to the first message.

```
This is a conversation with a human being. Provide the next message that would be said as a reply to the last line of text.

<FIRST_MESSAGE>
```

And then adds the response to this growing conversation. The POST command for the second message for the human is as follows:

```
This is a conversation with a human being. Provide the next message that would be said as a reply to the last line of text.

<FIRST_MESSAGE>
<FIRST_REPLY>
<SECOND_MESSAGE>
```

And so on. This means eventually the conversation gets so large that it causes an error.