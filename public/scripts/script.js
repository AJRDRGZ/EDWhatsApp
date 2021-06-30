import GetRandomName from './names-generator.js'

new lc_emoji_picker('input', {
  trigger_position: {
    top: '6px',
    right: '6px',
  },
})

// conn variable for the websocket connection
let conn

// chat elements
const chatTalk = document.getElementById('chat-talk')
const chatForm = document.getElementById('chat-form')
const chatInput = document.getElementById('chat-input')

const NICKNAME = GetRandomName()

window.addEventListener('load', event => {
  if (!window['WebSocket']) {
    chatTalk.insertAdjacentHTML('beforeend', `<b>Your browser does not support WebSockets.</b>`)
    return
  }

  conn = new WebSocket('ws://' + document.location.host + `/ws?nickname=${NICKNAME}`)

  conn.onclose = function (event) {
    chatTalk.insertAdjacentHTML('beforeend', `<b>Connection closed.</b>`)
  }

  conn.onmessage = function (event) {
    const message = JSON.parse(event.data)
    appendMessage(message.nickname, 'outgoing', message.content)
  }
})

chatForm.addEventListener('submit', event => {
  event.preventDefault()

  const message = chatInput.value
  if (!message) return

  appendMessage(NICKNAME, 'incoming', message)
  chatInput.value = ''

  conn.send(JSON.stringify({ nickname: NICKNAME, content: message }))
})

function appendMessage(nickname, type, message) {
  const msgBubble = `
    <div class="msg-bubble msg-${type}">
        <div class="msg-info">
            <span class="msg-name">${nickname}</span>
            <span class="msg-time">${formatDate(new Date())}</span>
        </div>
    
        <div class="msg-text">
            ${message}
        </div>
    </div>
  `

  const doScroll = chatTalk.scrollTop > chatTalk.scrollHeight - chatTalk.clientHeight - 1
  chatTalk.insertAdjacentHTML('beforeend', msgBubble)

  if (doScroll) {
    chatTalk.scrollTop = chatTalk.scrollHeight - chatTalk.clientHeight
  }
}

// ---------------------------------------------------------
// helpers
// ---------------------------------------------------------

function formatDate(date) {
  const hours = '0' + date.getHours()
  const minutes = '0' + date.getMinutes()

  return `${hours.slice(-2)}:${minutes.slice(-2)}`
}
