const messageContainer = document.getElementById('message-container');
const messageForm = document.getElementById('message-form');
const messageInput = document.getElementById('message-input');

// Connect to websocket
const socket = new WebSocket('ws://localhost:3000');

// Handle incoming messages
socket.onmessage = function(event) {
  const message = JSON.parse(event.data);
  const messageElement = document.createElement('div');
  messageElement.innerText = `${message.sender}: ${message.text}`;
  messageContainer.appendChild(messageElement);
};

// Handle form submission
messageForm.addEventListener('submit', function(event) {
  event.preventDefault();
  const message = {
    sender: 'You',
    text: messageInput.value
  };
  socket.send(JSON.stringify(message));
  messageInput.value = '';
});