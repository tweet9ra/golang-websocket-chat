import React, {useEffect} from "react";
const moment = require('moment')

const ChatHistory = ({chat, currentUserId}) => {
  let messagesContainer

  useEffect(() => {
    if (messagesContainer) {
      messagesContainer.scrollTop = 100000
    }
  })

  if (!chat) {
    return null
  }

  const usersById = {}
  chat.users.forEach((el) => {
    usersById[el.id] = el
  })


  return (
    <div className="chat-history" ref={node => messagesContainer = node}>
      <ul>
        {
          chat && chat.messages.map((message, index) => {
            let showMessageHeader = index === 0 || chat.messages[index - 1].user_id !== message.user_id
            return <li key={message.id} className='clearfix'>
              { showMessageHeader
                  ? <div className={'message-data ' + (message.user_id === currentUserId ? 'message-data-me' : '')}>
                    <span className="message-data-name">{usersById[message.user_id].email}</span>&nbsp;&nbsp;
                    <span className="message-data-time">{moment(message.created_at).format("hh:mm:ss")}</span>
                  </div> : ''
              }
              <div
                className={
                  'message '
                  + (message.user_id === currentUserId ? 'my-message' : 'other-message float-right')
                  + (showMessageHeader ? '' : ' hide-message-header')
                }>
                {message.text}
              </div>
            </li>
          })
        }
      </ul>

    </div>
  )
}

export default ChatHistory