const socket = io('ws://localhost:5001', {transports: ['websocket']});
const room = document.getElementById("room");
const form = room.querySelector("form");

// listen for messages
socket.on('/message', function(message) {

    console.log('new message');
    addMessage(message);
});

socket.on('connect', function () {

    console.log('socket connected');

    //send something
    // welcome msg
    const input = form.querySelector("input");
    const value = input.value;
    socket.emit('/', {name: "welcome", message: "welcome (userID)"}, function(result) {
        console.log(result);
    });
});

function addMessage(msg) {
    const ul = room.querySelector("ul");
    const li = document.createElement("li");
    li.innerText = msg;
    ul.appendChild(li);
}

var roomId = "";

function handleMsgClick(event) {
    event.preventDefault();
    const input = form.querySelector("input");
    const value = input.value;
    if (roomId === "") {
        roomId = "roomID";
        // enter room
    } else {
        // new message
    }
    socket.emit('/chat', {name: "my name", message: value}, function(result) {

        console.log('sended successfully');
        console.log(result);
    });

    input.value = "";
}

function init() {

}

form.addEventListener("submit", handleMsgClick);
