const chats = (state = [], action) => {
    switch (action.type) {
        case 'SET_CHATS':
            return action.chats
        case 'NEW_MESSAGE':
            let message = action.message
            return state.map((chat) => {
                if (chat.id === message.chat_id) {
                    chat.messages.push(message)
                }
                chat.lastMessageId = message.id
                return chat
            })
        default:
            return state
    }
}

export default chats